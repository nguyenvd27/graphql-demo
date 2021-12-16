package query

import "github.com/nguyenvd27/graphql-test/interactor"

type Resolver struct {
	Interactor interactor.Interactor
}

func NewQueryResolver(interactor interactor.Interactor) *Resolver {
	return &Resolver{
		Interactor: interactor,
	}
}
