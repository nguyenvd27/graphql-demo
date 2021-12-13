package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/nguyenvd27/graphql-test/graph/generated"
	"github.com/nguyenvd27/graphql-test/graph/model"
)

func (r *eventResolver) User(ctx context.Context, obj *model.Event) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "event user " + obj.UserID}, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	event := &model.Event{
		Title:   input.Title,
		Content: input.Content,
		ID:      fmt.Sprintf("T%d", rand.Int()),
		UserID:  input.UserID,
	}
	r.events = append(r.events, event)
	return event, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	for i := 0; i < len(r.todos); i++ {
		if id == r.todos[i].ID {
			return r.todos[i], nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	return r.events, nil
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	for i := 0; i < len(r.events); i++ {
		if id == r.events[i].ID {
			return r.events[i], nil
		}
	}
	return nil, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type eventResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
