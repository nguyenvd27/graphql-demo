package query

import (
	"context"

	"github.com/nguyenvd27/graphql-test/model"
)

func (r *Resolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.Interactor.TodoUsecase().GetAllTodosUsecase()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *Resolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	todo, err := r.Interactor.TodoUsecase().GetATodoUsecase(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
