CREATE TABLE todos (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    status      VARCHAR(60)  NOT NULL DEFAULT 'pending',
    user_id     INT NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);