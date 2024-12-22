package auth

type AuthRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type AuthResponse struct {
	SessionId string `json:"session_id"`
	Code      int    `json:"code"`
}

type VerifyRequest struct {
	SessionId string `json:"session_id" validate:"required"`
	Code      int    `json:"code" validate:"required"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
