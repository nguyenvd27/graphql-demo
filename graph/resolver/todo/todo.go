package todo

import (
	"context"

	"github.com/nguyenvd27/graphql-test/interactor"
	"github.com/nguyenvd27/graphql-test/model"
)

type Resolver struct {
	Interactor interactor.Interactor
}

func NewTodoResolver(interactor interactor.Interactor) *Resolver {
	return &Resolver{
		Interactor: interactor,
	}
}

func (r *Resolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "todo user " + obj.UserID}, nil
}
