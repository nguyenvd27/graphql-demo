package usecase

import (
	"github.com/nguyenvd27/graphql-test/model"
	"github.com/nguyenvd27/graphql-test/repository"
)

type EventUsecase interface {
	CreateEventUsecase(id, title, content, userID string) (*model.Event, error)
	GetAllEventsUsecase() ([]*model.Event, error)
	GetAEventUsecase(id string) (*model.Event, error)
}

type eventUsecase struct {
	eventRepo repository.Event
}

func NewEventUsecase(eventRepo repository.Event) EventUsecase {
	return &eventUsecase{
		eventRepo: eventRepo,
	}
}

func (u *eventUsecase) CreateEventUsecase(id, title, content, userID string) (*model.Event, error) {
	newEvent, err := u.eventRepo.CreateEvent(id, title, content, userID)
	if err != nil {
		return nil, err
	}
	return newEvent, err
}

func (u *eventUsecase) GetAllEventsUsecase() ([]*model.Event, error) {
	events, err := u.eventRepo.GetEvents()
	if err != nil {
		return nil, err
	}
	return events, err
}

func (u *eventUsecase) GetAEventUsecase(id string) (*model.Event, error) {
	event, err := u.eventRepo.GetAEvent(id)
	if err != nil {
		return nil, err
	}
	return event, err
}
