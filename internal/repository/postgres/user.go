package repository

// import (
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/repository"
// 	"golang.org/x/crypto/bcrypt"
// )

// type userStore struct {
// 	DB *sqlx.DB
// }

// func NewOrderRepo(db *sqlx.DB) repository.UserStorer {
// 	return &userStore{
// 		DB: db,
// 	}
// }

// func (us *userStore) CreateUser(user repository.User) (repository.User, error) {
// 	user.Id = uuid.New()
// 	byte, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
// 	user.Password = string(byte)
// 	user.Created_at = time.Now()
// 	user.Role_id, _ = uuid.Parse("6a55565e-3b0f-48fe-854e-ea22ce1ff991")
// 	err := us.DB.QueryRow("INSERT INTO users(id, firstname, lastname, email, password, created_at, role_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
// 		user.Id, user.FirstName, user.LastName, user.Email, user.Password, user.Created_at, user.Role_id).Scan(&user.Id)

// 	if err != nil {
// 		return repository.User{}, err
// 	}
// 	return user, nil
// }
