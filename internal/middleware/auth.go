package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/helper"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := helper.NewJWT().Parse(token)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", data.Id)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
