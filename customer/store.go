package customer

import (
	"github.com/go-kit/kit/log"
	"context"
	"database/sql"

	model "github.com/mikeleg/gokitexample/domain"

	"github.com/mikeleg/gokitexample/customer"
)

type storeCustomerRepository struct {
	db *sql.DB
	logger log.Logger
}

func New(db *sql.DB, logger log.Logger) customer.Repository {
	return &storeCustomerRepository{
		db:     db,
		logger: logger,
	}
}

func (m *storeCustomerRepository) GetByID(ctx context.Context, id uint) (res *model.Customer, err error) {

	return nil, nil
}

func (m *storeCustomerRepository) GetAll(ctx context.Context) (res []model.Customer, err error) {

	return [], nil
}
