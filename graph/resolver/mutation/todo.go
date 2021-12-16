package mutation

import (
	"context"
	"fmt"
	"math/rand"

	graphqlmodel "github.com/nguyenvd27/graphql-test/graph/model"
	"github.com/nguyenvd27/graphql-test/model"
)

func (r *Resolver) CreateTodo(ctx context.Context, input graphqlmodel.NewTodo) (*model.Todo, error) {
	newTodo, err := r.Interactor.TodoUsecase().CreateTodoUsecase(fmt.Sprintf("T%d", rand.Int()), input.Text, input.UserID, false)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}
