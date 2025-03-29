package models

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserCollection struct {
	Data []User `json:"data"`
}
