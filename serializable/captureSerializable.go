package serializable

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/transport"
)

//DecodeCaptureRequest create the decode for the capture request
func DecodeCaptureRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.CaptureRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeCaptureResponse create the encode for the capture response
func EncodeCaptureResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
