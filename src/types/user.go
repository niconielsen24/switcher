package types

import "uuid"

type User struct {
	Username string `json:"username"`
	Id       string `json:"id"`
}

func NewUser(username string) *User {
	return &User{Username: username, Id: uuid.New().String()}
}
