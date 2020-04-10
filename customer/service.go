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

func NewCustomerService(db *sqlx.DB) Service {
	return &service{
		db: db,
	}
}

func (s *customerservice) FetchByID(ctx context.Context, id int) (Customer, error) {
	customer := Customer{}

	err := s.db.Get(&customer, "SELECT * FROM customer where id=$", id)
	if err != nil {
		return nil, err
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
