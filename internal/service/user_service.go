package service

import (
	"errors"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/helper"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(model.UserRegistration) (model.AuthResponse, error)
	Login(model.UserLogin) (model.AuthResponse, error)
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Register(user model.UserRegistration) (model.AuthResponse, error) {
	var authResponse model.AuthResponse
	nameExist, err := u.userRepository.NameExist(user)
	if err != nil {
		return authResponse, errors.New(err.Error())
	}

	if nameExist {
		return authResponse, errors.New("user with this name already exist")
	}

	phoneExist, err := u.userRepository.PhoneExist(user)

	if err != nil {
		return authResponse, errors.New(err.Error())
	}

	if phoneExist {
		return authResponse, errors.New("phone already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return authResponse, err
	}

	user.Password = string(hashedPassword)

	lastInsertedId, err := u.userRepository.Register(user)
	if err != nil {
		return authResponse, err
	}

	token, err := helper.NewJWT().Create(helper.JWTData{
		Id:   lastInsertedId,
		Name: user.Name,
	})
	if err != nil {
		return authResponse, err
	}

	authResponse = model.AuthResponse{
		Token: token,
	}

	return authResponse, nil
}

func (u *UserService) Login(user model.UserLogin) (model.AuthResponse, error) {
	var authResponse model.AuthResponse
	userLogined, err := u.userRepository.Login(user)
	if err != nil {
		return authResponse, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userLogined.Password), []byte(user.Password))
	if err != nil {
		return authResponse, errors.New("wrong credentials")
	}

	token, err := helper.NewJWT().Create(helper.JWTData{
		Id:   int64(userLogined.Id),
		Name: user.Name,
	})
	if err != nil {
		return authResponse, err
	}

	authResponse = model.AuthResponse{
		Token: token,
	}

	return authResponse, nil
}
