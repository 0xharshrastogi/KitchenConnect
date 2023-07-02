package models

type Address struct {
	ID          int `gorm:"primaryKey"`
	Street      string
	City        string
	State       string
	ZipCode     string `gorm:"column:zip_code"`
	CountryCode string `gorm:"column:country_code"`
}

func (a *Address) TableName() string {
	return "address"
}
