package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateArtworkRequest struct {
	Name           string
	Description    string
	Image          string
	Starting_price float64
	Live_period    time.Time
	Status         string
	Owner_id       uuid.UUID
}
