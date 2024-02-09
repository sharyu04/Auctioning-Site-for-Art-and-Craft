package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/artwork"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/bid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/user"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type Dependencies struct {
	UserService    user.Service
	ArtworkService artwork.Service
	BidService     bid.Service
}

func NewServices(db *sqlx.DB) Dependencies {
	userRepo := repository.NewUserRepo(db)
	artworkRepo := repository.NewArtworkRepo(db)
	bidRepo := repository.NewBidRepo(db)

	userService := user.NewService(userRepo)
	artworkService := artwork.NewService(artworkRepo)
	bidService := bid.NewService(bidRepo)

	return Dependencies{
		UserService:    userService,
		ArtworkService: artworkService,
		BidService:     bidService,
	}
}
