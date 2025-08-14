package models

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type Authentication interface {
	Login(LoginRequest) (LoginResponse,error)
	Register(RegisterRequest) (RegisterResponse,error)
}