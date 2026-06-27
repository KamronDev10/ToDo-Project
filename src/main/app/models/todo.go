package models

import "time"

type Todo struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	UserId      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// CREATE TABLE todos (
//     id          SERIAL PRIMARY KEY,
//     title       VARCHAR(256) NOT NULL,
//     description VARCHAR(256) NOT NULL,
//     status      VARCHAR(60)  NOT NULL DEFAULT 'pending',
//     user_id     INT NOT NULL,
//     created_at  TIMESTAMP DEFAULT NOW(),

//     CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
// );
