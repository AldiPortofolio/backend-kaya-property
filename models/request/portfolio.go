package request

type FilterPortfolio struct {
	Status     string `json:"status"`
	CustomerID int    `json:"customer_id"`
}
