package repository

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
)

type ProductRepositoryInterface interface {
	GetAllProducts() ([]model.ProductResponse, error)
}

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}

func (p *ProductRepository) GetAllProducts() ([]model.ProductResponse, error) {
	var product model.ProductResponse
	var category model.ProductCategory
	var products []model.ProductResponse

	builderProduct := squirrel.Select("products.id", "products.name", "products.description", "products.price", "products.in_stock", "products.is_active", "products.created_at", "products.updated_at", "categories.id", "categories.name as category_name").
		From("products").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "products.is_active")), true).
		Join("categories on products.category_id = categories.id").
		OrderBy(fmt.Sprintf("%s %s", "products.id", "ASC"))

	query, args, err := builderProduct.ToSql()
	if err != nil {
		return products, err
	}

	rows, err := p.Db.Query(query, args...)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.InStock, &product.IsActive, &product.Created, &product.Updated, &category.Id, &category.Name)
		if err != nil {
			return products, err
		}

		pr := model.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Category:    category,
			InStock:     product.InStock,
			IsActive:    product.IsActive,
			Created:     product.Created,
			Updated:     product.Updated,
		}

		products = append(products, pr)
	}

	return products, nil
}
