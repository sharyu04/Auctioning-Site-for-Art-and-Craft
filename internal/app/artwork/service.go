package artwork

import (
	"time"

	"github.com/google/uuid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type service struct {
	artworkRepo repository.ArtworkStorer
}

type Service interface {
	CreateArtwork(artworkDetails dto.CreateArtworkRequest) (repository.Artworks, error)
	GetArtworks(category string, start int, count int) ([]dto.GetArtworkResponse, error)
	GetArtworkByID(id string) (dto.GetArtworkResponse, error)
	DeleteArtworkById(id string, owner_id string, role string) error
}

func NewService(artworkRepo repository.ArtworkStorer) Service {
	return &service{
		artworkRepo: artworkRepo,
	}
}

func (as *service) CreateArtwork(artworkDetails dto.CreateArtworkRequest) (repository.Artworks, error) {

	if artworkDetails.Name == "" || artworkDetails.Description == "" || artworkDetails.Image == "" || artworkDetails.Category == "" || artworkDetails.Duration == 0 {
		return repository.Artworks{}, apperrors.BadRequest{ErrorMsg: "Invalid Credentials"}
	}

	category, err := as.artworkRepo.GetCategory(artworkDetails.Category)
	if err != nil {
		return repository.Artworks{}, err
	}
	owner, _ := uuid.Parse(artworkDetails.Owner_id)
	closing_time := time.Now().Add(artworkDetails.Duration * time.Hour)
	artworkInfo := repository.Artworks{
		Name:           artworkDetails.Name,
		Description:    artworkDetails.Description,
		Image:          artworkDetails.Image,
		Starting_price: artworkDetails.Starting_price,
		Closing_time:   closing_time,
		Owner_id:       owner,
		Category_id:    category.Id,
	}

	return as.artworkRepo.CreateArtwork(artworkInfo)
}

func (as *service) GetArtworks(category string, start int, count int) ([]dto.GetArtworkResponse, error) {

	if category != "" {
		_, err := as.artworkRepo.GetCategory(category)
		if err != nil {
			return nil, err
		}

		artworkList, err := as.artworkRepo.GetFilterArtworks(category, start, count)
		if err != nil {
			return nil, err
		}
		if len(artworkList) == 0 {
			return nil, apperrors.NoContent{ErrorMsg: "No artworks found!"}
		}
		return artworkList, nil

	} else {
		artworkList, err := as.artworkRepo.GetAllArtworks(start, count)
		if err != nil {
			return nil, err
		}
		if len(artworkList) == 0 {
			return nil, apperrors.NoContent{ErrorMsg: "No artworks found!"}
		}
		return artworkList, nil
	}

}

func (as *service) GetArtworkByID(id string) (dto.GetArtworkResponse, error) {
	artworkId, err := uuid.Parse(id)
	if err != nil {
		return dto.GetArtworkResponse{}, err
	}
	return as.artworkRepo.GetArtworkById(artworkId)
}

func (as *service) DeleteArtworkById(id string, owner_id string, role string) error {
	artworkId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	ownerId, err := uuid.Parse(owner_id)
	if err != nil {
		return err
	}

	err = as.artworkRepo.DeleteArtworkById(artworkId, ownerId, role)
	if err != nil {
		return err
	}

	return nil

}
