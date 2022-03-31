package models

type User struct {
	UserId   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
