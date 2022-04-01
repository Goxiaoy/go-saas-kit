// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authz"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	gorm2 "github.com/goxiaoy/go-saas-kit/pkg/gorm"
	server2 "github.com/goxiaoy/go-saas-kit/pkg/server"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/saas/private/biz"
	conf2 "github.com/goxiaoy/go-saas-kit/saas/private/conf"
	"github.com/goxiaoy/go-saas-kit/saas/private/data"
	"github.com/goxiaoy/go-saas-kit/saas/private/server"
	"github.com/goxiaoy/go-saas-kit/saas/private/service"
	api2 "github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas-kit/user/remote"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/uow"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, confData *conf2.Data, logger log.Logger, config *uow.Config, gormConfig *gorm.Config, webMultiTenancyOption *http.WebMultiTenancyOption, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	sender, cleanup, err := data.NewEventSender(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	eventBus, cleanup2, err := data.NewEventbus(sender)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tenantRepo := data.NewTenantRepo(eventBus)
	tenantStore := data.NewTenantStore(tenantRepo)
	dbOpener, cleanup3 := gorm2.NewDbOpener()
	manager := uow2.NewUowManager(gormConfig, config, dbOpener)
	tenantUseCase := biz.NewTenantUserCase(tenantRepo)
	clientName := _wireClientNameValue
	option := api.NewDefaultOption(logger)
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer, logger)
	grpcConn, cleanup4 := api2.NewGrpcConn(clientName, services, option, inMemoryTokenManager, logger, arg...)
	permissionServiceClient := api2.NewPermissionGrpcClient(grpcConn)
	permissionChecker := remote.NewRemotePermissionChecker(permissionServiceClient)
	authzOption := service.NewAuthorizationOption()
	subjectResolverImpl := authz.NewSubjectResolver(authzOption)
	defaultAuthorizationService := authz.NewDefaultAuthorizationService(permissionChecker, subjectResolverImpl, logger)
	trustedContextValidator := api.NewClientTrustedContextValidator()
	factory := data.NewBlobFactory(confData)
	tenantService := service.NewTenantService(tenantUseCase, defaultAuthorizationService, trustedContextValidator, factory)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	userServiceClient := api2.NewUserGrpcClient(grpcConn)
	userTenantContributor := remote.NewUserTenantContributor(userServiceClient)
	httpServer := server.NewHTTPServer(services, security, tokenizer, tenantStore, manager, tenantService, webMultiTenancyOption, option, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, factory, confData, logger, trustedContextValidator, userTenantContributor)
	grpcServer := server.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, option, tenantService, userTenantContributor, trustedContextValidator, logger)
	connStrResolver := data.NewConnStrResolver(confData, tenantStore)
	dbProvider := gorm2.NewDbProvider(connStrResolver, gormConfig, dbOpener)
	dataData, cleanup5, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	migrate := data.NewMigrate(dataData)
	seeder := server.NewSeeder(confData, manager, migrate)
	app := newApp(logger, httpServer, grpcServer, seeder)
	return app, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireClientNameValue         = server.ClientName
	_wireDecodeRequestFuncValue  = server2.ReqDecode
	_wireEncodeResponseFuncValue = server2.ResEncoder
	_wireEncodeErrorFuncValue    = server2.ErrEncoder
)
