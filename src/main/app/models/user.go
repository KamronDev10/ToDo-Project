package models

import "time"

type User struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	RegisteredAt time.Time `json:"registered_at"`
}
