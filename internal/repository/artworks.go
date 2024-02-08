package repository

import (
	"time"

	"github.com/google/uuid"
)

type Artworks struct {
	id             uuid.UUID
	name           string
	desc           string
	image          string
	starting_price float64
	category_id    uuid.UUID
	live_period    time.Time
	status         string
	owner_id       uuid.UUID
	highest_bid    uuid.UUID
	created_at     time.Time
}
