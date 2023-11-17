package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"snapfood/constants/PrivateErrors"
	"snapfood/domain"
	"snapfood/repo/queries"
)

type TripsRepoInterface interface {
	TripsOrder(OrderID int) (*domain.Trip, error)
}

type TripsRepo struct {
	DB *sql.DB
}

func NewTripsRepo(DB *sql.DB) *TripsRepo {
	return &TripsRepo{DB: DB}
}

func (t *TripsRepo) TripsOrder(OrderID int) (*domain.Trip, error) {
	var trip domain.Trip
	err := t.DB.QueryRow(queries.FindOrderTrip, OrderID).Scan(&trip.TripID, &trip.OrderID, &trip.CourierID, &trip.TripStatus)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("error in find order trip: %v", err)
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, PrivateErrors.NotFound
	}
	return &trip, nil
}
