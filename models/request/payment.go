package request

type Payment struct {
	Lot            int    `json:"lot"`
	PropertyID     int    `json:"property_id"`
	SecondaryID    int    `json:"secondary_id"`
	CustomerID     int    `json:"customer_id"`
	PaymentMethode string `json:"payment_methode"`
}

type Withdrawal struct {
	Amount     float64 `json:"amount"`
	VerifyCode string  `json:"verify_code"`
}

type Topup struct {
	Amount float64 `json:"amount"`
}
