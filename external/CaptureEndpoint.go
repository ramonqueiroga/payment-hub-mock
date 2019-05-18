package external

import (
	"context"
	"errors"
	"payment-hub-mock/business"
	"payment-hub-mock/transport"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

//MakeCaptureEndpoint create the capture payment endpoint
func MakeCaptureEndpoint(ps business.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.CaptureRequest)

		paymentID, err := strconv.ParseUint(req.PaymentID, 10, 64)
		if err != nil {
			return nil, err
		}

		trans, err := ps.Capture(paymentID)

		//aqui depois da captura, precisamos persistir no banco
		if err != nil {
			return transport.CaptureResponse{
				Transaction: business.Transaction{},
				Error:       errCapturePayment.Error(),
			}, nil
		}

		return transport.CaptureResponse{
			Transaction: trans,
			Error:       "",
		}, nil
	}
}

var errCapturePayment = errors.New("Error on capture the payment")
