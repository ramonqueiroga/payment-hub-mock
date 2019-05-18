package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/business"

	"github.com/gorilla/mux"
)

//SearchRequest is the model to consumem the search transcation service
type SearchRequest struct {
	PaymentID string `json:"paymentID"`
}

//SearchResponse is the model that responses the searching transaction service
type SearchResponse struct {
	Transaction business.Transaction `json:"transaction"`
	Error       string               `json:"error"`
}

//DecodeSearchRequest creates the cancel decode request
func DecodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["payment_id"]

	if !ok {
		return nil, errBadRouting
	}

	return SearchRequest{PaymentID: id}, nil
}

//EncodeSearchResponse creates the cancel encode response
func EncodeSearchResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
