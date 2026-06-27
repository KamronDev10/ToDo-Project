package handler

import (
	"encoding/json"
	"todo_app/src/main/app/dto"
	"todo_app/src/main/app/models"

	"net/http"
	"strconv"
)

// @Summary Todo yaratish
// @Tags Todos
// @Accept json
// @Produce json
// @Param todo body dto.CreateTodoRequest true "Todo ma'lumotlari"
// @Success 201 {string} string "Yaratildi"
// @Security BearerAuth
// @Router /todos/create [post]
func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Noto'g'ri ma'lumot", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(int64)

	if err := h.TodoService.Create(&models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      "pending",
		UserId:      userID,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Todo yaratildi"))
}

// @Summary Barcha todolar
// @Tags Todos
// @Produce json
// @Success 200 {array} models.Todo
// @Security BearerAuth
// @Router /todos [get]
func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	todos, err := h.TodoService.GetAll(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// @Summary Bitta todo
// @Tags Todos
// @Produce json
// @Param id query int true "Todo ID"
// @Success 200 {object} models.Todo
// @Security BearerAuth
// @Router /todos/get [get]
func (h *Handler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	todo, err := h.TodoService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// @Summary Todo yangilash
// @Tags Todos
// @Accept json
// @Produce json
// @Param id query int true "Todo ID"
// @Param todo body dto.UpdateTodoRequest true "Todo ma'lumotlari"
// @Success 200 {string} string "Yangilandi"
// @Security BearerAuth
// @Router /todos/update [put]
func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	var req dto.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Noto'g'ri ma'lumot", http.StatusBadRequest)
		return
	}

	if err := h.TodoService.Update(&models.Todo{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo yangilandi"))
}

// @Summary Todo o'chirish
// @Tags Todos
// @Param id query int true "Todo ID"
// @Success 200 {string} string "O'chirildi"
// @Security BearerAuth
// @Router /todos/delete [delete]
func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	if err := h.TodoService.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo o'chirildi"))
}
