package service

import "github.com/Vladimir5577/go_shop_meat_factures/internal/repository"

type UserServiceInterface interface {
	Register() (string, error)
	Login() (string, error)
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Register() (string, error) {
	return u.userRepository.Register()
}

func (u *UserService) Login() (string, error) {
	return u.userRepository.Login()
}
