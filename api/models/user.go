package models

import (
	"github.com/satori/go.uuid"
)


type User struct {
	ID uuid.UUID `json:"id"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName string `json:"lastname" form:"lastname"`
}