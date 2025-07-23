package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
)

type UserRepositoryInterface interface {
	Register(model.UserRegistration) (int64, error)
	Login(model.UserLogin) (model.UserLogin, error)
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

func (u *UserRepository) Register(user model.UserRegistration) (int64, error) {
	query, args, err := squirrel.Insert("users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("name", "password", "phone", "address").
		Values(user.Name, user.Password, user.Phone, user.Address).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var lastInsertedId int64
	err = u.Db.QueryRow(query, args...).Scan(&lastInsertedId)
	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}

func (u *UserRepository) Login(user model.UserLogin) (model.UserLogin, error) {
	var userLogined model.UserLogin
	query, args, err := squirrel.Select("id", "name", "password", "phone").
		From("users").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "phone")), user.Phone).
		Limit(1).
		ToSql()
	if err != nil {
		return userLogined, err
	}

	row := u.Db.QueryRow(query, args...)
	err = row.Scan(&userLogined.Id, &userLogined.Name, &userLogined.Password, &userLogined.Phone)
	if err != nil {
		return userLogined, errors.New("wrong credentials")
	}
	return userLogined, nil
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
	_ = row.Scan(&userExist.Name)
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
	_ = row.Scan(&userExist.Phone)
	// if err != nil {
	// 	return true, err
	// }

	if userExist.Phone == user.Phone {
		return true, nil
	}

	return false, nil
}
