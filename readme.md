# Todo API 📝

A simple Todo REST API built with Go.

## Tech Stack
- Go
- PostgreSQL
- JWT Authentication
- Swagger UI

## Getting Started

Fill in the .env file:
```
DB_URL=postgres://...
JWT_SECRET=your-secret-key
```

```bash
go run src/main/main.go
```

Swagger: `http://localhost:8080/swagger/index.html`

## API Endpoints

### Auth
- `POST /auth/sign-up` — Register
- `POST /auth/sign-in` — Login

### Todos
- `GET /todos` — Get all todos 🔒
- `GET /todos/get` — Get single todo 🔒
- `POST /todos/create` — Create todo 🔒
- `PUT /todos/update` — Update todo 🔒
- `DELETE /todos/delete` — Delete todo 🔒

> 🔒 — Requires token