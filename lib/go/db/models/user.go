package models

type User struct {
	ID        int    `gorm:"primaryKey"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string
	Password  string
	AddressId int      `gorm:"column:address_id"`
	Address   *Address `gorm:"foreignKey:AddressId"`
}

func (u *User) TableName() string {
	return "user"
}
