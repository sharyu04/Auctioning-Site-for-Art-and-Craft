package dto

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:id`
	FirstName  string    `json:firstname`
	LastName   string    `json:lastname`
	Email      string    `json:email`
	Role_id    uuid.UUID `json:role_id`
	Created_at time.Time `json:created_at`
}

type CreateUserRequest struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
	Email     string `json:email`
	Password  string `json:password`
}

type LoginRequest struct {
	Email    string `json:email`
	Password string `json:password`
}
