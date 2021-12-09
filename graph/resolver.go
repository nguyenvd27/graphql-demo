package graph

import "github.com/nguyenvd27/graphql-test/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos  []*model.Todo
	events []*model.Event
}
