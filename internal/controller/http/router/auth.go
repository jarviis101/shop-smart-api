package router

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/controller/http/validator"
	"shop-smart-api/internal/usecase"
)

type authRouteManager struct {
	group     *echo.Group
	validator *validator.Validator
	useCase   usecase.UserUseCase
}

func CreateAuthRouterManager(g *echo.Group, v *validator.Validator, u usecase.UserUseCase) RouteManager {
	return &authRouteManager{g, v, u}
}

func (r *authRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/auth", r.auth)
}

func (r *authRouteManager) auth(c echo.Context) error {
	u := &types.LoginUserRequest{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := r.validator.Validate(u); err != nil {
		return err
	}

	token, err := r.useCase.PreAuthenticate(context.Background(), u.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response := &types.TokenResponse{Token: token}

	// TODO: send otp
	log.Println("Send OTP")

	return c.JSON(http.StatusOK, response)
}
