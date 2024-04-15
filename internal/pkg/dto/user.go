package dto

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role_id"`
	Created_at time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Id   uuid.UUID `json:"email"`
	Role string
	jwt.StandardClaims
}

type UserSignupResponse struct {
	Id         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Email      string    `json:"email"`
	Role_id    uuid.UUID `json:"role_id"`
	Created_at time.Time `json:"created_at"`
}

type GetAllUserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	RoleID    string    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResBodyStruct struct {
	Token  string    `json:"token"`
	UserId uuid.UUID `json:"userId"`
	Role   string    `json:"role"`
}
