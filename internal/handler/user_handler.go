package handler

import (
	"net/http"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/service"
)

type UserHandlerInterface interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
}

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.userService.Register()
		if err != nil {
			return
		}
		w.Write([]byte(resp))
	}
}

func (h *UserHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.userService.Login()
		if err != nil {
			return
		}
		w.Write([]byte(resp))
	}
}
