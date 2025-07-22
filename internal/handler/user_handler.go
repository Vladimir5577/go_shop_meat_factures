package handler

import (
	"fmt"
	"net/http"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/helper"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
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
		u, err := helper.HandleBody[model.UserRegistration](&w, r)
		if err != nil {
			return
		}

		fmt.Printf("Received person: Name=%v, password=%v\n", u.Name, u.Password)

		resp, err := h.userService.Register(*u)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, resp, http.StatusCreated)
	}
}

func (h *UserHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.UserRegistration
		resp, err := h.userService.Login(user)
		if err != nil {
			return
		}
		w.Write([]byte(resp))
	}
}
