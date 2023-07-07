package dto

type (
	AddressOutDTO struct {
		Street      string `json:"street"`
		City        string `json:"city"`
		State       string `json:"state"`
		ZipCode     string `json:"zipCode"`
		CountryCode string `json:"countryCode"`
	}

	UserOutDTO struct {
		FirstName string         `json:"firstName"`
		LastName  string         `json:"lastName"`
		Email     string         `json:"email"`
		Address   *AddressOutDTO `json:"address"`
	}
)
