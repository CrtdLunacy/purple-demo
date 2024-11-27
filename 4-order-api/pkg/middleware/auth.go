package middleware

import (
	"context"
	"go/order-api/configs"
	"go/order-api/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	CtxPhoneKey   key = "CtxPhoneKey"
	CtxSessionKey key = "CtxSessionKey"
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

		ctx := context.WithValue(r.Context(), CtxPhoneKey, data.Phone)
		ctx = context.WithValue(r.Context(), CtxSessionKey, data.SessionId)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
