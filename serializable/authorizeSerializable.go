package serializable

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/transport"
)

//DecodeAuthorizeRequest creates the authorize decode request
func DecodeAuthorizeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.AuthorizeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeAuthorizeResponse creates the authorize encode response
func EncodeAuthorizeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
