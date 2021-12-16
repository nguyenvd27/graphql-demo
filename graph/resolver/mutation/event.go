package mutation

import (
	"context"
	"fmt"
	"math/rand"

	graphqlmodel "github.com/nguyenvd27/graphql-test/graph/model"
	"github.com/nguyenvd27/graphql-test/model"
)

func (r *Resolver) CreateEvent(ctx context.Context, input graphqlmodel.NewEvent) (*model.Event, error) {
	newEvent, err := r.Interactor.EventUsecase().CreateEventUsecase(fmt.Sprintf("T%d", rand.Int()), input.Title, input.Content, input.UserID)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}
