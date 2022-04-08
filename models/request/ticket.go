package request

type FilterTicket struct {
	GeneralRequest
	CustomerID int `json:"customer_id"`
}
