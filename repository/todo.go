package repository

import (
	"github.com/nguyenvd27/graphql-test/model"
	staticdatabase "github.com/nguyenvd27/graphql-test/static_database"
	"gorm.io/gorm"
)

type Todo interface {
	CreateTodo(id, text, userId string, done bool) (*model.Todo, error)
	GetTodos() ([]*model.Todo, error)
	GetATodo(id string) (*model.Todo, error)
}

type todo struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) Todo {
	return &todo{
		db: db,
	}
}

func (r *todo) CreateTodo(id, text, userId string, done bool) (*model.Todo, error) {
	todo := &model.Todo{
		ID:     id,
		Text:   text,
		Done:   done,
		UserID: userId,
	}
	staticdatabase.Todos = append(staticdatabase.Todos, todo)
	return todo, nil
}

func (r *todo) GetTodos() ([]*model.Todo, error) {
	return staticdatabase.Todos, nil
}

func (r *todo) GetATodo(id string) (*model.Todo, error) {
	for i := 0; i < len(staticdatabase.Todos); i++ {
		if id == staticdatabase.Todos[i].ID {
			return staticdatabase.Todos[i], nil
		}
	}
	return &model.Todo{}, nil
}
