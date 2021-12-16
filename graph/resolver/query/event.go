package query

import (
	"context"

	"github.com/nguyenvd27/graphql-test/model"
)

func (r *Resolver) Events(ctx context.Context) ([]*model.Event, error) {
	events, err := r.Interactor.EventUsecase().GetAllEventsUsecase()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *Resolver) Event(ctx context.Context, id string) (*model.Event, error) {
	event, err := r.Interactor.EventUsecase().GetAEventUsecase(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}
