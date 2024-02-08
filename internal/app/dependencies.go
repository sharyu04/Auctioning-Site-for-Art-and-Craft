package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/artwork"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/user"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type Dependencies struct {
	UserService    user.Service
	ArtworkService artwork.Service
}

func NewServices(db *sqlx.DB) Dependencies {
	userRepo := repository.NewUserRepo(db)
	artworkRepo := repository.NewArtworkRepo(db)

	userService := user.NewService(userRepo)
	artworkService := artwork.NewService(artworkRepo)

	return Dependencies{
		UserService:    userService,
		ArtworkService: artworkService,
	}
}
