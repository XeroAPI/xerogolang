package accounting

//Overpayment is used when a debtor overpays an invoice
type Overpayment struct {

	// See Overpayment Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// The date the overpayment is created YYYY-MM-DD
	Date string `json:"Date,omitempty" xml:"Date,omitempty"`

	// See Overpayment Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Overpayment Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// See Overpayment Line Items
	LineItems []LineItem `json:"LineItems,omitempty" xml:"LineItems,omitempty"`

	// The subtotal of the overpayment excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// The total tax on the overpayment
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// The total of the overpayment (subtotal + total tax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// UTC timestamp of last update to the overpayment
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`

	// Currency used for the overpayment
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// Xero generated unique identifier
	OverpaymentID string `json:"OverpaymentID,omitempty" xml:"OverpaymentID,omitempty"`

	// The currency rate for a multicurrency overpayment. If no rate is specified, the XE.com day rate is used
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The remaining credit balance on the overpayment
	RemainingCredit string `json:"RemainingCredit,omitempty" xml:"RemainingCredit,omitempty"`

	// See Allocations
	Allocations []Allocation `json:"Allocations,omitempty" xml:"Allocations,omitempty"`

	// See Payments
	Payments []Payment `json:"Payments,omitempty" xml:"Payments,omitempty"`

	// boolean to indicate if a overpayment has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`
}
