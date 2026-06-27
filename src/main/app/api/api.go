package api

import (
	"net/http"
	"todo_app/src/main/app/api/handler"
)

func RegisterUserRoutes(router *http.ServeMux, h *handler.Handler) {

	// api.go
	router.HandleFunc("POST /auth/sign-up", h.CreateUser)
	router.HandleFunc("POST /auth/sign-in", h.Login)
}
