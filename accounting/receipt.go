package accounting

type Receipt struct {

	// Date of receipt – YYYY-MM-DD
	Date string `json:"Date"`

	// See LineItems
	Lineitems []LineItem `json:"Lineitems"`

	// Additional reference number
	Reference string `json:"Reference,omitempty"`

	// See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty"`

	// Total of receipt excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty"`

	// Total tax on receipt
	TotalTax float32 `json:"TotalTax,omitempty"`

	// Total of receipt tax inclusive (i.e. SubTotal + TotalTax)
	Total float32 `json:"Total,omitempty"`

	// Xero generated unique identifier for receipt
	ReceiptID string `json:"ReceiptID,omitempty"`

	// Current status of receipt – see status types
	Status string `json:"Status,omitempty"`

	// Xero generated sequence number for receipt in current claim for a given user
	ReceiptNumber string `json:"ReceiptNumber,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`

	// boolean to indicate if a receipt has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`

	// URL link to a source document – shown as “Go to [appName]” in the Xero app
	Url string `json:"Url,omitempty"`
}
