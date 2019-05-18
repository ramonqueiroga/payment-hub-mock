package external

import (
	"context"
	"errors"
	"payment-hub-mock/business"
	"payment-hub-mock/transport"

	"github.com/go-kit/kit/endpoint"
)

//MakeCancelEndpoint creates the cancel payment endpoint
func MakeCancelEndpoint(ps business.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.CancelRequest)
		trans, err := ps.Cancel(req.PaymentID)
		if err != nil {
			return transport.CancelResponse{
				Transaction: business.Transaction{},
				Error:       errCancelPayment.Error(),
			}, nil
		}

		return transport.CancelResponse{
			Transaction: trans,
			Error:       "",
		}, nil
	}
}

var errCancelPayment = errors.New("Error on canceling the payment")
