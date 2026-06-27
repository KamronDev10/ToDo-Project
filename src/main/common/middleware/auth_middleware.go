package middleware

import (
	"context"
	"net/http"
	"strings"
	"todo_app/src/main/common/token"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. Headerdan tokenni ol
		authHeader := r.Header.Get("Authorization")

		// 2. Token bormi?
		if authHeader == "" {
			http.Error(w, "Token yo'q", http.StatusUnauthorized) // 401
			return
		}

		// 3. "Bearer " ni olib tashla
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. ParseToken — haqiqiyми tekshir
		claims, err := token.AuthToken(tokenString)
		if err != nil {
			http.Error(w, "Token yaroqsiz", http.StatusForbidden) // 403
			return
		}

		// 5. Context ga sol — handler ishlatadi
		ctx := context.WithValue(r.Context(), "userID", claims.Id)

		// 6. Hammasi to'g'ri → handler ga o't
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
