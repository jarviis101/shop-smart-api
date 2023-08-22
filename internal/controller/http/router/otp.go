package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/controller/http/validator"
	"shop-smart-api/internal/service"
	"shop-smart-api/pkg"
)

type otpRouteManager struct {
	group        *echo.Group
	validator    *validator.Validator
	userUseCase  service.UserService
	otpUseCase   service.OTPService
	serverConfig pkg.Server
}

func CreateOTPRouterManager(
	g *echo.Group,
	v *validator.Validator,
	uc service.UserService,
	oc service.OTPService,
	sc pkg.Server,
) RouteManager {
	return &otpRouteManager{g, v, uc, oc, sc}
}

func (r *otpRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/verify", r.verify, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
	r.group.Add("POST", "/send", r.send, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
}

func (r *otpRouteManager) send(c echo.Context) error {
	currentUser := c.Get(middleware.CurrentUserKey).(int64)

	user, err := r.userUseCase.Get(currentUser)
	if err != nil {
		return err
	}

	t := types.Phone

	if err := r.otpUseCase.Send(user, &t); err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, "")
}

func (r *otpRouteManager) verify(c echo.Context) error {
	verifyRequest := &types.VerifyOTPRequest{}
	if err := c.Bind(verifyRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := r.validator.Validate(verifyRequest); err != nil {
		return err
	}

	currentUser := c.Get(middleware.CurrentUserKey).(int64)
	user, err := r.userUseCase.Get(currentUser)
	if err != nil {
		return err
	}

	if err := r.otpUseCase.Verify(user, verifyRequest.Code); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := r.userUseCase.Authenticate(user)
	if err != nil {
		return err
	}
	response := &types.TokenResponse{Token: token}

	return c.JSON(http.StatusOK, response)
}
