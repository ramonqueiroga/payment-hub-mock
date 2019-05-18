package external

import (
	"context"
	"errors"
	"payment-hub-mock/business"
	"payment-hub-mock/transport"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

//MakeSearchEndpoint creates the search endpoint
func MakeSearchEndpoint(ps business.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.SearchRequest)

		paymentID, err := strconv.ParseUint(req.PaymentID, 10, 64)
		if err != nil {
			return nil, err
		}

		trans, err := ps.Search(paymentID)

		if err != nil {
			return transport.SearchResponse{
				Transaction: business.Transaction{},
				Error:       errSearchPayment.Error(),
			}, nil
		}

		return transport.SearchResponse{
			Transaction: trans,
			Error:       "",
		}, nil
	}
}

var errSearchPayment = errors.New("Error searching the payment")
