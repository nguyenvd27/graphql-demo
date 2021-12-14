package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/nguyenvd27/graphql-test/graph/generated"
	graphqlmodel "github.com/nguyenvd27/graphql-test/graph/model"
	"github.com/nguyenvd27/graphql-test/model"
)

func (r *eventResolver) User(ctx context.Context, obj *model.Event) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "event user " + obj.UserID}, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input graphqlmodel.NewEvent) (*model.Event, error) {
	newEvent, err := r.Interactor.EventUsecase().CreateEventUsecase(fmt.Sprintf("T%d", rand.Int()), input.Title, input.Content, input.UserID)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	events, err := r.Interactor.EventUsecase().GetAllEventsUsecase()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	event, err := r.Interactor.EventUsecase().GetAEventUsecase(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
