package request

type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	CustomerID uint64
}

type ResetPassword struct {
	Token      string `json:"token"`
	CustomerID int    `json:"customer_id"`
	Password   string `json:"password"`
}
