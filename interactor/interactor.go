package interactor

import "github.com/nguyenvd27/graphql-test/usecase"

type Interactor interface {
	// Resolver
	EventUsecase() usecase.EventUsecase
	TodoUsecase() usecase.TodoUsecase
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

func (i *interactor) TodoUsecase() usecase.TodoUsecase {
	return usecase.NewTodoUsecase(i.reg.TodoRepository())
}
