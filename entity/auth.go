package entity

import "time"

// RegisterRequest is used as the request body on register
type RegisterRequest struct {
	Phone string `json:"phone" form:"phone"`
	Name  string `json:"name" form:"name"`
	Role  string `json:"role" form:"role"`
}

// LoginRequest is used as the request body on login
type LoginRequest struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

// LoginResponse is used as the response on login
type LoginResponse struct {
	Token string `json:"token"`
}

// ValidateTokenResponse is used as the response on ValidateToken
type ValidateTokenResponse struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Timestamp time.Time `json:"timestamp"`
}
