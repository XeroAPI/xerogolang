package accounting

//Prepayment are payments made before the associated document has been created
type Prepayment struct {

	// See Prepayment Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// The date the prepayment is created YYYY-MM-DD
	Date string `json:"Date,omitempty" xml:"Date,omitempty"`

	// See Prepayment Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Prepayment Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// See Prepayment Line Items
	LineItems []LineItem `json:"LineItems,omitempty" xml:"LineItems,omitempty"`

	// The subtotal of the prepayment excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// The total tax on the prepayment
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// The total of the prepayment(subtotal + total tax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// UTC timestamp of last update to the prepayment
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`

	// Currency used for the prepayment
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// Xero generated unique identifier
	PrepaymentID string `json:"PrepaymentID,omitempty" xml:"PrepaymentID,omitempty"`

	// The currency rate for a multicurrency prepayment. If no rate is specified, the XE.com day rate is used
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The remaining credit balance on the prepayment
	RemainingCredit string `json:"RemainingCredit,omitempty" xml:"RemainingCredit,omitempty"`

	// See Allocations
	Allocations []Allocation `json:"Allocations,omitempty" xml:"Allocations,omitempty"`

	// boolean to indicate if a prepayment has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`
}
