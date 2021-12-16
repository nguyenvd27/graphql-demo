package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/nguyenvd27/graphql-test/graph/generated"
	"github.com/nguyenvd27/graphql-test/graph/resolver/event"
	"github.com/nguyenvd27/graphql-test/graph/resolver/mutation"
	"github.com/nguyenvd27/graphql-test/graph/resolver/query"
	"github.com/nguyenvd27/graphql-test/graph/resolver/todo"
	"github.com/nguyenvd27/graphql-test/interactor"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Interactor interactor.Interactor
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return mutation.NewMutationResolver(r.Interactor)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return query.NewQueryResolver(r.Interactor)
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver {
	return event.NewEventResolver(r.Interactor)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver {
	return todo.NewTodoResolver(r.Interactor)
}
