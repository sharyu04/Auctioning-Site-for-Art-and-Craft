package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateArtworkRequest struct {
	Name           string        `json: name`
	Description    string        `json: desc`
	Image          string        `json: image`
	Starting_price float64       `json: starting_price`
	Duration       time.Duration `json: duration`
	Owner_id       string
	Category       string `json: category`
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
