package customer

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type CustomerService interface {
	FetchByID(ctx context.Context, id int) (Customer, error)
	FetchAll(ctx context.Context) ([]Customer, error)
}

type customerservice struct {
	db *sqlx.DB
}

func NewCustomerService(db *sqlx.DB) CustomerService {
	return &customerservice{
		db: db,
	}
}

func (s *customerservice) FetchByID(ctx context.Context, id int) (Customer, error) {
	customer := Customer{}

	err := s.db.Get(&customer, `SELECT * FROM customer where id = $1`, id)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (s *customerservice) FetchAll(ctx context.Context) ([]Customer, error) {
	customers := []Customer{}
	err := s.db.Select(&customers, "SELECT * FROM customer")
	if err != nil {
		return customers, err
	}
	return customers, nil
}
