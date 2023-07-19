package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/pkg/jwt"
)

const (
	CurrentUserKey = "currentUser"
	header         = "Authorization"
	excludedString = "Bearer "
)

func OTPAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			verifiedClaims, err := receiveClaims(c, secret)
			if err != nil {
				return err
			}

			if verifiedClaims.IsFully {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			c.Set(CurrentUserKey, verifiedClaims.UserId)
			return next(c)
		}
	}
}

func AuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			verifiedClaims, err := receiveClaims(c, secret)
			if err != nil {
				return err
			}

			if !verifiedClaims.IsFully {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			c.Set(CurrentUserKey, verifiedClaims.UserId)
			return next(c)
		}
	}
}

func receiveClaims(c echo.Context, secret string) (*jwt.UserClaims, error) {
	jwtManager := jwt.CreateManager(secret)
	token := c.Request().Header.Get(header)
	if token == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "not authorized")
	}
	accessToken := token[len(excludedString):]
	verifiedClaims, err := jwtManager.Verify(accessToken)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return verifiedClaims, nil
}
