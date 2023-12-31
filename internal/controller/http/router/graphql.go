package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/pkg"
)

type graphqlRouterManager struct {
	group  *echo.Group
	server *handler.Server
	pg     http.HandlerFunc
	sc     pkg.Server
}

func CreateGraphqlRouterManager(g *echo.Group, s *handler.Server, pg http.HandlerFunc, sc pkg.Server) RouteManager {
	return &graphqlRouterManager{g, s, pg, sc}
}

func (r *graphqlRouterManager) PopulateRoutes() {
	r.group.Add("GET", "/graphql", r.playground)
	r.group.Add("POST", "/query", r.query, middleware.AuthMiddleware(r.sc.Secret))
}

func (r *graphqlRouterManager) query(c echo.Context) error {
	r.server.ServeHTTP(c.Response(), c.Request())
	return nil
}

func (r *graphqlRouterManager) playground(c echo.Context) error {
	r.pg.ServeHTTP(c.Response(), c.Request())
	return nil
}
