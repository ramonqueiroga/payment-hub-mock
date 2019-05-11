package business

// Payments ....
type Payments struct {
	Payments []struct {
		Amount     float64 `json:"amount"`
		Capture    bool    `json:"capture"`
		CreditCard struct {
			Brand           string `json:"brand"`
			CardNumber      string `json:"card_number"`
			ExpirationMonth int    `json:"expiration_month"`
			ExpirationYear  int    `json:"expiration_year"`
			Holder          string `json:"holder"`
			SecurityCode    string `json:"security_code"`
			SoftDescriptor  string `json:"soft_descriptor"`
		} `json:"credit_card"`
		Installments int    `json:"installments"`
		Reference    string `json:"reference"`
		Type         string `json:"type"`
	} `json:"payments"`
}
