package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sjhitchner/soapcalc/graph/generated"
	"github.com/sjhitchner/soapcalc/graph/model"
)

func (r *mutationResolver) CalculateRecipe(ctx context.Context, input model.CreateRecipe) (*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
