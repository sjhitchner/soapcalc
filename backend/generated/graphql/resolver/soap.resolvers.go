package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphql1 "github.com/sjhitchner/soapcalc/backend/generated/graphql/models"
	"github.com/sjhitchner/soapcalc/backend/generated/graphql/server"
)

func (r *mutationResolver) CalculateRecipe(ctx context.Context, input graphql1.CreateRecipe) (*graphql1.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Additive(ctx context.Context, id string) (*graphql1.Additive, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListAdditives(ctx context.Context, first *int, after *string) ([]*graphql1.Additive, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Lye(ctx context.Context, id string) (*graphql1.Lye, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListLye(ctx context.Context, first *int, after *string) ([]*graphql1.Lye, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Lipid(ctx context.Context, id string) (*graphql1.Lipid, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListLipids(ctx context.Context, first *int, after *string) ([]*graphql1.Lipid, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Fragrance(ctx context.Context, id string) (*graphql1.Fragrance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListFragrances(ctx context.Context, first *int, after *string) ([]*graphql1.Fragrance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Recipe(ctx context.Context, id string) (*graphql1.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListRecipes(ctx context.Context, first *int, after *string) ([]*graphql1.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchRecipes(ctx context.Context, search string, first *int, after *string) ([]*graphql1.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
