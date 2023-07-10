package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/pkg/jwt"
)

const (
	header                  = "Authorization"
	excludedStringFromToken = "Bearer "
)

var claims *jwt.UserClaims

func OTPAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := receiveClaims(c, secret)
			if err != nil {
				return err
			}

			if claims.IsFully {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			if err := next(c); err != nil {
				return err
			}

			return nil
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

			claims = verifiedClaims

			if !verifiedClaims.IsFully {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}

func GetClaims() *jwt.UserClaims {
	// TODO: provide claims must be with another way
	return claims
}

func receiveClaims(c echo.Context, secret string) (*jwt.UserClaims, error) {
	jwtManager := jwt.CreateManager(secret)
	token := c.Request().Header.Get(header)
	if token == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "not authorized")
	}
	accessToken := token[len(excludedStringFromToken):]
	verifiedClaims, err := jwtManager.Verify(accessToken)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return verifiedClaims, nil
}
