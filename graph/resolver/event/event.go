package event

import (
	"context"

	"github.com/nguyenvd27/graphql-test/interactor"
	"github.com/nguyenvd27/graphql-test/model"
)

type Resolver struct {
	Interactor interactor.Interactor
}

func NewEventResolver(interactor interactor.Interactor) *Resolver {
	return &Resolver{
		Interactor: interactor,
	}
}

func (r *Resolver) User(ctx context.Context, obj *model.Event) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "event user " + obj.UserID}, nil
}
