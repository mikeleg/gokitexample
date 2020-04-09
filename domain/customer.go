package entity

type Customer struct {
	ID        uint   `json:"id,omitempty"`
	Legalname string `json:"legalname"`
}

type CustomerRepository interface {
	GetByID(ctx context.Context, id uint) (model.Customer, error)
	GetAll(ctx context.Context) ([]model.Customer, error)
}