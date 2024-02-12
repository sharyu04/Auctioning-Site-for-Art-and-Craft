package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

type ArtworkStorer interface {
	CreateArtwork(artwork Artworks) (Artworks, error)
	GetCategory(categoryName string) (Category, error)
	GetAllArtworks(start, count int) ([]dto.GetArtworkResponse, error)
	GetFilterArtworks(category string, start, count int) ([]dto.GetArtworkResponse, error)
	GetArtworkById(artworkId uuid.UUID) (dto.GetArtworkResponse, error)
	DeleteArtworkById(artworkId uuid.UUID, ownerId uuid.UUID, role string) error
}

type Artworks struct {
	Id             uuid.UUID `db: id`
	Name           string    `db:name`
	Description    string    `db:description`
	Image          string    `db:image`
	Starting_price float64   `db:starting_price`
	Category_id    uuid.UUID `db:category_id`
	Closing_time   time.Time `db:closing_time`
	Owner_id       uuid.UUID `db:owner_id`
	Highest_bid    uuid.UUID `db:highest_bid`
	Created_at     time.Time `db:created_at`
}

type Category struct {
	Id   uuid.UUID `db:id`
	Name string    `db:name`
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

	err := as.DB.QueryRow("INSERT INTO artworks(id, name, image, starting_price, category_id, closing_time, owner_id, created_at, description) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		artwork.Id, artwork.Name, artwork.Image, artwork.Starting_price, artwork.Category_id, artwork.Closing_time, artwork.Owner_id, artwork.Created_at, artwork.Description).Scan(&artwork.Id)

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
	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return Category{}, err
		}
	}

	if i == 0 {
		return Category{}, errors.New("Please choose correct category")
	}

	return category, nil

}

func (as *artworkStore) GetAllArtworks(start, count int) ([]dto.GetArtworkResponse, error) {
	artworks := []dto.GetArtworkResponse{}
	row, err := as.DB.Query("select artworks.id, artworks.name, artworks.description, artworks.image, artworks.starting_price, category.name, artworks.closing_time, artworks.owner_id, artworks.created_at, artworks.highest_bid from artworks inner join category on artworks.category_id = category.id LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return artworks, err
	}
	defer row.Close()
	for row.Next() {
		var a dto.GetArtworkResponse
		var highest_bid uuid.UUID
		if err := row.Scan(&a.Id, &a.Name, &a.Description, &a.Image, &a.Starting_price, &a.Category, &a.Closing_time, &a.Owner_id, &a.Created_at, &highest_bid); err != nil {
			return nil, err
		}
		if highest_bid != uuid.Nil {
			err := as.DB.QueryRow("SELECT amount FROM bids where bids.id = $1", highest_bid).Scan(&a.Highest_bid)
			if err != nil {
				return []dto.GetArtworkResponse{}, err
			}
		} else {
			a.Highest_bid = 0
		}
		artworks = append(artworks, a)
	}
	return artworks, nil

}

func (as *artworkStore) GetFilterArtworks(category string, start, count int) ([]dto.GetArtworkResponse, error) {
	artworks := []dto.GetArtworkResponse{}
	row, err := as.DB.Query("select artworks.id, artworks.name, artworks.description, artworks.image, artworks.starting_price, category.name, artworks.closing_time, artworks.owner_id, artworks.created_at, artworks.highest_bid from artworks inner join category on artworks.category_id = category.id where category.name = $1 LIMIT $2 OFFSET $3", category, count, start)
	if err != nil {
		return artworks, err
	}
	defer row.Close()
	for row.Next() {
		var a dto.GetArtworkResponse
		var highest_bid uuid.UUID
		if err := row.Scan(&a.Id, &a.Name, &a.Description, &a.Image, &a.Starting_price, &a.Category, &a.Closing_time, &a.Owner_id, &a.Created_at, &highest_bid); err != nil {
			return []dto.GetArtworkResponse{}, err
		}
		if highest_bid != uuid.Nil {
			err := as.DB.QueryRow("SELECT amount FROM bids where bids.id = $1", highest_bid).Scan(&a.Highest_bid)
			if err != nil {
				return []dto.GetArtworkResponse{}, err
			}
		} else {
			a.Highest_bid = 0
		}
		artworks = append(artworks, a)
	}

	return artworks, nil

}

func (as *artworkStore) GetArtworkById(artworkId uuid.UUID) (dto.GetArtworkResponse, error) {
	row, err := as.DB.Query("select artworks.id, artworks.name, artworks.description, artworks.image, artworks.starting_price, category.name, artworks.closing_time, artworks.owner_id, artworks.created_at, artworks.highest_bid from artworks inner join category on artworks.category_id = category.id where artworks.id = $1", artworkId)
	if err != nil {
		return dto.GetArtworkResponse{}, err
	}
	defer row.Close()
	var a dto.GetArtworkResponse
	i := 0
	for row.Next() {
		var highest_bid uuid.UUID
		if err := row.Scan(&a.Id, &a.Name, &a.Description, &a.Image, &a.Starting_price, &a.Category, &a.Closing_time, &a.Owner_id, &a.Created_at, &highest_bid); err != nil {
			return dto.GetArtworkResponse{}, err
		}
		if highest_bid != uuid.Nil {
			err := as.DB.QueryRow("SELECT amount FROM bids where bids.id = $1", highest_bid).Scan(&a.Highest_bid)
			if err != nil {
				return dto.GetArtworkResponse{}, err
			}
		} else {
			a.Highest_bid = 0
		}
		i++
	}

	if i == 0 {
		return dto.GetArtworkResponse{}, errors.New("Wrong atrwork id")
	}
	return a, nil
}

func (as *artworkStore) DeleteArtworkById(artworkId uuid.UUID, ownerId uuid.UUID, role string) error {
	if role == "user" {
		rows, err := as.DB.Query("SELECT * FROM artworks where id = $1 and owner_id = $2", artworkId, ownerId)
		if err != nil {
			return err
		}
		i := 0
		for rows.Next() {
			i++
		}
		if i == 0 {
			return errors.New("Logged in user does not own any such artwork!")
		}
	}

	_, err := as.DB.Query("Update artworks set highest_bid = null where id = $1", artworkId)
	if err != nil {

		return err
	}

	_, err = as.DB.Query("Delete from bids where artwork_id = $1", artworkId)
	if err != nil {
		return err
	}

	_, err = as.DB.Query("Delete from artworks where id = $1", artworkId)
	if err != nil {
		return err
	}

	_, err = as.DB.Query("Delete from artworks where id = $1", artworkId)
	if err != nil {
		return err
	}

	return nil

}
