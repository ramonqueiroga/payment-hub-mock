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
		transModel, err := ps.Cancel(req.PaymentID)
		if err != nil {
			return transport.CancelResponse{
				TransactionModel: business.TransactionModel{},
				Error:            errCancelPayment.Error(),
			}, nil
		}

		return transport.CancelResponse{
			TransactionModel: transModel,
			Error:            "",
		}, nil
	}
}

var errCancelPayment = errors.New("Error on canceling the payment")
