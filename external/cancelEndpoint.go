package external

import (
	"context"
	"errors"
	"payment-hub-mock/business"
	"payment-hub-mock/transport"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

//MakeCancelEndpoint creates the cancel payment endpoint
func MakeCancelEndpoint(ps business.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.CancelRequest)

		paymentID, err := strconv.ParseUint(req.PaymentID, 10, 64)
		if err != nil {
			return nil, err
		}

		trans, err := ps.Cancel(paymentID)
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
