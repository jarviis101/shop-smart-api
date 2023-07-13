package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/controller/http/validator"
	"shop-smart-api/internal/usecase"
	"shop-smart-api/pkg"
)

type otpRouteManager struct {
	group        *echo.Group
	validator    *validator.Validator
	userUseCase  usecase.UserUseCase
	otpUseCase   usecase.OTPUseCase
	serverConfig pkg.Server
}

func CreateOTPRouterManager(
	g *echo.Group,
	v *validator.Validator,
	uc usecase.UserUseCase,
	oc usecase.OTPUseCase,
	sc pkg.Server,
) RouteManager {
	return &otpRouteManager{g, v, uc, oc, sc}
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

	verifyRequest := &types.VerifyOTPRequest{}
	if err := c.Bind(verifyRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := r.validator.Validate(verifyRequest); err != nil {
		return err
	}

	currentUser := c.Get(middleware.Key).(string)
	user, err := r.userUseCase.Get(ctx, currentUser)
	if err != nil {
		return err
	}

	if err := r.otpUseCase.Verify(ctx, user, verifyRequest.Code); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := r.userUseCase.Authenticate(user)
	if err != nil {
		return err
	}
	response := &types.TokenResponse{Token: token}

	return c.JSON(http.StatusOK, response)
}
