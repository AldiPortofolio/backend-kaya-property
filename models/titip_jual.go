package models

type TitipJual struct {
	ID                 int               `gorm:"column:id" json:"id"`
	FullName           string            `gorm:"column:full_name" json:"full_name"`
	Email              string            `gorm:"column:email" json:"email"`
	PhoneNumber        string            `gorm:"column:no_handphone" json:"no_handphone"`
	Address1           string            `gorm:"column:address1" json:"address1"`
	Address2           string            `gorm:"column:address2" json:"address2"`
	Price              float64           `gorm:"column:price" json:"price"`
	Reason             string            `gorm:"column:reason" json:"reason"`
	LuasTanah          float32           `gorm:"column:luas_tanah" json:"luas_tanah"`
	LuasBangunan       float32           `gorm:"column:luas_bangunan" json:"luas_bangunan"`
	JumlahTingkat      int               `gorm:"column:jumlah_tingkat" json:"jumlah_tingkat"`
	JumlahKamarTidur   int               `gorm:"column:jumlah_kamar_tidur" json:"jumlah_kamar_tidur"`
	JumlahKamarMandi   int               `gorm:"column:jumlah_kamar_mandi" json:"jumlah_kamar_mandi"`
	Legalitas          string            `gorm:"column:legalitas" json:"legalitas"`
	LegalitasCondition string            `gorm:"column:legalitas_condition" json:"legalitas_condition"`
	HadapRumah         string            `gorm:"column:hadap_rumah" json:"hadap_rumah"`
	OtherInformation   string            `gorm:"column:other_information" json:"other_information"`
	Status             string            `gorm:"column:status" json:"status"`
	CityID             int32             `gorm:"column:city_id" json:"city_id"`
	Photos             []TitipJualPhotos `gorm:"foreignKey:TitipJualID;references:id" json:"photos"`
}

func (t *TitipJual) TableName() string {
	return "titip_juals"
}
