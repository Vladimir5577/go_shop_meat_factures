package service

import (
	"errors"
	"fmt"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/repository"
)

type OrderServiceInterface interface {
	CreateOrder(model.Ordering) (string, error)
}

type OrderService struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewOrderService(orderRepository repository.OrderRepositoryInterface) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (o *OrderService) CreateOrder(ordering model.Ordering) (string, error) {
	for _, pr := range ordering.Products {
		exist, err := o.orderRepository.CheckProductExist(pr.ProductId)
		if !exist {
			return "", errors.New(fmt.Sprintf("Product with id = %v does not exist!", pr.ProductId))
		}
		if err != nil {
			return "", err
		}
		fmt.Println(pr.ProductId)
	}

	resp, err := o.orderRepository.CreateOrder(ordering)
	// if err != nil {
	// 	return resp, err
	// }

	return resp, err
}
