package controller

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shop-smart-api/internal/controller/graphql/directives"
	"shop-smart-api/internal/controller/graphql/graph"
	"shop-smart-api/internal/controller/graphql/transformers"
	http_context "shop-smart-api/internal/controller/http/context"
	"shop-smart-api/internal/controller/http/router"
	http_validator "shop-smart-api/internal/controller/http/validator"
	"shop-smart-api/internal/service"
	"shop-smart-api/pkg"
)

type http struct {
	serverConfig    pkg.Server
	userUseCase     service.UserUseCase
	otpUseCase      service.OTPUseCase
	userTransformer transformers.UserTransformer
	validator       *http_validator.Validator
	echo            *echo.Echo
}

func CreateServer(
	sc pkg.Server,
	u service.UserUseCase,
	o service.OTPUseCase,
) Server {
	v := http_validator.CreateValidator(validator.New())
	e := echo.New()
	e.Validator = v
	e.Use(http_context.EchoContextToContextMiddleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	baseTransformer := transformers.CreateBaseTransformer()
	ut := transformers.CreateUserTransformer(baseTransformer)

	return &http{
		serverConfig:    sc,
		userUseCase:     u,
		otpUseCase:      o,
		userTransformer: ut,
		validator:       v,
		echo:            e,
	}
}

func (h *http) RunServer() error {
	h.appendRestRoutes(h.echo)
	h.appendGraphqlRoutes(h.echo)

	host := fmt.Sprintf(":%s", h.serverConfig.Port)
	return h.echo.Start(host)
}

func (h *http) appendRestRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")
	authRouter := router.CreateAuthRouterManager(apiGroup, h.validator, h.userUseCase, h.otpUseCase)
	authRouter.PopulateRoutes()

	otpGroup := apiGroup.Group("/otp")
	otpRouter := router.CreateOTPRouterManager(otpGroup, h.validator, h.userUseCase, h.otpUseCase, h.serverConfig)
	otpRouter.PopulateRoutes()
}

func (h *http) appendGraphqlRoutes(e *echo.Echo) {
	resolver := graph.CreateResolver(h.userUseCase, h.userTransformer)
	c := graph.Config{Resolvers: resolver}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	pg := playground.Handler("GraphQL playground", "/query")

	graphqlRouter := router.CreateGraphqlRouterManager(e.Group(""), srv, pg, h.serverConfig)

	graphqlRouter.PopulateRoutes()
}
