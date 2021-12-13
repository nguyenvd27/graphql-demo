package resolver

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

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
