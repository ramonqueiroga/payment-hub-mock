package main

import (
	"log"
	"net/http"
	"payment-hub-mock/business"
	"payment-hub-mock/external"
	"payment-hub-mock/serializable"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	paymentService := business.PaymentServiceImpl{}

	captureHandler := httptransport.NewServer(
		external.MakeCaptureEndpoint(paymentService),
		serializable.DecodeCaptureRequest,
		serializable.EncodeCaptureResponse,
	)

	authorizeHandler := httptransport.NewServer(
		external.MakeAuthorizationEndpoint(paymentService),
		serializable.DecodeAuthorizeRequest,
		serializable.EncodeAuthorizeResponse,
	)

	cancelHandler := httptransport.NewServer(
		external.MakeCancelEndpoint(paymentService),
		serializable.DecodeCancelRequest,
		serializable.EncodeCancelResponse,
	)

	searchHandler := httptransport.NewServer(
		external.MakeSearchEndpoint(paymentService),
		serializable.DecodeSearchRequest,
		serializable.EncodeSearchResponse,
	)

	http.Handle("/capture", captureHandler)
	http.Handle("/authorize", authorizeHandler)
	http.Handle("/cancel", cancelHandler)
	http.Handle("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
