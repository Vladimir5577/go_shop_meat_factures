package handler

import (
	"net/http"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/helper"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/service"
)

type ProductHandlerInterface interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
}

type ProductHandler struct {
	productService service.ProductServiceInterface
}

func NewProductHandler(productService service.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (p *ProductHandler) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := p.productService.GetAllProducts()
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadGateway)
			return
		}
		helper.JsonResponse(w, resp, http.StatusOK)
	}
}
