package repository

import (
	"github.com/nguyenvd27/graphql-test/model"
	staticdatabase "github.com/nguyenvd27/graphql-test/static_database"
	"gorm.io/gorm"
)

type Event interface {
	CreateEvent(id, title, content, userID string) (*model.Event, error)
	GetEvents() ([]*model.Event, error)
	GetAEvent(id string) (*model.Event, error)
}

type event struct {
	db *gorm.DB
}

func NewEvent(db *gorm.DB) Event {
	return &event{db: db}
}

func (r *event) CreateEvent(id, title, content, userId string) (*model.Event, error) {
	event := &model.Event{
		ID:      id,
		Title:   title,
		Content: content,
		UserID:  userId,
	}

	staticdatabase.Events = append(staticdatabase.Events, event)
	return event, nil
}

func (r *event) GetEvents() ([]*model.Event, error) {
	return staticdatabase.Events, nil
}

func (r *event) GetAEvent(id string) (*model.Event, error) {
	for i := 0; i < len(staticdatabase.Events); i++ {
		if id == staticdatabase.Events[i].ID {
			return staticdatabase.Events[i], nil
		}
	}
	return &model.Event{}, nil
}
