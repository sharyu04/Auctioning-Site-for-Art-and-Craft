package repository

import (
	"time"

	"github.com/google/uuid"
)

type Bids struct {
	id         uuid.UUID
	artwork_id uuid.UUID
	amount     float64
	status     uuid.UUID
	bidder_id  uuid.UUID
	created_at time.Time
}
