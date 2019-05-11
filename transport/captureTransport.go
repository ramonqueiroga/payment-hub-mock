package transport

import "payment-hub-mock/business"

//CaptureRequest is the model for consume the capture service
type CaptureRequest struct {
	Payments business.Payments `json:"payments"`
}

//CaptureResponse is the model that returns in the capture service
type CaptureResponse struct {
	TransactionModel business.TransactionModel `json:"transactionModel"`
	Error            string                    `json:"error"`
}
