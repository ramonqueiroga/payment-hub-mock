package serializable

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/transport"
)

//DecodeCancelRequest creates the cancel decode request
func DecodeCancelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.CancelRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeCancelResponse creates the cancel encode response
func EncodeCancelResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
