package api

import (
	"net/http"
	"todo_app/src/main/app/api/handler"
	"todo_app/src/main/common/middleware"
)

func RegisterUserRoutes(router *http.ServeMux, h *handler.Handler) {

	// api.go
	router.HandleFunc("POST /auth/sign-up", h.CreateUser)
	router.HandleFunc("POST /auth/sign-in", h.Login)

}

func RegisterTodoRoutes(router *http.ServeMux, h *handler.Handler) {

	router.Handle("GET /todos",
		middleware.AuthMiddleware(http.HandlerFunc(h.GetAllTodos)))

	router.Handle("GET /todos/get",
		middleware.AuthMiddleware(http.HandlerFunc(h.GetTodoByID)))

	router.Handle("POST /todos/create",
		middleware.AuthMiddleware(http.HandlerFunc(h.CreateTodo)))

	router.Handle("PUT /todos/update",
		middleware.AuthMiddleware(http.HandlerFunc(h.UpdateTodo)))

	router.Handle("DELETE /todos/delete",
		middleware.AuthMiddleware(http.HandlerFunc(h.DeleteTodo)))
}
