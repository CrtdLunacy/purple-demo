package verify

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type SendResponse struct {
	RegisterMsg string `json:"register_msg"`
}
