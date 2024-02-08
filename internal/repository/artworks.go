package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ArtworkStorer interface {
	CreateArtwork(artwork Artworks) (Artworks, error)
}

type Artworks struct {
	Id             uuid.UUID
	Name           string
	Description    string
	Image          string
	Starting_price float64
	Category_id    uuid.UUID
	Live_period    time.Time
	Status         string
	Owner_id       uuid.UUID
	Highest_bid    uuid.UUID
	Created_at     time.Time
}

type artworkStore struct {
	DB *sqlx.DB
}

func NewArtworkRepo(db *sqlx.DB) ArtworkStorer {
	return &artworkStore{DB: db}
}

func (as *artworkStore) CreateArtwork(artwork Artworks) (Artworks, error) {
	artwork.Id = uuid.New()
	artwork.Created_at = time.Now()
	artwork.Status = "open"
	//Change uuids by inserting data
	artwork.Category_id, _ = uuid.Parse("6a55565e-3b0f-48fe-854e-ea22ce1ff991")
	artwork.Owner_id, _ = uuid.Parse("6a55565e-3b0f-48fe-854e-ea22ce1ff991")
	err := as.DB.QueryRow("INSERT INTO artworks(id, name, image, starting_price, category_id, live_period,status, owner_id, created_at, description) VALUES($1, $2, $3, $4, $5, 'open', $7, $8, $9, $10) RETURNING id",
		artwork.Id, artwork.Name, artwork.Image, artwork.Starting_price, artwork.Category_id, artwork.Owner_id, artwork.Created_at, artwork.Description).Scan(&artwork.Id)

	if err != nil {
		return Artworks{}, err
	}
	return artwork, nil
}
