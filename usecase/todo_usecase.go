package usecase

import (
	"github.com/nguyenvd27/graphql-test/model"
	"github.com/nguyenvd27/graphql-test/repository"
)

type TodoUsecase interface {
	CreateTodoUsecase(id, text, userId string, done bool) (*model.Todo, error)
	GetAllTodosUsecase() ([]*model.Todo, error)
	GetATodoUsecase(id string) (*model.Todo, error)
}

type todoUsecase struct {
	todoRepo repository.Todo
}

func NewTodoUsecase(todoRepo repository.Todo) TodoUsecase {
	return &todoUsecase{
		todoRepo: todoRepo,
	}
}

func (u *todoUsecase) CreateTodoUsecase(id, text, userId string, done bool) (*model.Todo, error) {
	newTodo, err := u.todoRepo.CreateTodo(id, text, userId, done)
	if err != nil {
		return nil, err
	}
	return newTodo, err
}

func (u *todoUsecase) GetAllTodosUsecase() ([]*model.Todo, error) {
	todos, err := u.todoRepo.GetTodos()
	if err != nil {
		return nil, err
	}
	return todos, err
}

func (u *todoUsecase) GetATodoUsecase(id string) (*model.Todo, error) {
	todo, err := u.todoRepo.GetATodo(id)
	if err != nil {
		return nil, err
	}
	return todo, err
}
