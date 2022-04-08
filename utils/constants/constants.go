package constants

const (
	CREATED_CODE    = 201
	SUCCESS_CODE    = 200
	NO_CONTENT_CODE = 204

	BAD_REQUEST_CODE  = 400
	UNAUTHORIZED_CODE = 401
	NOTFOUND_CODE     = 404

	INTERNAL_SERVER_ERROR_CODE = 500
)

const (
	HttpMethodGet    = "GET"
	HttpMethodPost   = "POST"
	HttpMethodPut    = "PUT"
	HttpMethodDelete = "DELETE"
)

const (
	NONACTIVE = "NonActive"
	ACTIVE    = "Active"

	TOPUP              = "TOPUP"
	WITHDRAWAL         = "WITHDRAWAL"
	PENJUALAN_PROPERTY = "PENJUALAN_PROPERTY"
	PEMBELIAN_LOT      = "PEMBELIAN_LOT"
	PENJUALAN_LOT      = "PENJUALAN_LOT"

	STATUS_PENDING  = "PENDING"
	STATUS_DONE     = "DONE"
	STATUS_CANCELED = "CANCELED"

	PEMBELIAN = "BELI"
	PENJUALAN = "JUAL"
)

const (
	SECONDARY_STATUS_OPEN    = "OPEN"
	SECONDARY_STATUS_CLOSE   = "CLOSE"
	SECONDARY_STATUS_BATAL   = "BATAL"
	SECONDARY_STATUS_PENDING = "PENDING"
)

const (
	PENDING = 1000
	SUCCESS = 1001
	FAILED  = 1002
	EXPIRED = 1003
)

const (
	SALT = "kaya345678901223"
)
