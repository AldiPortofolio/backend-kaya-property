package request

type FilterPropertySecondary struct {
	GeneralRequest
	Status     string `json:"status"`
	CustomerID int    `json:"customer_id"`
}
