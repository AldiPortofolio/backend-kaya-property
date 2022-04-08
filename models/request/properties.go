package request

type FilterProperty struct {
	GeneralRequest
	Name       string `json:"name"`
	CityID     int    `json:"city_id"`
	CustomerID int    `json:"customer_id"`
	Status     string `json:"status"`
}
