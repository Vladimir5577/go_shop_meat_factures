package repository

import (
	"database/sql"
	"errors"
)

type UserRepositoryInterface interface {
	Register() (string, error)
	Login() (string, error)
}

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) Register() (string, error) {
	return "Register page from reposityr", errors.New("something went wrong")
}

func (u *UserRepository) Login() (string, error) {
	return "Login page from repository", nil
}
