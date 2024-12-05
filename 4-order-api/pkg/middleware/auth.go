package middleware

import (
	"context"
	"go/order-api/configs"
	"go/order-api/pkg/jwt"
	"net/http"
	"strings"
)

type key string
type AuthContext struct {
	SessionId string
	Phone     string
}

const (
	CtxAuthData key = "authData"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func AuthCheck(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeUnauthorized(w)
			return
		}

		if authHeader == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthorized(w)
			return
		}

		authCtx := AuthContext{
			SessionId: data.SessionId,
			Phone:     data.Phone,
		}

		ctx := context.WithValue(r.Context(), "authData", authCtx)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
