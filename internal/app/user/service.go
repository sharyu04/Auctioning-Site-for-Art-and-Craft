package user

import (
	"errors"
	"regexp"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	CreateUser(userDetails dto.CreateUserRequest, role string) (dto.UserSignupResponse, error)
	LoginUser(credentials dto.LoginRequest) (string, error)
	GetAllUsers(start, count int, role string) ([]dto.GetAllUserResponse, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}

func createUserValidations(userDetails dto.CreateUserRequest) error {
	if userDetails.FirstName == "" || userDetails.LastName == "" || userDetails.Email == "" || userDetails.Password == "" {
		return errors.New("Invalid Input")

	}

	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matches := pattern.MatchString(userDetails.Email)

	if !matches {
		return errors.New("Invalid email")
	}

	sevenOrMore, number, upper, special := verifyPassword(userDetails.Password)

	if sevenOrMore == false || number == false || upper == false || special == false {
		return errors.New("Invalid password")
	}

	return nil
}

func (us *service) CreateUser(userDetails dto.CreateUserRequest, role string) (dto.UserSignupResponse, error) {

	err := createUserValidations(userDetails)
	if err != nil {
		return dto.UserSignupResponse{}, apperrors.BadRequest{ErrorMsg: err.Error()}
	}

	userInfo := repository.User{
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Email:     userDetails.Email,
		Password:  userDetails.Password,
	}

	err = us.userRepo.CheckEmailExists(userInfo)
	if err != nil {
		return dto.UserSignupResponse{}, apperrors.BadRequest{ErrorMsg: "User Already Exists"}
	}

	userInfo.Id = uuid.New()
	byte, _ := bcrypt.GenerateFromPassword([]byte(userInfo.Password), 14)
	userInfo.Password = string(byte)
	userInfo.Created_at = time.Now()
	if role == "" {
		role = "user"
	}
	userInfo.Role_id, err = us.userRepo.GetRoleID(role)
	if err != nil {
		return dto.UserSignupResponse{}, err
	}

	return us.userRepo.CreateUser(userInfo)
}

func (us *service) LoginUser(credentials dto.LoginRequest) (string, error) {

	expirationTime := time.Now().Add(time.Minute * 5)

	user, err := us.userRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", apperrors.BadRequest{ErrorMsg: "Invalid Credentials"}
	}

	claims := &dto.Claims{
		Id:   user.Id,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	var jwtKey = []byte("secret_key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (us *service) GetAllUsers(start, count int, role string) ([]dto.GetAllUserResponse, error) {

	if role != "" {
		if role != "admin" && role != "super_admin" && role != "user" {
			return nil, apperrors.BadRequest{ErrorMsg: "Invalid role"}
		}
		userList, err := us.userRepo.GetAllUsersByRole(start, count, role)
		if err != nil {
			return nil, err
		}
		if len(userList) == 0 {
			return userList, apperrors.NoContent{ErrorMsg: "No Users Found!"}
		}
		return userList, nil
	}

	userList, err := us.userRepo.GetAllUsers(start, count)
	if err != nil {
		return nil, err
	}

	if len(userList) == 0 {
		return userList, apperrors.NoContent{ErrorMsg: "No Users Found!"}
	}

	return userList, nil
}
