// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-eventbus"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authz"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/casbin"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/dal"
	"github.com/goxiaoy/go-saas-kit/pkg/gorm"
	"github.com/goxiaoy/go-saas-kit/pkg/job"
	"github.com/goxiaoy/go-saas-kit/pkg/redis"
	server2 "github.com/goxiaoy/go-saas-kit/pkg/server"
	"github.com/goxiaoy/go-saas-kit/pkg/uow"
	api2 "github.com/goxiaoy/go-saas-kit/saas/api"
	api3 "github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas-kit/user/private/biz"
	conf2 "github.com/goxiaoy/go-saas-kit/user/private/conf"
	"github.com/goxiaoy/go-saas-kit/user/private/data"
	"github.com/goxiaoy/go-saas-kit/user/private/server"
	"github.com/goxiaoy/go-saas-kit/user/private/service"
	http2 "github.com/goxiaoy/go-saas-kit/user/private/service/http"
	"github.com/goxiaoy/go-saas/common/http"
)

import (
	_ "github.com/goxiaoy/go-saas-kit/saas/api"
	_ "github.com/goxiaoy/go-saas-kit/sys/api"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, userConf *conf2.UserConf, confData *conf.Data, logger log.Logger, webMultiTenancyOption *http.WebMultiTenancyOption, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	config := _wireConfigValue
	dbCache, cleanup := gorm.NewDbCache(confData)
	manager := uow.NewUowManager(config, dbCache)
	option := api.NewDefaultOption(logger)
	clientName := _wireClientNameValue
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer, logger)
	grpcConn, cleanup2 := api2.NewGrpcConn(clientName, services, option, inMemoryTokenManager, logger, arg...)
	tenantServiceServer := api2.NewTenantGrpcClient(grpcConn)
	tenantStore := api2.NewTenantStore(tenantServiceServer)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	connStrResolver := dal.NewConnStrResolver(confData, tenantStore)
	dbProvider := gorm.NewDbProvider(dbCache, connStrResolver, confData)
	dataData, cleanup3, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	passwordHasher := biz.NewPasswordHasher()
	userValidator := biz.NewUserValidator()
	passwordValidator := biz.NewPasswordValidator(userConf)
	lookupNormalizer := biz.NewLookupNormalizer()
	userTokenRepo := data.NewUserTokenRepo(dataData)
	refreshTokenRepo := data.NewRefreshTokenRepo(dataData)
	userTenantRepo := data.NewUserTenantRepo(dataData)
	connName := _wireConnNameValue
	client := dal.NewRedis(confData, connName)
	emailTokenProvider := biz.NewEmailTokenProvider(client)
	phoneTokenProvider := biz.NewPhoneTokenProvider(client)
	cache := redis.NewCache(client)
	twoStepTokenProvider := biz.NewTwoStepTokenProvider(cache)
	userManager := biz.NewUserManager(userConf, userRepo, passwordHasher, userValidator, passwordValidator, lookupNormalizer, userTokenRepo, refreshTokenRepo, userTenantRepo, emailTokenProvider, phoneTokenProvider, twoStepTokenProvider, logger)
	eventBus := _wireEventBusValue
	roleRepo := data.NewRoleRepo(dataData, eventBus)
	roleManager := biz.NewRoleManager(roleRepo, lookupNormalizer)
	enforcerProvider, err := data.NewEnforcerProvider(logger, dbProvider)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	permissionService := casbin.NewPermissionService(enforcerProvider)
	userRoleContributor := service.NewUserRoleContributor(userRepo)
	authzOption := server.NewAuthorizationOption(userRoleContributor)
	subjectResolverImpl := authz.NewSubjectResolver(authzOption)
	defaultAuthorizationService := authz.NewDefaultAuthorizationService(permissionService, subjectResolverImpl, logger)
	factory := dal.NewBlobFactory(confData)
	trustedContextValidator := api.NewClientTrustedContextValidator()
	userService := service.NewUserService(userManager, roleManager, defaultAuthorizationService, factory, trustedContextValidator, logger)
	userServiceServer := service.NewUserServiceServer(userService)
	userTenantContributor := api3.NewUserTenantContributor(userServiceServer)
	lazyClient := dal.NewEmailer(confData)
	emailSender := biz.NewEmailSender(lazyClient, confData)
	authService := service.NewAuthService(userManager, roleManager, tokenizer, tokenizerConfig, passwordValidator, refreshTokenRepo, emailSender, security, defaultAuthorizationService, trustedContextValidator, logger)
	authServer := service.NewAuthServiceServer(authService)
	refreshTokenProvider := api3.NewRefreshProvider(authServer, logger)
	userSettingRepo := data.NewUserSettingRepo(dataData, eventBus)
	userAddressRepo := data.NewUserAddrRepo(dataData, eventBus)
	accountService := service.NewAccountService(userManager, factory, tenantServiceServer, userSettingRepo, userAddressRepo, lookupNormalizer)
	roleService := service.NewRoleServiceService(roleManager, defaultAuthorizationService, permissionService)
	servicePermissionService := service.NewPermissionService(defaultAuthorizationService, permissionService, subjectResolverImpl, trustedContextValidator)
	signInManager := biz.NewSignInManager(userManager, security)
	apiClient := server.NewHydra(security)
	auth := http2.NewAuth(decodeRequestFunc, userManager, logger, signInManager, apiClient)
	httpServerRegister := service.NewHttpServerRegister(userService, encodeResponseFunc, encodeErrorFunc, accountService, authService, roleService, servicePermissionService, auth, confData, factory)
	httpServer := server.NewHTTPServer(services, security, tokenizer, manager, webMultiTenancyOption, option, tenantStore, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, logger, userTenantContributor, trustedContextValidator, refreshTokenProvider, httpServerRegister)
	grpcServerRegister := service.NewGrpcServerRegister(userService, accountService, authService, roleService, servicePermissionService)
	grpcServer := server.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, option, logger, trustedContextValidator, userTenantContributor, grpcServerRegister)
	redisConnOpt := job.NewAsynqClientOpt(client)
	migrate := data.NewMigrate(dataData)
	roleSeed := biz.NewRoleSeed(roleManager, permissionService)
	userSeed := biz.NewUserSeed(userManager, roleManager)
	permissionSeeder := biz.NewPermissionSeeder(permissionService, permissionService, roleManager)
	seeding := server.NewSeeding(manager, migrate, roleSeed, userSeed, permissionSeeder)
	seeder := server.NewSeeder(seeding)
	sender, cleanup4, err := dal.NewEventSender(confData, logger, connName)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userMigrationTaskHandler := biz.NewUserMigrationTaskHandler(seeder, sender)
	jobServer := server.NewJobServer(redisConnOpt, userMigrationTaskHandler)
	asynqClient, cleanup5 := job.NewAsynqClient(redisConnOpt)
	tenantSeedEventHandler := biz.NewTenantSeedEventHandler(asynqClient)
	userEventHandler := biz.NewRemoteEventHandler(logger, manager, tenantSeedEventHandler)
	handler := server.NewEventHandler(userEventHandler)
	receiver, cleanup6, err := dal.NewRemoteEventReceiver(confData, logger, handler, connName)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := newApp(userConf, logger, httpServer, grpcServer, jobServer, seeder, receiver)
	return app, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireConfigValue             = dal.UowCfg
	_wireClientNameValue         = server.ClientName
	_wireDecodeRequestFuncValue  = server2.ReqDecode
	_wireEncodeResponseFuncValue = server2.ResEncoder
	_wireEncodeErrorFuncValue    = server2.ErrEncoder
	_wireConnNameValue           = biz.ConnName
	_wireEventBusValue           = eventbus.Default
)
