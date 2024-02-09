package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ArtworkStorer interface {
	CreateArtwork(artwork Artworks) (Artworks, error)
	GetCategory(categoryName string) (Category, error)
}

type Artworks struct {
	Id             uuid.UUID
	Name           string
	Description    string
	Image          string
	Starting_price float64
	Category_id    uuid.UUID
	Live_period    string
	Status         string
	Owner_id       uuid.UUID
	Highest_bid    uuid.UUID
	Created_at     time.Time
}

type Category struct {
	Id   uuid.UUID
	Name string
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

	err := as.DB.QueryRow("INSERT INTO artworks(id, name, image, starting_price, category_id, live_period,status, owner_id, created_at, description) VALUES($1, $2, $3, $4, $5, 'open', $6, $7, $8, $9) RETURNING id",
		artwork.Id, artwork.Name, artwork.Image, artwork.Starting_price, artwork.Category_id, artwork.Live_period, artwork.Owner_id, artwork.Created_at, artwork.Description).Scan(&artwork.Id)

	if err != nil {
		return Artworks{}, err
	}
	return artwork, nil
}

func (as *artworkStore) GetCategory(categoryName string) (Category, error) {

	var category Category

	rows, err := as.DB.Query("SELECT id, name FROM category where name = $1", categoryName)
	if err != nil {
		return Category{}, err
	}

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return Category{}, err
		}
	}

	return category, nil

}
