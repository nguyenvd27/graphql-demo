package interactor

import "github.com/nguyenvd27/graphql-test/usecase"

type Interactor interface {
	// Resolver
	EventUsecase() usecase.EventUsecase
}

type interactor struct {
	reg Registry
}

func NewInteractor(reg Registry) Interactor {
	return &interactor{
		reg: reg,
	}
}

func (i *interactor) EventUsecase() usecase.EventUsecase {
	return usecase.NewEventUsecase(i.reg.EventRepository())
}
