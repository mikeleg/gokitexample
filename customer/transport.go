package customer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeFetchAllCustomer(cs CustomerService, logger kitlog.Logger) *kithttp.Server {
	return kithttp.NewServer(
		makeFetchAllCustomerEndpoint(cs),
		kithttp.NopRequestDecoder,
		encodeResponse,
	)
}

var errBadRoute = errors.New("bad route")

type errorer interface {
	error() error
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	kithttp.EncodeJSONResponse(ctx, w, err)
	return
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		errorEncoder(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
