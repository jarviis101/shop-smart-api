package graph

import (
	"shop-smart-api/internal/controller/graphql/transformers"
	"shop-smart-api/internal/service"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase     service.UserUseCase
	userTransformer transformers.UserTransformer
}

func CreateResolver(u service.UserUseCase, ut transformers.UserTransformer) *Resolver {
	return &Resolver{u, ut}
}
