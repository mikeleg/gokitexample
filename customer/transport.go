package customer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func FetchAllCustomerHandelr(cs CustomerService, logger kitlog.Logger) *kithttp.Server {
	return kithttp.NewServer(
		makeFetchAllCustomerEndpoint(cs),
		kithttp.NopRequestDecoder,
		encodeResponse,
	)
}

func FetchCustomerByIDHandler(cs CustomerService, logger kitlog.Logger) *kithttp.Server {
	return kithttp.NewServer(
		makeFetchCustomerByIDEndpoint(cs),
		decodeGetByIDRequest,
		encodeResponse,
	)
}

var errBadRoute = errors.New("bad route")

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["ID"]
	if !ok {
		return nil, errBadRoute
	}
	intID, _ := strconv.Atoi(id)
	return GetCustomerByIDRequest{ID: intID}, nil
}

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
