package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"snapfood/constants/PrivateErrors"
	"snapfood/domain"
	"snapfood/repo/queries"
)

type OrdersRepoInterface interface {
	GetOrders(OrderID int) (*domain.Order, error)
}

type OrdersRepo struct {
	DB *sql.DB
}

func NewOrdersRepo(DB *sql.DB) *OrdersRepo {
	return &OrdersRepo{DB: DB}
}

func (t *OrdersRepo) GetOrders(OrderID int) (*domain.Order, error) {
	var order domain.Order
	err := t.DB.QueryRow(queries.FindOrders, OrderID).Scan(&order.OrderID, &order.UserID, &order.VendorID, &order.OrderTime, &order.TimeDelivery)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("error in find order: %v", err)
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, PrivateErrors.NotFound
	}
	return &order, nil
}
