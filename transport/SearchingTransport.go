package transport

import "payment-hub-mock/business"

//SearchRequest is the model to consumem the search transcation service
type SearchRequest struct {
	PaymentID string `json:"paymentID"`
}

//SearchResponse is the model that responses the searching transaction service
type SearchResponse struct {
	TransactionModel business.TransactionModel `json:"transactionModel"`
	Error            string                    `json:"error"`
}
