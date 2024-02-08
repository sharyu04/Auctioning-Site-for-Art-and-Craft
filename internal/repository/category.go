package repository

import "github.com/google/uuid"

type Category struct {
	id   uuid.UUID
	name string
}
