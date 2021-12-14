package interactor

import (
	"github.com/nguyenvd27/graphql-test/repository"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	EventRepository() repository.Event
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{
		db: db,
	}
}

func (r *registry) EventRepository() repository.Event {
	return repository.NewEvent(r.db)
}
