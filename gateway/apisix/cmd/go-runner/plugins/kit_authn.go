package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/session"
	conf2 "github.com/goxiaoy/go-saas-kit/pkg/conf"
	v1 "github.com/goxiaoy/go-saas-kit/saas/api/tenant/v1"
	uremote "github.com/goxiaoy/go-saas-kit/user/remote"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/sessions"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strings"
)

func init() {
	err := plugin.RegisterPlugin(&KitAuthn{})
	if err != nil {
		log.Fatalf("failed to register plugin kit_authn: %s", err)
	}
}

type KitAuthn struct {
}

type KitAuthConf struct {
}

var (
	tokenizer           jwt.Tokenizer
	tokenManager        api.TokenManager
	apiClient           *conf2.Client
	apiOpt              *api.Option
	sessionInfoStore    sessions.Store
	rememberStore       sessions.Store
	securityCfg         *conf2.Security
	userTenantValidator *uremote.UserTenantContributor
	refreshProvider     session.RefreshTokenProvider
	ts                  common.TenantStore
)

func Init(
	t jwt.Tokenizer,
	tmr api.TokenManager,
	clientName api.ClientName,
	services *conf2.Services,
	security *conf2.Security,
	userTenant *uremote.UserTenantContributor,
	tenantStore common.TenantStore,
	refreshTokenProvider session.RefreshTokenProvider,
	logger klog.Logger,
) error {
	tokenizer = t
	tokenManager = tmr
	clientCfg, ok := services.Clients[string(clientName)]
	if !ok {
		return errors.New(fmt.Sprintf(" %v client not found", clientName))
	}
	apiClient = clientCfg
	apiOpt = api.NewOption(true, api.NewUserPropagator(logger), api.NewClientPropagator(false, logger))
	securityCfg = security
	sessionInfoStore = session.NewSessionInfoStore(security)
	rememberStore = session.NewRememberStore(security)
	userTenantValidator = userTenant
	refreshProvider = refreshTokenProvider
	ts = tenantStore
	return nil
}

func (p *KitAuthn) Name() string {
	return "kit_authn"
}

func (p *KitAuthn) ParseConf(in []byte) (interface{}, error) {
	conf := KitAuthConf{}
	err := json.Unmarshal(in, &conf)
	if err != nil {
		return nil, err
	}
	return conf, err
}

func abortWithError(err error, w http.ResponseWriter) {
	//use error codec
	fr := kerrors.FromError(err)
	w.WriteHeader(int(fr.Code))
	khttp.DefaultErrorEncoder(w, &http.Request{}, err)
}

func (p *KitAuthn) Filter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	var err error
	ctx := r.Context()
	tracer := tracing.NewTracer(trace.SpanKindServer)
	var span trace.Span
	ctx, span = tracer.Start(ctx, p.Name(), propagation.HeaderCarrier(r.Header().View()))
	defer func() { tracer.End(ctx, span, nil, nil) }()

	//clean internal headers
	for s, _ := range r.Header().View() {
		if strings.HasPrefix(strings.ToLower(s), api.InternalKeyPrefix) {
			//just clean headers
			log.Infof("clean untrusted internal header %s", s)
			r.Header().Set(s, "")
		}
	}

	trace := r.Header().Get("traceparent")
	//https://github.com/apache/apisix/issues/6728
	log.Infof("trace: %s", trace)

	ctx, err = Saas(ctx, ts, "", w, r)
	//format error
	if err != nil {
		if errors.Is(err, common.ErrTenantNotFound) {
			err = v1.ErrorTenantNotFound("")
		}
		abortWithError(err, w)
		//stop
		return
	}

	uid := ""
	clientId := ""

	//session auth
	header := r.Header().View()
	ctx = sessions.NewRegistryContext(ctx, header)

	s, _ := session.GetSession(ctx, header, sessionInfoStore, securityCfg)

	rs, _ := session.GetRememberSession(ctx, header, rememberStore, securityCfg)

	stateWriter := session.NewClientStateWriter(s, rs, w, header)

	ctx = session.NewClientStateWriterContext(ctx, stateWriter)
	state := session.NewClientState(s, rs)
	ctx = session.NewClientStateContext(ctx, state)

	if len(state.GetUid()) == 0 && len(state.GetRememberToken()) > 0 {
		//call refresh
		log.Infof("call refresh token")
		err := refreshProvider(ctx, state.GetRememberToken())
		if err != nil {
			log.Errorf("refresh fail %v", err)
			//abort with error
			abortWithError(err, w)
			return
		}
	}

	//set uid from cookie
	uid = state.GetUid()
	ctx = authn.NewUserContext(ctx, authn.NewUserInfo(uid))

	//extract token
	var t = ""
	if auth := r.Header().Get(jwt.AuthorizationHeader); len(auth) > 0 {
		t = jwt.ExtractHeaderToken(auth)
	}
	if len(t) == 0 {
		t = r.Args().Get(jwt.AuthorizationQuery)
	}
	//jwt auth
	if len(t) > 0 {
		if claims, err := jwt.ExtractAndValidate(tokenizer, t); err != nil {
			log.Errorf("fail to extract and validate token %s", err)
		} else {
			if claims.Subject != "" {
				uid = claims.Subject
			} else if claims.Subject != "" {
				uid = claims.Uid
			}
			clientId = claims.ClientId
		}
	}

	ctx = authn.NewUserContext(ctx, authn.NewUserInfo(uid))

	//check tenant and user mismatch
	ti, _ := common.FromCurrentTenant(ctx)
	trCtx := common.NewTenantResolveContext(ctx)
	trCtx.TenantIdOrName = ti.GetId()

	log.Infof("resolve user: %s client: %s tenantId: %s", uid, clientId, ti.GetId())
	err = userTenantValidator.Resolve(trCtx)
	if err != nil {
		log.Errorf("%s", err)
		// user can not in this tenant
		// use error codec
		fr := kerrors.FromError(err)
		w.WriteHeader(int(fr.Code))
		khttp.DefaultErrorEncoder(w, &http.Request{}, err)
		return
	}

	ctx = trCtx.Context()

	//keep previous client id
	ctx = authn.NewClientContext(ctx, clientId)
	ctx = authn.NewUserContext(ctx, authn.NewUserInfo(uid))

	//set auth token
	//use token mgr
	token, err := tokenManager.GetOrGenerateToken(ctx, &conf2.Client{
		ClientId:     apiClient.ClientId,
		ClientSecret: apiClient.ClientSecret,
	})
	if err != nil {
		log.Errorf("%s", err)
		w.WriteHeader(500)
		return
	}
	//replace with internal token
	r.Header().Set(jwt.AuthorizationHeader, fmt.Sprintf("%s %s", jwt.BearerTokenType, token))

	headers := api.HeaderCarrier(map[string]string{})
	//inject header
	for _, contributor := range apiOpt.Propagators {
		//do not handle error
		contributor.Inject(ctx, headers)
	}
	for k, v := range headers {
		log.Infof("set header: %s value: %s", k, v)
		r.Header().Set(k, v)
	}
	//continue request
	return
}
