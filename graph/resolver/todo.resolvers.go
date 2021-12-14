package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nguyenvd27/graphql-test/graph/generated"
	graphqlmodel "github.com/nguyenvd27/graphql-test/graph/model"
	"github.com/nguyenvd27/graphql-test/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input graphqlmodel.NewTodo) (*model.Todo, error) {
	return &model.Todo{}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{}, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	return &model.Todo{}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
