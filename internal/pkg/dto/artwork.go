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
	Duration       time.Duration
	Owner_id       string
	Category       string
}

type GetArtworkResponse struct {
	Id             uuid.UUID
	Name           string
	Description    string
	Image          string
	Starting_price float64
	Category       string
	Closing_time   string
	Owner_id       uuid.UUID
	Highest_bid    int
	Created_at     time.Time
}
