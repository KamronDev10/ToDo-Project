package handler

import (
	"encoding/json"
	"net/http"
	"todo_app/src/main/app/dto"
	"todo_app/src/main/app/models"
)

// @Summary Foydalanuvchi yaratish (Sign Up)
// @Description Yangi foydalanuvchini bazaga qo'shadi va token qaytaradi
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.SignUpRequest true "Foydalanuvchi ma'lumotlari"
// @Success 201 {object} map[string]string "token"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/sign-up [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Noto'g'ri ma'lumot", http.StatusBadRequest)
		return
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	// 1. Service orqali yaratish (token qaytadi)
	token, err := h.ServiceUser.Register(user)
	if err != nil {
		// Agar xatolik bo'lsa (masalan, email bazada bor bo'lsa)
		http.Error(w, "Foydalanuvchi yaratilmadi: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Muvaffaqiyatli javob qaytarish
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Javob sifatida token yoki user ma'lumotlarini JSON qilib qaytaramiz
	json.NewEncoder(w).Encode(map[string]string{
		"token":   token,
		"message": "Foydalanuvchi muvaffaqiyatli yaratildi",
	})
}

// @Summary Foydalanuvchi kirishi (Sign In)
// @Description Email va parol bilan kirish, token qaytaradi
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.SignInRequest true "Email va parol"
// @Success 200 {object} map[string]string "token"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/sign-in [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req dto.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Noto'g'ri malumot ", http.StatusBadRequest)
	}

	token, err := h.ServiceUser.Login(req.Email, req.Password)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}
