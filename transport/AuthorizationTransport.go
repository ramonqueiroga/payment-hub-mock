//AuthorizeRequest is the model to consume the authorization service for the generated payment transaction
package transport

import "payment-hub-mock/business"

type AuthorizeRequest struct {
	PaymentID string `json:"paymentID"`
}

//AuthorizeResponse is the model that responses the authorization service
type AuthorizeResponse struct {
	TransactionModel business.TransactionModel `json:"transactionModel"`
	Error            string                    `json:"error"`
}
