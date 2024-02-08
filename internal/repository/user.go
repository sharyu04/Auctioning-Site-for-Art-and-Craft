package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserStorer interface {
	CreateUser(user User) (User, error)
}

type User struct {
	Id         uuid.UUID `db:id`
	FirstName  string    `db:firstname`
	LastName   string    `db:lastname`
	Email      string    `db:email`
	Password   string    `db:password`
	Role_id    uuid.UUID `db:role_id`
	Created_at time.Time `db:created_at`
}

type userStore struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserStorer {
	return &userStore{
		DB: db,
	}
}

func (us *userStore) CreateUser(user User) (User, error) {
	user.Id = uuid.New()
	byte, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(byte)
	user.Created_at = time.Now()
	user.Role_id, _ = uuid.Parse("6a55565e-3b0f-48fe-854e-ea22ce1ff991")
	err := us.DB.QueryRow("INSERT INTO users(id, firstname, lastname, email, password, created_at, role_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.Id, user.FirstName, user.LastName, user.Email, user.Password, user.Created_at, user.Role_id).Scan(&user.Id)

	if err != nil {
		return User{}, err
	}
	return user, nil
}
