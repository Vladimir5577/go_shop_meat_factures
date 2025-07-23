package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/helper"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/service"
)

type OrderHandlerInterface interface {
	CreateOrder() http.HandlerFunc
}

type OrderHandler struct {
	orderService service.OrderServiceInterface
}

func NewOrderHandler(orderService service.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (o *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id, ok := r.Context().Value("user_id").(int64)
		if ok {
			fmt.Println("user id ", user_id)
		}
		if !ok {
			helper.JsonResponse(w, errors.New("require bearer token"), http.StatusUnauthorized)
			return
		}

		ordering, err := helper.HandleBody[model.Ordering](&w, r)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ordering.UserId = uint(user_id)

		resp, err := o.orderService.CreateOrder(*ordering)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, resp, http.StatusOK)
	}
}
