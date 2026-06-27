package service

import (
	"todo_app/src/main/app/models"
	"todo_app/src/main/app/repository"
)

type TodoServiceI interface {
	Create(todo *models.Todo) error
	GetAll(userID int64) ([]*models.Todo, error)
	GetByID(id int64) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id int64) error
}

type TodoService struct {
	todoRepo repository.TodoRepoI
}

func NewTodoService(todoRepo repository.TodoRepoI) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (ts *TodoService) Create(todo *models.Todo) error {
	return ts.todoRepo.Create(todo)
}

func (ts *TodoService) GetAll(userID int64) ([]*models.Todo, error) {
	todos, err := ts.todoRepo.GetAll(userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (ts *TodoService) GetByID(id int64) (*models.Todo, error) {
	todo, err := ts.todoRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (ts *TodoService) Update(todo *models.Todo) error {
	return ts.todoRepo.Update(todo)
}

func (ts *TodoService) Delete(id int64) error {
	return ts.todoRepo.Delete(id)
}
