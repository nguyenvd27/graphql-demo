package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/nguyenvd27/graphql-test/interactor"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Interactor interactor.Interactor
}
