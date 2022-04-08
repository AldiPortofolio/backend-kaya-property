package request

type GeneralRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
type ByID struct {
	ID int `json:"id"`
}
type ByEmail struct {
	Email string `json:"email"`
}
