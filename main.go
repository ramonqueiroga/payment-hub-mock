package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"payment-hub-mock/business"
	"payment-hub-mock/external"
	"payment-hub-mock/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	paymentService := business.PaymentServiceImpl{}

	captureHandler := httptransport.NewServer(
		external.MakeCaptureEndpoint(paymentService),
		decodeCaptureRequest,
		encodeCaptureResponse,
	)

	authorizeHandler := httptransport.NewServer(
		external.MakeAuthorizationEndpoint(paymentService),
		decodeAuthorizeRequest,
		encodeAuthorizeResponse,
	)

	http.Handle("/capture", captureHandler)
	http.Handle("/authorize", authorizeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//capture
func decodeCaptureRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.CaptureRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCaptureResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

//authorize
func decodeAuthorizeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.AuthorizeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeAuthorizeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
