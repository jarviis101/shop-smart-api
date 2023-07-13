package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/usecase"
	"shop-smart-api/pkg"
)

type otpRouteManager struct {
	userUseCase  usecase.UserUseCase
	group        *echo.Group
	serverConfig pkg.Server
}

func CreateOTPRouterManager(uc usecase.UserUseCase, g *echo.Group, sc pkg.Server) RouteManager {
	return &otpRouteManager{uc, g, sc}
}

func (r *otpRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/verify", r.verify, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
	r.group.Add("POST", "/send", r.send, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
}

func (r *otpRouteManager) send(c echo.Context) error {
	_ = c.Get(middleware.Key).(string)
	return c.JSON(http.StatusOK, "send")
}

func (r *otpRouteManager) verify(c echo.Context) error {
	ctx := c.Request().Context()
	currentUser := c.Get(middleware.Key).(string)
	user, err := r.userUseCase.Get(ctx, currentUser)
	if err != nil {
		return err
	}

	// TODO Check OTP
	token, err := r.userUseCase.Authenticate(user)
	if err != nil {
		return err
	}
	response := &types.TokenResponse{Token: token}

	return c.JSON(http.StatusOK, response)
}
