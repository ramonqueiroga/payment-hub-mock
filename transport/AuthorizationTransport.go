package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/business"
)

//AuthorizeRequest authorize request object to transport authorize request message
type AuthorizeRequest struct {
	PaymentID string `json:"paymentID"`
}

//AuthorizeResponse is the model that responses the authorization service
type AuthorizeResponse struct {
	Transaction business.Transaction `json:"transaction"`
	Error       string               `json:"error"`
}

//DecodeAuthorizeRequest creates the authorize decode request
func DecodeAuthorizeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AuthorizeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeAuthorizeResponse creates the authorize encode response
func EncodeAuthorizeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
