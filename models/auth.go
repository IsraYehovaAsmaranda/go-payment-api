package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
