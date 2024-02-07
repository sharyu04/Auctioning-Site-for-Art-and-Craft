package repository

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	firstName  string
	lastName   string
	email      string
	password   string
	role_id    uuid.UUID
	created_at time.Time
}
