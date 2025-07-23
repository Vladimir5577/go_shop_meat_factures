package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
)

type OrderRepositoryInterface interface {
	CreateOrder(model.Ordering) (string, error)
	CheckProductExist(uint) (bool, error)
}

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (o *OrderRepository) CreateOrder(ordering model.Ordering) (string, error) {

	query, args, err := squirrel.Insert("orders").
		PlaceholderFormat(squirrel.Dollar).
		Columns("name", "password", "phone", "address").
		Values(user.Name, user.Password, user.Phone, user.Address).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var lastInsertedId int64
	err = u.Db.QueryRow(query, args...).Scan(&lastInsertedId)
	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil

	for _, pr := range ordering.Products {

	}
	return "Hello from order repository!", nil
}

func (p *OrderRepository) CheckProductExist(id uint) (bool, error) {
	var product model.ProductResponse
	query, args, err := squirrel.Select("id").
		From("products").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "id")), id).
		Limit(1).
		ToSql()
	if err != nil {
		return false, err
	}

	row := p.Db.QueryRow(query, args...)
	err = row.Scan(&product.Id)
	if err != nil {
		return false, err
	}

	if product.Id != id {
		return false, errors.New("product does not exist")
	}

	return true, nil
}
