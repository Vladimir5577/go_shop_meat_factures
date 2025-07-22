package service

import (
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/repository"
)

type ProductServiceInterface interface {
	GetAllProducts() ([]model.ProductResponse, error)
}

type ProductService struct {
	productRepository repository.ProductRepositoryInterface
}

func NewProductService(productRepository repository.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) GetAllProducts() ([]model.ProductResponse, error) {
	// var products model.ProductResponse
	resp, err := p.productRepository.GetAllProducts()
	if err != nil {
		return resp, err
	}

	return resp, err
}
