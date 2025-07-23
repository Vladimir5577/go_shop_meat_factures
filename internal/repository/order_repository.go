package repository

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/model"
)

type OrderRepositoryInterface interface {
	CreateOrder(model.Ordering) (string, error)
	GetProductById(uint) (model.ProductResponse, error)
	GetOrdersByUser(uint) ([]model.OrderResponse, error)
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
		Columns("user_id", "total_summ", "status", "comment").
		Values(ordering.UserId, ordering.TotalSumm, ordering.Status, ordering.Comment).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return "", err
	}

	var orderInsertedId int64
	err = o.Db.QueryRow(query, args...).Scan(&orderInsertedId)
	if err != nil {
		return "", err
	}

	for _, pr := range ordering.Products {
		query, args, err := squirrel.Insert("order_products").
			PlaceholderFormat(squirrel.Dollar).
			Columns("order_id", "product_id", "amount", "total_summ").
			Values(orderInsertedId, pr.ProductId, pr.Amount, pr.SummItem).
			Suffix("RETURNING id").
			ToSql()
		if err != nil {
			return "", err
		}
		err = o.Db.QueryRow(query, args...).Err()
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("Order with id = %v created successfully! Status - %v", orderInsertedId, ordering.Status), nil
}

func (p *OrderRepository) GetProductById(id uint) (model.ProductResponse, error) {
	var product model.ProductResponse
	var category model.ProductCategory

	query, args, err := squirrel.Select("products.id", "products.name", "products.description", "products.price", "products.in_stock", "products.is_active", "products.created_at", "products.updated_at", "categories.id", "categories.name as category_name").
		From("products").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "products.id")), id).
		Join("categories on products.category_id = categories.id").
		ToSql()

	if err != nil {
		return product, err
	}

	row := p.Db.QueryRow(query, args...)
	err = row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.InStock, &product.IsActive, &product.Created, &product.Updated, &category.Id, &category.Name)
	if err != nil {
		return product, fmt.Errorf("product with id = %v does not exeist", id)
	}

	product.Category = category

	return product, nil
}

func (o *OrderRepository) GetOrdersByUser(id uint) ([]model.OrderResponse, error) {
	var order model.OrderResponse
	var orders []model.OrderResponse

	_, err := o.CheckUserWithIdExists(id)
	if err != nil {
		return orders, fmt.Errorf("user with id = %v does not exist", id)
	}

	_, err = o.CheckUserHaveOrders(id)
	if err != nil {
		return orders, err
	}

	builderProduct := squirrel.Select(
		"orders.id",
		"orders.user_id",
		"orders.total_summ",
		"orders.status",
		"orders.comment",
		"orders.created_at",
		"orders.updated_at",
	).
		From("orders").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "orders.user_id")), id).
		OrderBy(fmt.Sprintf("%s %s", "orders.id", "DESC"))

	query, args, err := builderProduct.ToSql()

	if err != nil {
		return orders, err
	}

	rows, err := o.Db.Query(query, args...)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.UserId, &order.TotalSumm, &order.Status, &order.Comment, &order.Created, &order.Updated)
		if err != nil {
			return orders, err
		}

		products, err := o.GetProductsByOrder(order.Id)
		if err != nil {
			return orders, err
		}

		or := model.OrderResponse{
			Id:        order.Id,
			UserId:    order.UserId,
			Products:  products,
			TotalSumm: order.TotalSumm,
			Status:    order.Status,
			Comment:   order.Comment,
			Created:   order.Created,
			Updated:   order.Updated,
		}

		orders = append(orders, or)
	}

	return orders, nil
}

func (p *OrderRepository) GetProductsByOrder(orderId uint) ([]model.OrderItemResponse, error) {
	var product model.OrderItemProductResponse
	var orderItem model.OrderItemResponse
	var orderItems []model.OrderItemResponse

	builderProduct := squirrel.Select(
		"order_products.id",
		"order_products.order_id",
		"order_products.amount",
		"order_products.total_summ",

		"products.id",
		"products.name",
		"products.price",
		"products.in_stock").
		From("order_products").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "order_products.order_id")), orderId).
		Join("products on order_products.product_id = products.id")

	query, args, err := builderProduct.ToSql()
	if err != nil {
		return orderItems, err
	}

	rows, err := p.Db.Query(query, args...)
	if err != nil {
		return orderItems, err
	}

	for rows.Next() {
		err = rows.Scan(
			&orderItem.Id,
			&orderItem.OrderId,
			&orderItem.Amount,
			&orderItem.SummItem,

			&product.ProductId,
			&product.ProductName,
			&product.Price,
			&product.InStock)
		if err != nil {
			return orderItems, err
		}

		or := model.OrderItemResponse{
			Id:       orderItem.Id,
			OrderId:  orderItem.OrderId,
			Product:  product,
			Amount:   orderItem.Amount,
			SummItem: orderItem.SummItem,
			InStock:  orderItem.InStock,
		}

		orderItems = append(orderItems, or)
	}

	return orderItems, nil
}

func (o *OrderRepository) CheckUserWithIdExists(id uint) (bool, error) {
	var userExistId uint
	query, args, err := squirrel.Select("id").
		From("users").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "id")), id).
		Limit(1).
		ToSql()
	if err != nil {
		return false, err
	}

	row := o.Db.QueryRow(query, args...)
	err = row.Scan(&userExistId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (o *OrderRepository) CheckUserHaveOrders(id uint) (bool, error) {
	var count uint
	sqlQuery := fmt.Sprintf("SELECT COUNT(id) FROM orders WHERE user_id = %v", id)
	err := o.Db.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, fmt.Errorf("user with id = %v does not have orders", id)
}
