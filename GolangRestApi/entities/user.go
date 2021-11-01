package entities

import (
	"github.com/google/uuid"
)

type ResponseGetlistUser struct {
	User []*User `json:"data"`
}

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}
