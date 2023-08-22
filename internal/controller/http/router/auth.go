package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/controller/http/validator"
	"shop-smart-api/internal/service"
)

type authRouteManager struct {
	group       *echo.Group
	validator   *validator.Validator
	userUseCase service.UserService
	otpUseCase  service.OTPService
}

func CreateAuthRouterManager(
	g *echo.Group,
	v *validator.Validator,
	uc service.UserService,
	oc service.OTPService,
) RouteManager {
	return &authRouteManager{g, v, uc, oc}
}

func (r *authRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/auth", r.sendCode)
}

func (r *authRouteManager) sendCode(c echo.Context) error {
	authRequest, err := r.resolveAuthRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	channel, err := types.ResolveByChannel(authRequest.Channel)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, token, err := r.userUseCase.ProvideOrCreate(authRequest.Resource, channel)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := r.otpUseCase.Send(user, channel); err != nil {
		return err
	}

	response := &types.TokenResponse{Token: token}
	return c.JSON(http.StatusOK, response)
}

func (r *authRouteManager) resolveAuthRequest(c echo.Context) (*types.AuthUserRequest, error) {
	authRequest := &types.AuthUserRequest{}
	if err := c.Bind(authRequest); err != nil {
		return nil, err
	}

	if err := r.validator.Validate(authRequest); err != nil {
		return nil, err
	}

	return authRequest, nil
}
