package user

import (
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	CreateUser(userDetails dto.CreateUserRequest) (repository.User, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (us *service) CreateUser(userDetails dto.CreateUserRequest) (user repository.User, err error) {
	userInfo := repository.User{
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Email:     userDetails.Email,
		Password:  userDetails.Password,
	}
	return us.userRepo.CreateUser(userInfo)
}
