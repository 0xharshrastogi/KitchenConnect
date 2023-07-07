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

	AddressInDTO struct {
		Street      string `json:"street" validate:"required"`
		City        string `json:"city" validate:"required"`
		State       string `json:"state" validate:"required"`
		ZipCode     string `json:"zipCode" validate:"required"`
		CountryCode string `json:"countryCode" validate:"required"`
	}

	UserInDTO struct {
		FirstName string        `json:"firstName" validate:"required"`
		LastName  string        `json:"lastName" validate:"required"`
		Email     string        `json:"email" validate:"required"`
		Password  string        `json:"password" validate:"required"`
		Address   *AddressInDTO `json:"address" validate:"required"`
	}
)
