package entity

import "time"

// RegisterRequest is used as the request body on register
type RegisterRequest struct {
	Phone string `json:"phone" form:"phone"`
	Name  string `json:"name" form:"name"`
	Role  string `json:"role" form:"role"`
}

// RegisterResponse is used as the response on register
type RegisterResponse struct {
	RegisterRequest RegisterRequest `json:"register_request"`
	Password        string          `json:"password"`
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

// PrivateClaimsRequest is used as the request body on PrivateClaims
type PrivateClaimsRequest struct {
	Token string `json:"token" form:"token"`
}

// PrivateClaimsResponse is used as the response on PrivateClaims
type PrivateClaimsResponse struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Timestamp time.Time `json:"timestamp"`
}
