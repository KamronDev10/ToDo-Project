package repository

import (
	"database/sql"
	"todo_app/src/main/app/models"
)

type UserRepoI interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoI {
	return &UserRepo{db: db}
}

func (ur *UserRepo) Create(user *models.User) error {
	query := `
        INSERT INTO users (username, email, password_hash)
        VALUES ($1, $2, $3)
    `
	_, err := ur.db.Exec(query,
		user.Username,
		user.Email,
		user.PasswordHash,
	)
	return err
}

func (ur *UserRepo) GetByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, registered_at FROM users WHERE email = $1`

	var user models.User
	err := ur.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.RegisteredAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
