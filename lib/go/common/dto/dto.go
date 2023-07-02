package dto

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Address   UserAddress
}

type UserAddress struct {
	City    string
	State   string
	Country string
}
