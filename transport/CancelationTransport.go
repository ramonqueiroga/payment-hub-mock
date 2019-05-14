package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/business"

	"github.com/gorilla/mux"
)

//CancelRequest is the model to consume the cancelation service
type CancelRequest struct {
	PaymentID string `json:"paymentID"`
}

//CancelResponse is the model that responses the cancelation service
type CancelResponse struct {
	TransactionModel business.TransactionModel `json:"transactionModel"`
	Error            string                    `json:"error"`
}

//DecodeCancelRequest creates the cancel decode request
func DecodeCancelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["payment_id"]

	if !ok {
		return nil, errBadRouting
	}

	return CancelRequest{PaymentID: id}, nil
}

//EncodeCancelResponse creates the cancel encode response
func EncodeCancelResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
