package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"

	// "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"

	// "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserStorer interface {
	CreateUser(user User) (User, error)
	CreateAdmin(user User) (User, error)
	GetUserByEmail(reqEmail string) (dto.User, error)
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

type Role struct {
	id   uuid.UUID
	name string
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
	rows, err := us.DB.Query("Select * from  users where email=$1", user.Email)
	if err != nil {
		return User{}, err
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		err = errors.New("User with this email id already exists!")
		return User{}, err
	}

	user.Id = uuid.New()
	byte, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(byte)
	user.Created_at = time.Now()
	user.Role_id, _ = uuid.Parse("6a55565e-3b0f-48fe-854e-ea22ce1ff991")
	err = us.DB.QueryRow("INSERT INTO users(id, firstname, lastname, email, password, created_at, role_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.Id, user.FirstName, user.LastName, user.Email, user.Password, user.Created_at, user.Role_id).Scan(&user.Id)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (us *userStore) CreateAdmin(user User) (User, error) {
	rows, err := us.DB.Query("Select * from  users where email=$1", user.Email)
	if err != nil {
		return User{}, err
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		err = errors.New("User with this email id already exists!")
		return User{}, err
	}

	user.Id = uuid.New()
	byte, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(byte)
	user.Created_at = time.Now()
	user.Role_id, _ = uuid.Parse("d832c38a-ce9b-43ae-8755-0a66f669acf2")
	err = us.DB.QueryRow("INSERT INTO users(id, firstname, lastname, email, password, created_at, role_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.Id, user.FirstName, user.LastName, user.Email, user.Password, user.Created_at, user.Role_id).Scan(&user.Id)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (us *userStore) GetUserByEmail(reqEmail string) (dto.User, error) {
	rows, err := us.DB.Query("Select * from  users where email=$1", reqEmail)
	if err != nil {
		return dto.User{}, err
	}

	var user dto.User
	defer rows.Close()
	for rows.Next() {
		var roleId uuid.UUID
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Created_at, &roleId)
		if err != nil {
			return dto.User{}, err
		}

		err = us.DB.QueryRow("Select name from role where id=$1", roleId).Scan(&user.Role)
		if err != nil {
			return dto.User{}, err
		}

	}
	return user, nil
}
