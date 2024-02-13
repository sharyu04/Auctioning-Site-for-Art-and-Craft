package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateArtworkRequest struct {
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Image          string        `json:"image"`
	Starting_price float64       `json:"starting_price"`
	Duration       time.Duration `json:"duration"`
	Owner_id       string
	Category       string `json:"category"`
}

type GetArtworkResponse struct {
	Id             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	Image          string    `db:"image"`
	Starting_price float64   `db:"starting_price"`
	Category       string
	Closing_time   string    `db:"closing_time"`
	Owner_id       uuid.UUID `db:"owner_id"`
	Highest_bid    int       `db:"highest_bid"`
	Created_at     time.Time `db:"created_at"`
}
