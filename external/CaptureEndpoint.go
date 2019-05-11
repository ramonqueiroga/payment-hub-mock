package external

import (
	"context"
	"errors"
	"payment-hub-mock/business"
	"payment-hub-mock/transport"

	"github.com/go-kit/kit/endpoint"
)

//MakeCaptureEndpoint create the capture payment endpoint
func MakeCaptureEndpoint(ps business.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.CaptureRequest)
		transModel, err := ps.Capture(req.Payments)

		if err != nil {
			return transport.CaptureResponse{
				TransactionModel: business.TransactionModel{},
				Error:            errCapturePayment.Error(),
			}, nil
		}

		return transport.CaptureResponse{
			TransactionModel: transModel,
			Error:            "",
		}, nil
	}
}

var errCapturePayment = errors.New("Error on capture the payment")
