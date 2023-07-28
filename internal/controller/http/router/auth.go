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
	r.group.Add("POST", "/auth", r.auth)
}

func (r *authRouteManager) auth(c echo.Context) error {
	authRequest := &types.AuthUserRequest{}
	if err := c.Bind(authRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := r.validator.Validate(authRequest); err != nil {
		return err
	}

	token, err := r.userUseCase.PreAuthenticate(authRequest.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: In a future rework this
	user, err := r.userUseCase.GetByPhone(authRequest.Phone)
	if err != nil {
		return err
	}

	if err := r.otpUseCase.Send(user); err != nil {
		return err
	}

	response := &types.TokenResponse{Token: token}
	return c.JSON(http.StatusOK, response)
}
