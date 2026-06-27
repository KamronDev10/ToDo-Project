package service

import (
	"errors"
	"todo_app/src/main/app/models"
	"todo_app/src/main/app/repository"
	"todo_app/src/main/common/token"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceI interface {
	Register(user *models.User) (string, error)
	Login(email, password string) (string, error)
}

type UserService struct {
	userRepo repository.UserRepoI
}

func NewUserService(userRepo repository.UserRepoI) UserServiceI {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) Register(user *models.User) (string, error) {
	// 1. Parolni hash qilish
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.PasswordHash = string(hashedPassword)

	// 2. DB ga saqlash
	err = us.userRepo.Create(user)
	if err != nil {
		return "", err
	}
	token, err := token.GetToken(user.Id, user.Username, user.Email, "user")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) Login(email, password string) (string, error) {
	// Email bo'yicha userni topish
	user, err := us.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("user topilmadi")
	}

	// Parolni tekshirish
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("parol noto'g'ri")
	}

	// Token yaratish
	tokenString, err := token.GetToken(user.Id, user.Username, user.Email, "user")
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
