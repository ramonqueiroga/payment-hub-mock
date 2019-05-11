package external

import (
	"context"
	"errors"
	"payment-hub-mock/business"
	"payment-hub-mock/transport"

	"github.com/go-kit/kit/endpoint"
)

//MakeAuthorizationEndpoint creates the authorization payment endpoint
func MakeAuthorizationEndpoint(ps business.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.AuthorizeRequest)
		transModel, err := ps.Authorize(req.PaymentID)

		if err != nil {
			return transport.AuthorizeResponse{
				TransactionModel: business.TransactionModel{},
				Error:            errAuthorization.Error(),
			}, nil
		}

		return transport.AuthorizeResponse{
			TransactionModel: transModel,
			Error:            "",
		}, nil
	}
}

var errAuthorization = errors.New("Error on authorize the payment")
