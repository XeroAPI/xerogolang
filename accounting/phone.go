package accounting

type Phone struct {

	// max length = 50
	PhoneNumber string `json:"PhoneNumber,omitempty"`

	// max length = 10
	PhoneAreaCode string `json:"PhoneAreaCode,omitempty"`

	// max length = 20
	PhoneCountryCode string `json:"PhoneCountryCode,omitempty"`
}
