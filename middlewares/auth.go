package middlewares

import (
	"context"
	"ecommerce-api/utils"

	"net/http"
	"strings"
)

func JWTAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// 🔥 inject ke context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)

		if !ok || role != "admin" {
			utils.Error(w, http.StatusForbidden, "FORBIDDEN", "forbidden")
			return
		}

		next(w, r)
	}
}
