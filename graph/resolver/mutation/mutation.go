package mutation

import "github.com/nguyenvd27/graphql-test/interactor"

type Resolver struct {
	Interactor interactor.Interactor
}

func NewMutationResolver(interactor interactor.Interactor) *Resolver {
	return &Resolver{
		Interactor: interactor,
	}
}
