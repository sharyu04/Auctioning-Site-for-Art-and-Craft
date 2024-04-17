package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

type ArtworkStorer interface {
	CreateArtwork(artwork Artworks) (Artworks, error)
	GetCategory(categoryName string) (Category, error)
	GetAllArtworks(start, count int) ([]dto.GetArtworkResponse, int, error)
	GetFilterArtworks(category string, start, count int) ([]dto.GetArtworkResponse, int, error)
	GetArtworkById(artworkId uuid.UUID) (dto.GetArtworkResponse, error)
	DeleteArtworkById(artworkId uuid.UUID, ownerId uuid.UUID, role string) error
	GetAllCategories() ([]Category, error)
	CreateCategory(name string) error
}

type Artworks struct {
	Id             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	Image          string    `db:"image"`
	Starting_price float64   `db:"starting_price"`
	Category_id    uuid.UUID `db:"category_id"`
	Closing_time   time.Time `db:"closing_time"`
	Owner_id       uuid.UUID `db:"owner_id"`
	Highest_bid    uuid.UUID `db:"highest_bid"`
	Created_at     time.Time `db:"created_at"`
}

type Category struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
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

	ctx := context.Background()
	tx, err := as.DB.BeginTxx(ctx, nil)
	if err != nil {
		return Artworks{}, err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO artworks(id, name, image, starting_price, category_id, closing_time, owner_id, created_at, description) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		artwork.Id, artwork.Name, artwork.Image, artwork.Starting_price, artwork.Category_id, artwork.Closing_time, artwork.Owner_id, artwork.Created_at, artwork.Description)

	// err = as.DB.QueryRow("INSERT INTO artworks(id, name, image, starting_price, category_id, closing_time, owner_id, created_at, description) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
	// artwork.Id, artwork.Name, artwork.Image, artwork.Starting_price, artwork.Category_id, artwork.Closing_time, artwork.Owner_id, artwork.Created_at, artwork.Description).Scan(&artwork.Id)

	if err != nil {
		tx.Rollback()
		return Artworks{}, err
	}

	err = tx.Commit()
	if err != nil {
		return Artworks{}, err
	} else {
		fmt.Println("Transaction commited successfully")
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
		return Category{}, apperrors.BadRequest{ErrorMsg: "Invalid Category"}
	}

	return category, nil

}

func (as *artworkStore) GetAllCategories() ([]Category, error) {

	category := []Category{}

	rows, err := as.DB.Query("SELECT * FROM category")
	if err != nil {
		return category, err
	}
	i := 0
	for rows.Next() {
		i++
		var cat Category
		err := rows.Scan(&cat.Id, &cat.Name)
		if err != nil {
			return category, err
		}
		category = append(category, cat)
	}

	if i == 0 {
		return category, apperrors.BadRequest{ErrorMsg: "No categories exist"}
	}

	return category, nil

}

func (as *artworkStore) GetAllArtworks(start, count int) ([]dto.GetArtworkResponse, int, error) {
	artworks := []dto.GetArtworkResponse{}
	totalRows, err := as.DB.Query("select count(*) from artworks")
	if err != nil {
		return artworks, 0, err
	}
	totalCount := 0
	for totalRows.Next() {
		err = totalRows.Scan(&totalCount)
		fmt.Println("TotalCount", totalCount)
	}
	row, err := as.DB.Query("select artworks.id, artworks.name, artworks.description, artworks.image, artworks.starting_price, category.name, artworks.closing_time, artworks.owner_id, artworks.created_at, artworks.highest_bid from artworks inner join category on artworks.category_id = category.id LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return artworks, 0, err
	}
	defer row.Close()
	for row.Next() {
		var a dto.GetArtworkResponse
		var highest_bid uuid.UUID
		if err := row.Scan(&a.Id, &a.Name, &a.Description, &a.Image, &a.Starting_price, &a.Category, &a.Closing_time, &a.Owner_id, &a.Created_at, &highest_bid); err != nil {
			return nil, 0, err
		}
		if highest_bid != uuid.Nil {
			err := as.DB.QueryRow("SELECT amount FROM bids where bids.id = $1", highest_bid).Scan(&a.Highest_bid)
			if err != nil {
				return []dto.GetArtworkResponse{}, 0, err
			}
		} else {
			a.Highest_bid = 0
		}
		artworks = append(artworks, a)
	}
	return artworks, totalCount, nil

}

func (as *artworkStore) GetFilterArtworks(category string, start, count int) ([]dto.GetArtworkResponse, int, error) {
	artworks := []dto.GetArtworkResponse{}
	totalRows, err := as.DB.Query("select count(*) from artworks")
	if err != nil {
		return artworks, 0, err
	}
	totalCount := 0
	for totalRows.Next() {
		err = totalRows.Scan(&totalCount)
		fmt.Println("TotalCount", totalCount)
	}
	row, err := as.DB.Query("select artworks.id, artworks.name, artworks.description, artworks.image, artworks.starting_price, category.name, artworks.closing_time, artworks.owner_id, artworks.created_at, artworks.highest_bid from artworks inner join category on artworks.category_id = category.id where category.name = $1 LIMIT $2 OFFSET $3", category, count, start)
	if err != nil {
		return artworks, 0, err
	}
	defer row.Close()
	for row.Next() {
		var a dto.GetArtworkResponse
		var highest_bid uuid.UUID
		if err := row.Scan(&a.Id, &a.Name, &a.Description, &a.Image, &a.Starting_price, &a.Category, &a.Closing_time, &a.Owner_id, &a.Created_at, &highest_bid); err != nil {
			return []dto.GetArtworkResponse{}, 0, err
		}
		if highest_bid != uuid.Nil {
			err := as.DB.QueryRow("SELECT amount FROM bids where bids.id = $1", highest_bid).Scan(&a.Highest_bid)
			if err != nil {
				return []dto.GetArtworkResponse{}, 0, err
			}
		} else {
			a.Highest_bid = 0
		}
		artworks = append(artworks, a)
	}

	return artworks, totalCount, nil

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
		errMsg := fmt.Sprintf("Artwork not found with id: %v", artworkId)
		return dto.GetArtworkResponse{}, apperrors.BadRequest{ErrorMsg: errMsg}
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
			return apperrors.UnAuthorizedAccess{ErrorMsg: "Unauthorized Delete"}
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

func (as *artworkStore) CreateCategory(name string) error {

	rows, err := as.DB.Query("SELECT id, name FROM category where name = $1", name)
	if err != nil {
		return err
	}
	i := 0
	for rows.Next() {
		i++
	}
	if i != 0 {
		return apperrors.BadRequest{ErrorMsg: "Category already exists"}
	}

	_, err = as.DB.Query("Insert into category(id, name) values(gen_random_uuid(),$1)", name)
	if err != nil {
		return err
	}

	return nil
}
