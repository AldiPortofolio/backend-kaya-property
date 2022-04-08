package request

type CustomerUpload struct {
	UploadKtp    string `json:"upload_ktp"`
	UploadSelfie string `json:"upload_selfie"`
}

type CustomerVerification struct {
	VerifyCode string `json:"verify_code"`
}

type CustomerAccount struct {
	BankID        int    `json:"bank_id"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
}

type CustomerPassword struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
