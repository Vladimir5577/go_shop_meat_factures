package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
)

type UserRepositoryInterface interface {
	Register(model.UserRegistration) (string, error)
	Login(model.UserRegistration) (string, error)
	NameExist(model.UserRegistration) (bool, error)
	PhoneExist(model.UserRegistration) (bool, error)
}

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) Register(user model.UserRegistration) (string, error) {
	if time.Now().Unix()%2 == 0 {
		return "Success", nil
	}
	return "Register page from reposityr", errors.New("something went wrong")
}

func (u *UserRepository) Login(user model.UserRegistration) (string, error) {
	return "Login page from repository", nil
}

func (u *UserRepository) NameExist(user model.UserRegistration) (bool, error) {
	var userExist model.User
	query, args, err := squirrel.Select("name").
		From("users").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "name")), user.Name).
		Limit(1).
		ToSql()
	if err != nil {
		return true, err
	}

	row := u.Db.QueryRow(query, args...)
	err = row.Scan(&userExist.Name)
	// if err != nil {
	// 	return true, err
	// }

	if userExist.Name == user.Name {
		return true, nil
	}

	return false, nil
}

func (u *UserRepository) PhoneExist(user model.UserRegistration) (bool, error) {
	var userExist model.User
	query, args, err := squirrel.Select("phone").
		From("users").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "phone")), user.Phone).
		Limit(1).
		ToSql()
	if err != nil {
		return true, err
	}

	row := u.Db.QueryRow(query, args...)
	err = row.Scan(&userExist.Phone)
	// if err != nil {
	// 	return true, err
	// }

	if userExist.Phone == user.Phone {
		return true, nil
	}

	return false, nil
}
