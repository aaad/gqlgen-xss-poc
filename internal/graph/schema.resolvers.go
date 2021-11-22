package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aaad/gqlgen-xss-poc/internal/graph/generated"
	"github.com/aaad/gqlgen-xss-poc/internal/graph/model"
)

func (r *queryResolver) Test(ctx context.Context) (*model.TestResult, error) {
	return &model.TestResult{
		Description: "Description <h1>This page has a XSS vulnerability.</h1>",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
