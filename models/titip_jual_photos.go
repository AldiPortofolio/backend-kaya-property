package models

type TitipJualPhotos struct {
	ID          int    `gorm:"column:id" json:"id"`
	TitipJualID int    `gorm:"column:titip_jual_id" json:"titip_jual_id"`
	Photo       string `gorm:"column:photo" json:"photo"`
}

func (t *TitipJualPhotos) TableName() string {
	return "titip_jual_photos"
}
