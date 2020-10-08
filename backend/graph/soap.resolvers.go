package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/sjhitchner/soapcalc/backend/graph/generated"
	model1 "github.com/sjhitchner/soapcalc/backend/graph/model"
)

func (r *mutationResolver) CalculateRecipe(ctx context.Context, input model1.CreateRecipe) (*model1.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
