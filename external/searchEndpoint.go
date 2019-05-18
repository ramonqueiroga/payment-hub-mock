package external

import (
	"context"
	"errors"
	"payment-hub-mock/business"
	"payment-hub-mock/transport"

	"github.com/go-kit/kit/endpoint"
)

//MakeSearchEndpoint creates the search endpoint
func MakeSearchEndpoint(ps business.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.SearchRequest)
		trans, err := ps.Search(req.PaymentID)

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
