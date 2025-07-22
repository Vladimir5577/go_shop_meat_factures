package service

import (
	"errors"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(model.UserRegistration) (string, error)
	Login(model.UserRegistration) (string, error)
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Register(user model.UserRegistration) (string, error) {
	nameExist, err := u.userRepository.NameExist(user)
	if err != nil {
		return "", errors.New(err.Error())
	}

	if nameExist {
		return "", errors.New("user with this name already exist")
	}

	phoneExist, err := u.userRepository.PhoneExist(user)

	if err != nil {
		return "", errors.New(err.Error())
	}

	if phoneExist {
		return "", errors.New("phone already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)

	return u.userRepository.Register(user)
}

func (u *UserService) Login(user model.UserRegistration) (string, error) {
	return u.userRepository.Login(user)
}
