package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-hub-mock/business"
)

//CaptureRequest is the model for consume the capture service
type CaptureRequest struct {
	Payments business.Payments `json:"payments"`
}

//CaptureResponse is the model that returns in the capture service
type CaptureResponse struct {
	TransactionModel business.TransactionModel `json:"transactionModel"`
	Error            string                    `json:"error"`
}

//DecodeCaptureRequest create the decode for the capture request
func DecodeCaptureRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CaptureRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeCaptureResponse create the encode for the capture response
func EncodeCaptureResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
