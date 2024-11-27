package auth

import (
	"go/order-api/configs"
	"go/order-api/pkg/jwt"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth", handler.Auth())
	router.HandleFunc("POST /auth/verify", handler.Verify())
}

func (authHandler *AuthHandler) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[AuthRequest](w, req)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		session, err := authHandler.AuthService.Auth(&AuthData{
			Phone: body.Phone,
		})
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := AuthResponse{
			SessionId: session.SessionId,
			Code:      session.Code,
		}
		response.ResponseJSON(w, res, http.StatusAccepted)
	}
}

func (authHandler *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[VerifyRequest](w, req)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := authHandler.AuthService.Verify(&VerifyData{
			SessionId: body.SessionId,
			Code:      body.Code,
		})
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jwtToken, err := jwt.NewJWT(authHandler.Config.Auth.Secret).Create(&jwt.TokenData{
			Phone:     user.Phone,
			SessionId: user.SessionId,
			Code:      user.Code,
		})

		if err != nil {
			response.ResponseJSON(w, ErrCommon, http.StatusInternalServerError)
			return
		}

		res := VerifyResponse{
			Token: jwtToken,
		}

		response.ResponseJSON(w, res, http.StatusCreated)
	}
}
