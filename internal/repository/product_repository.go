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

func (p *ProductRepository) CountRows() (uint64, error) {
	var count uint64
	sqlQuery := "SELECT COUNT(id) FROM products "
	err := p.Db.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *ProductRepository) GetAllProducts() ([]model.ProductResponse, error) {
	var product model.ProductResponse
	var products []model.ProductResponse

	builderProduct := squirrel.Select("id", "name", "description", "price", "in_stock", "is_active", "created_at", "updated_at").
		From("products").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "is_active")), true).
		Join("categories WHERE (category_id = )").
		OrderBy(fmt.Sprintf("%s %s", "id", "ASC"))

	query, args, err := builderProduct.ToSql()
	if err != nil {
		return products, err
	}

	rows, err := p.Db.Query(query, args...)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.InStock, &product.IsActive, &product.Created, &product.Updated)
		if err != nil {
			return products, err
		}

		pr := model.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			InStock:     product.InStock,
			IsActive:    product.IsActive,
			Created:     product.Created,
			Updated:     product.Updated,
		}

		products = append(products, pr)
	}

	return products, nil
}
