package middlewares

import (
	"context"
	"ecommerce-api/utils"
	"net/http"

	"github.com/google/uuid"
)

func RequestID(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := uuid.NewString()

		ctx := context.WithValue(
			r.Context(),
			utils.RequestIDKey,
			requestID,
		)

		w.Header().Set(
			"X-Request-ID",
			requestID,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})

}
