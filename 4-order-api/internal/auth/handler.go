package auth

import (
	"go/order-api/configs"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (authHandler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if _, err := request.HandleBody[LoginRequest](w, req); err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := LoginResponse{
			Token: authHandler.Config.Auth.Secret,
		}
		response.ResponseJSON(w, res, http.StatusAccepted)
	}
}

func (authHandler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if _, err := request.HandleBody[RegisterRequest](w, req); err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := RegisterResponse{
			RegisterMsg: "Регистрация успешна",
		}
		response.ResponseJSON(w, res, http.StatusCreated)
	}
}
