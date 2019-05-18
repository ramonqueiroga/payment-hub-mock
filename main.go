package main

import (
	"net/http"
	"os"
	"payment-hub-mock/business"
	"payment-hub-mock/external"
	"payment-hub-mock/logging"
	"payment-hub-mock/transport"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	router := mux.NewRouter()
	router = router.PathPrefix("/v1/payment/creditcard").Subrouter()
	logger := log.NewLogfmtLogger(os.Stderr)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var paymentService business.PaymentService
	paymentService = business.PaymentServiceImpl{}
	paymentService = logging.LoggerMiddleware{Logger: logger, Next: paymentService}

	var authorize endpoint.Endpoint
	authorize = external.MakeAuthorizationEndpoint(paymentService)
	authorize = logging.MiddlewareEndpoint(log.With(logger, "method", "authorize"))(authorize)
	authorizeHandler := httptransport.NewServer(
		authorize,
		transport.DecodeAuthorizeRequest,
		transport.EncodeAuthorizeResponse,
	)

	var captureEndpoint endpoint.Endpoint
	captureEndpoint = external.MakeCaptureEndpoint(paymentService)
	captureEndpoint = logging.MiddlewareEndpoint(log.With(logger, "method", "capture"))(captureEndpoint)
	captureHandler := httptransport.NewServer(
		captureEndpoint,
		transport.DecodeCaptureRequest,
		transport.EncodeCaptureResponse,
	)

	var cancelEndpoint endpoint.Endpoint
	cancelEndpoint = external.MakeCancelEndpoint(paymentService)
	cancelEndpoint = logging.MiddlewareEndpoint(log.With(logger, "method", "cancel"))(cancelEndpoint)
	cancelHandler := httptransport.NewServer(
		cancelEndpoint,
		transport.DecodeCancelRequest,
		transport.EncodeCancelResponse,
	)

	var searchEndpoint endpoint.Endpoint
	searchEndpoint = external.MakeSearchEndpoint(paymentService)
	searchEndpoint = logging.MiddlewareEndpoint(log.With(logger, "method", "search"))(searchEndpoint)
	searchHandler := httptransport.NewServer(
		searchEndpoint,
		transport.DecodeSearchRequest,
		transport.EncodeSearchResponse,
	)

	//routes
	router.Methods("POST").Handler(authorizeHandler)
	router.Methods("PUT").Path("/{payment_id}/capture").Handler(captureHandler)
	router.Methods("DELETE").Path("/{payment_id}").Handler(cancelHandler)
	router.Methods("GET").Path("/{payment_id}").Handler(searchHandler)

	//up and running
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", router))
}
