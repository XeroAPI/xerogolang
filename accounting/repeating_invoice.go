package accounting

type RepeatingInvoice struct {

	// See Invoice Types
	Type string `json:"Type,omitempty"`

	// See LineItems
	LineItems []LineItem `json:"LineItems,omitempty"`

	// Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty"`

	// ACCREC only – additional reference number
	Reference string `json:"Reference,omitempty"`

	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty"`

	// The currency that invoice has been raised in (see Currencies)
	CurrencyCode string `json:"CurrencyCode,omitempty"`

	// One of the following : DRAFT or AUTHORISED – See Invoice Status Codes
	Status string `json:"Status,omitempty"`

	// Total of invoice excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty"`

	// Total tax on invoice
	TotalTax float32 `json:"TotalTax,omitempty"`

	// Total of Invoice tax inclusive (i.e. SubTotal + TotalTax)
	Total float32 `json:"Total,omitempty"`

	// Xero generated unique identifier for repeating invoice template
	RepeatingInvoiceID string `json:"RepeatingInvoiceID,omitempty"`

	// boolean to indicate if an invoice has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`
}
