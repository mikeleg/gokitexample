package customer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type GetCustomerByIDRequest struct {
	ID int `json:"id"`
}

type CustomerResponse Customer

func makeFetchAllCustomerEndpoint(svc CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customers, err := svc.FetchAll(ctx)
		return customers, err
	}
}

func makeFetchCustomerByIDEndpoint(svc CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIDRequest)
		customer, err := svc.FetchByID(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return CustomerResponse(customer), nil
	}
}
