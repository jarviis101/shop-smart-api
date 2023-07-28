package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	http_context "shop-smart-api/internal/controller/http/context"
	"shop-smart-api/internal/controller/http/middleware"
)

type AuthKey string

const Key = "auth"

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	echoContext, err := http_context.EchoContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	currentUser, ok := echoContext.Get(middleware.CurrentUserKey).(int64)
	if !ok {
		return nil, &gqlerror.Error{Message: "Access Denied"}
	}

	c := context.WithValue(ctx, AuthKey(Key), currentUser)
	return next(c)
}
