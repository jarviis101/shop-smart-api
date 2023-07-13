package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/controller/http/validator"
	"shop-smart-api/internal/usecase"
)

type authRouteManager struct {
	group       *echo.Group
	validator   *validator.Validator
	userUseCase usecase.UserUseCase
	otpUseCase  usecase.OTPUseCase
}

func CreateAuthRouterManager(
	g *echo.Group,
	v *validator.Validator,
	u usecase.UserUseCase,
	o usecase.OTPUseCase,
) RouteManager {
	return &authRouteManager{g, v, u, o}
}

func (r *authRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/auth", r.auth)
}

func (r *authRouteManager) auth(c echo.Context) error {
	ctx := c.Request().Context()
	u := &types.LoginUserRequest{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := r.validator.Validate(u); err != nil {
		return err
	}

	token, err := r.userUseCase.PreAuthenticate(ctx, u.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: In a future rework this
	user, err := r.userUseCase.GetByPhone(ctx, u.Phone)
	if err != nil {
		return err
	}

	if err := r.otpUseCase.Send(ctx, user); err != nil {
		return err
	}

	response := &types.TokenResponse{Token: token}
	return c.JSON(http.StatusOK, response)
}
