package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/V4"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
	"golang.org/x/crypto/bcrypt"
	// "github.com/dgrijalva/jwt-go"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	CreateUser(userDetails dto.CreateUserRequest) (repository.User, error)
	CreateAdmin(userDetails dto.CreateUserRequest) (repository.User, error)
	LoginUser(credentials dto.LoginRequest) (string, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (us *service) CreateUser(userDetails dto.CreateUserRequest) (user repository.User, err error) {
	if userDetails.FirstName == "" || userDetails.LastName == "" || userDetails.Email == "" || userDetails.Password == "" {
		return repository.User{}, errors.New("Invalid Input")
	}
	userInfo := repository.User{
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Email:     userDetails.Email,
		Password:  userDetails.Password,
	}
	return us.userRepo.CreateUser(userInfo)
}

func (us *service) CreateAdmin(userDetails dto.CreateUserRequest) (user repository.User, err error) {
	if userDetails.FirstName == "" || userDetails.LastName == "" || userDetails.Email == "" || userDetails.Password == "" {
		return repository.User{}, errors.New("Invalid Input")
	}
	userInfo := repository.User{
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Email:     userDetails.Email,
		Password:  userDetails.Password,
	}
	return us.userRepo.CreateAdmin(userInfo)
}

func (us *service) LoginUser(credentials dto.LoginRequest) (string, error) {

	expirationTime := time.Now().Add(time.Minute * 5)

	user, err := us.userRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		err = errors.New("Invalid credentials")
		return "", err
	}

	claims := &dto.Claims{
		Id:   user.Id,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	var jwtKey = []byte("secret_key")

	// token, err := getToken(credentials.Email)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
