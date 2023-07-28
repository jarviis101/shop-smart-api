package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"shop-smart-api/internal/controller/graphql/directives"
	"shop-smart-api/internal/controller/graphql/graph/model"
)

// UpdatePersonalInfo is the resolver for the updatePersonalInfo field.
func (r *mutationResolver) UpdatePersonalInfo(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	currentUser := ctx.Value(directives.AuthKey(directives.Key)).(int64)
	user, err := r.userUseCase.Get(currentUser)
	if err != nil {
		return nil, err
	}

	updatedUser, err := r.userUseCase.Update(user, input.FirstName, input.LastName, input.MiddleName)
	if err != nil {
		return nil, err
	}

	return r.userTransformer.TransformToModel(updatedUser), nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	currentUser := ctx.Value(directives.AuthKey(directives.Key)).(int64)
	user, err := r.userUseCase.Get(currentUser)
	if err != nil {
		return nil, err
	}

	return r.userTransformer.TransformToModel(user), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }