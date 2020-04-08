package repository

import (
	"context"
	"database/sql"

	model "github.com/mikeleg/gokitexample/entity"

	"github.com/mikeleg/gokitexample/customer"
)

type storeCustomerRepository struct {
	Conn *sql.DB
}

func NewStoreCustomerRepository(Conn *sql.DB) customer.Repository {
	return &storeCustomerRepository{Conn}
}

func (m *storeCustomerRepository) GetByID(ctx context.Context, id uint) (res *model.Customer, err error) {

	return nil, nil
}
