package service

import (
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/repository"
)

type OrderServiceInterface interface {
	CreateOrder(model.Ordering) (string, error)
	GetOrdersByUser(uint) ([]model.OrderResponse, error)
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
	for k, pr := range ordering.Products {
		product, err := o.orderRepository.GetProductById(pr.ProductId)
		if err != nil {
			return "", err
		}

		summItem := float32(pr.Amount) * product.Price
		ordering.Products[k].SummItem = summItem
		ordering.TotalSumm += summItem
	}

	ordering.Status = "new"

	resp, err := o.orderRepository.CreateOrder(ordering)
	// if err != nil {
	// 	return resp, err
	// }

	return resp, err
}

func (o *OrderService) GetOrdersByUser(userId uint) ([]model.OrderResponse, error) {
	return o.orderRepository.GetOrdersByUser(userId)
}
