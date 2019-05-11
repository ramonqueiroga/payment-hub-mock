package transport

import "payment-hub-mock/business"

//CancelRequest is the model to consume the cancelation service
type CancelRequest struct {
	PaymentID string `json:"paymentID"`
}

//CancelResponse is the model that responses the cancelation service
type CancelResponse struct {
	TransactionModel business.TransactionModel `json:"transactionModel"`
	Error            string                    `json:"error"`
}
