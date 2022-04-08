package request

type ActivitasFilter struct {
	Filter string `json:"filter"`
	CustomerID int `json:"customer_id"`
	GeneralRequest
}
