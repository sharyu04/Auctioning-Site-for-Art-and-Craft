package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/user"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type Dependencies struct {
	UserService user.Service
}

func NewServices(db *sqlx.DB) Dependencies {
	userRepo := repository.NewUserRepo(db)

	userService := user.NewService(userRepo)

	return Dependencies{
		UserService: userService,
	}
}