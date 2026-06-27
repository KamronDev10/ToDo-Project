package repository

import (
	"database/sql"
	"todo_app/src/main/app/models"
)

type TodoRepoI interface {
	Create(todo *models.Todo) error
	GetAll(userId int64) ([]*models.Todo, error)
	GetById(id int64) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id int64) error
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) TodoRepoI {
	return &TodoRepo{db: db}
}

func (tr *TodoRepo) Create(todo *models.Todo) error {
	query := `INSERT INTO todos (title, description, status, user_id) VALUES ($1 , $2 , $3 , $4)`
	_, err := tr.db.Exec(query, todo.Title, todo.Description, todo.Status, todo.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TodoRepo) GetAll(userId int64) ([]*models.Todo, error) {
	query := `SELECT id, title, description, status, user_id, created_at FROM todos WHERE user_id = $1`
	rows, err := tr.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*models.Todo

	for rows.Next() {
		todo := models.Todo{}
		err := rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.Status,
			&todo.UserId,
			&todo.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)

	}
	return todos, nil
}

func (tr *TodoRepo) GetById(id int64) (*models.Todo, error) {
	query := `SELECT id, title, description, status, user_id, created_at FROM todos WHERE id = $1`

	todo := &models.Todo{}
	err := tr.db.QueryRow(query, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Description,
		&todo.Status,
		&todo.UserId,
		&todo.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (tr *TodoRepo) Update(todo *models.Todo) error {
	query := `UPDATE todos SET title=$1, description=$2, status=$3 WHERE id=$4`
	_, err := tr.db.Exec(query, todo.Title, todo.Description, todo.Status, todo.Id)
	return err
}

func (tr *TodoRepo) Delete(id int64) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := tr.db.Exec(query, id)
	return err
}
