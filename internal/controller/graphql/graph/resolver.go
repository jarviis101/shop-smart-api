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
	userService             service.UserService
	userTransformer         transformers.UserTransformer
	organizationService     service.OrganizationService
	organizationTransformer transformers.OrganizationTransformer
	transactionService      service.TransactionService
	transactionTransformer  transformers.TransactionTransformer
}

func CreateResolver(
	u service.UserService,
	ut transformers.UserTransformer,
	o service.OrganizationService,
	ot transformers.OrganizationTransformer,
	ts service.TransactionService,
	tt transformers.TransactionTransformer,
) *Resolver {
	return &Resolver{
		userService:             u,
		userTransformer:         ut,
		organizationService:     o,
		organizationTransformer: ot,
		transactionService:      ts,
		transactionTransformer:  tt,
	}
}
