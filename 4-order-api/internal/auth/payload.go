package auth

type BaseFields struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	BaseFields
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	BaseFields
	Name string `json:"name" validate:"required,max=20"`
}

type RegisterResponse struct {
	RegisterMsg string `json:"register_msg"`
}
