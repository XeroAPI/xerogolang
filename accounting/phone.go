package accounting

type Phone struct {

	// max length = 50
	PhoneNumber string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`

	// max length = 10
	PhoneAreaCode string `json:"PhoneAreaCode,omitempty" xml:"PhoneAreaCode,omitempty"`

	// max length = 20
	PhoneCountryCode string `json:"PhoneCountryCode,omitempty" xml:"PhoneCountryCode,omitempty"`
}
