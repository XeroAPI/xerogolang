package accounting

type Address struct {
	AddressType string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`

	// max length = 500
	AddressLine1 string `json:"AddressLine1,omitempty" xml:"AddressLine1,omitempty"`

	// max length = 500
	AddressLine2 string `json:"AddressLine2,omitempty" xml:"AddressLine2,omitempty"`

	// max length = 500
	AddressLine3 string `json:"AddressLine3,omitempty" xml:"AddressLine3,omitempty"`

	// max length = 500
	AddressLine4 string `json:"AddressLine4,omitempty" xml:"AddressLine4,omitempty"`

	// max length = 255
	City string `json:"City,omitempty" xml:"City,omitempty"`

	// max length = 255
	Region string `json:"Region,omitempty" xml:"Region,omitempty"`

	// max length = 50
	PostalCode string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`

	// max length = 50, [A-Z], [a-z] only
	Country string `json:"Country,omitempty" xml:"Country,omitempty"`

	// max length = 255
	AttentionTo string `json:"AttentionTo,omitempty" xml:"AttentionTo,omitempty"`
}
