package artwork

import (
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type service struct {
	artworkRepo repository.ArtworkStorer
}

type Service interface {
	CreateArtwork(artworkDetails dto.CreateArtworkRequest) (repository.Artworks, error)
}

func NewService(artworkRepo repository.ArtworkStorer) Service {
	return &service{
		artworkRepo: artworkRepo,
	}
}

func (as *service) CreateArtwork(artworkDetails dto.CreateArtworkRequest) (artwork repository.Artworks, err error) {
	artworkInfo := repository.Artworks{
		Name:           artworkDetails.Name,
		Description:    artworkDetails.Description,
		Image:          artworkDetails.Image,
		Starting_price: artworkDetails.Starting_price,
		Live_period:    artworkDetails.Live_period,
		Owner_id:       artworkDetails.Owner_id,
	}
	return as.artworkRepo.CreateArtwork(artworkInfo)
}
