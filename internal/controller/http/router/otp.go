package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/pkg"
)

type otpRouteManager struct {
	group        *echo.Group
	serverConfig pkg.Server
}

func CreateOTPRouterManager(g *echo.Group, sc pkg.Server) RouteManager {
	return &otpRouteManager{g, sc}
}

func (r *otpRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/verify", r.verify, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
	r.group.Add("POST", "/send", r.send, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
}

func (r *otpRouteManager) send(c echo.Context) error {

	return c.JSON(http.StatusOK, "send")
}

func (r *otpRouteManager) verify(c echo.Context) error {
	return c.JSON(http.StatusOK, "verify")
}
