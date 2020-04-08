package customer

import (
	"context"

	model "github.com/mikeleg/gokitexample/entity"
)

type Repository interface {
	GetByID(ctx context.Context, id uint) (model.Customer, error)
}
