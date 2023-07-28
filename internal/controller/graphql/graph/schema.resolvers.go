package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"shop-smart-api/internal/controller/graphql/directives"
	"shop-smart-api/internal/controller/graphql/graph/model"
)

// UpdatePersonalInfo is the resolver for the updatePersonalInfo field.
func (r *mutationResolver) UpdatePersonalInfo(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	currentUser := ctx.Value(directives.AuthKey(directives.Key)).(int64)
	user, _ := r.userService.Get(currentUser)

	updatedUser, err := r.userService.Update(user, input.FirstName, input.LastName, input.MiddleName)
	if err != nil {
		return nil, &gqlerror.Error{Message: "something went wrong"}
	}

	return r.userTransformer.TransformToModel(updatedUser), nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	currentUser := ctx.Value(directives.AuthKey(directives.Key)).(int64)
	user, _ := r.userService.Get(currentUser)

	return r.userTransformer.TransformToModel(user), nil
}

// GetOrganization is the resolver for the getOrganization field.
func (r *queryResolver) GetOrganization(ctx context.Context) (*model.Organization, error) {
	currentUser := ctx.Value(directives.AuthKey(directives.Key)).(int64)
	user, _ := r.userService.Get(currentUser)

	if user.OrganizationID == nil {
		return nil, &gqlerror.Error{Message: "organization not found"}
	}

	organization, err := r.organizationService.Get(*user.OrganizationID)
	if err != nil {
		return nil, &gqlerror.Error{Message: "organization not found"}
	}

	return r.organizationTransformer.TransformToModel(organization), nil
}

// GetOrganizationUsers is the resolver for the getOrganizationUsers field.
func (r *queryResolver) GetOrganizationUsers(ctx context.Context) ([]*model.User, error) {
	currentUser := ctx.Value(directives.AuthKey(directives.Key)).(int64)
	user, _ := r.userService.Get(currentUser)

	if user.OrganizationID == nil {
		return nil, &gqlerror.Error{Message: "organization not found"}
	}

	users, err := r.userService.GetByOrganization(*user.OrganizationID)
	if err != nil {
		return nil, &gqlerror.Error{Message: err.Error()}
	}

	return r.userTransformer.TransformManyToModel(users), nil
}

// GetTransactions is the resolver for the getTransactions field.
func (r *queryResolver) GetTransactions(ctx context.Context) ([]*model.Transaction, error) {
	currentUser := ctx.Value(directives.AuthKey(directives.Key)).(int64)
	user, _ := r.userService.Get(currentUser)

	transactions, err := r.transactionService.GetTransactions(user)
	if err != nil {
		return nil, err
	}

	return r.transactionTransformer.TransformManyToModel(transactions), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
