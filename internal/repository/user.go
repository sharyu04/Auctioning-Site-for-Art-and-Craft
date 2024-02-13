package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

type UserStorer interface {
	CreateUser(user User) (dto.UserSignupResponse, error)
	GetUserByEmail(reqEmail string) (dto.User, error)
	CheckEmailExists(user User) error
	GetRoleID(role string) (uuid.UUID, error)
	GetAllUsers(start, count int) ([]dto.GetAllUserResponse, error)
	GetAllUsersByRole(start, count int, role string) ([]dto.GetAllUserResponse, error)
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
	id   uuid.UUID `db:id`
	name string    `db:name`
}

type userStore struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserStorer {
	return &userStore{
		DB: db,
	}
}

func (us *userStore) CheckEmailExists(user User) error {
	rows, err := us.DB.Query("Select * from  users where email=$1", user.Email)
	if err != nil {
		return err
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		err = errors.New("User with this email id already exists!")
		return err
	}

	return nil
}

func (us *userStore) CreateUser(user User) (dto.UserSignupResponse, error) {

	err := us.DB.QueryRow("INSERT INTO users(id, firstname, lastname, email, password, created_at, role_id) VALUES($1, $2, $3, $4, $5, $6,$7) RETURNING id",
		user.Id, user.FirstName, user.LastName, user.Email, user.Password, user.Created_at, user.Role_id).Scan(&user.Id)

	if err != nil {
		return dto.UserSignupResponse{}, err
	}

	resUser := dto.UserSignupResponse{
		Id:         user.Id,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Created_at: user.Created_at,
		Role_id:    user.Role_id,
	}

	return resUser, nil
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

func (us *userStore) GetRoleID(role string) (uuid.UUID, error) {
	var roleId uuid.UUID
	err := us.DB.QueryRow("Select id from role where name=$1", role).Scan(&roleId)
	if err != nil {
		fmt.Println("Error here ", role)
		return uuid.Nil, err
	}
	return roleId, nil
}

func (us *userStore) GetAllUsers(start, count int) ([]dto.GetAllUserResponse, error) {
	rows, err := us.DB.Query("select u.id, u.firstname, u.lastname, u.email, u.created_at, role.name from users as u join role on u.role_id=role.id Limit $1 offset $2;", count, start)
	if err != nil {
		return nil, err
	}

	var userList []dto.GetAllUserResponse
	for rows.Next() {
		var user dto.GetAllUserResponse
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.RoleID)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}

func (us *userStore) GetAllUsersByRole(start, count int, role string) ([]dto.GetAllUserResponse, error) {
	rows, err := us.DB.Query("select u.id, u.firstname, u.lastname, u.email, u.created_at, role.name from users as u join role on u.role_id=role.id where role.name=$1 Limit $2 offset $3;", role, count, start)
	if err != nil {
		return nil, err
	}

	var userList []dto.GetAllUserResponse
	for rows.Next() {
		var user dto.GetAllUserResponse
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.RoleID)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}
