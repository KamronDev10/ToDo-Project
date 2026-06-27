package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`

	jwt.RegisteredClaims
}

var key = []byte(os.Getenv("JWT_SECRET"))

func GetToken(id int64, username string, email string, role string) (string, error) {

	expriationTime := time.Now().Add(48 * time.Hour)
	claims := &CustomClaims{
		Id:       id,
		Username: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expriationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "blog-server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}

// tokenni haqiqiyligini tekshirish uchun
func AuthToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("noto'g'ri algoritm")
			}

			return key, nil
		},
	)

	if err != nil {
		return nil, err
	}

	// 2. Claims ni ol
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token yaroqsiz")
	}

	return claims, nil
}
