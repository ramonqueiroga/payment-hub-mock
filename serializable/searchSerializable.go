package serializable

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/transport"
)

//DecodeSearchRequest creates the cancel decode request
func DecodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeSearchResponse creates the cancel encode response
func EncodeSearchResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
