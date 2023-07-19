package context

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

type contextType string

const Key = "context"

func EchoContextToContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), contextType(Key), c)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func EchoContextFromContext(ctx context.Context) (echo.Context, error) {
	contextValue := ctx.Value(contextType(Key))
	if contextValue == nil {
		return nil, fmt.Errorf("could not retrieve echo.Context")
	}

	echoContext, ok := contextValue.(echo.Context)
	if !ok {
		return nil, fmt.Errorf("echo.Context has wrong type")
	}

	return echoContext, nil
}
