package main

import (
	"log"
	"net/http"
	"payment-hub-mock/business"
	"payment-hub-mock/external"
	"payment-hub-mock/transport"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	paymentService := business.PaymentServiceImpl{}
	router := mux.NewRouter()
	router = router.PathPrefix("/v1/payment/creditcard").Subrouter()

	authorizeHandler := httptransport.NewServer(
		external.MakeAuthorizationEndpoint(paymentService),
		transport.DecodeAuthorizeRequest,
		transport.EncodeAuthorizeResponse,
	)

	captureHandler := httptransport.NewServer(
		external.MakeCaptureEndpoint(paymentService),
		transport.DecodeCaptureRequest,
		transport.EncodeCaptureResponse,
	)

	cancelHandler := httptransport.NewServer(
		external.MakeCancelEndpoint(paymentService),
		transport.DecodeCancelRequest,
		transport.EncodeCancelResponse,
	)

	searchHandler := httptransport.NewServer(
		external.MakeSearchEndpoint(paymentService),
		transport.DecodeSearchRequest,
		transport.EncodeSearchResponse,
	)

	router.Methods("POST").Handler(authorizeHandler)
	router.Methods("PUT").Path("/{payment_id}/capture").Handler(captureHandler)
	router.Methods("DELETE").Path("/{payment_id}").Handler(cancelHandler)
	router.Methods("GET").Path("/{payment_id}").Handler(searchHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
