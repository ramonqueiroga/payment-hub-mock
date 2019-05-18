package business

// Transaction ....
type Transaction struct {
	PaymentID      string  `json:"payment_id"`
	Amount         float64 `json:"amount"`
	Backend        string  `json:"backend"`
	Status         string  `json:"status"`
	AdditionalInfo struct {
	} `json:"additional_info"`
	AuthorizationCode int    `json:"authorization_code"`
	ReturnCode        int    `json:"return_code"`
	ReturnMessage     string `json:"return_message"`
	Tid               int64  `json:"tid"`
	Nsu               int    `json:"nsu"`
}
