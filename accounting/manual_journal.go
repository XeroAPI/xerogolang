package accounting

type ManualJournal struct {

	// Description of journal being posted
	Narration string `json:"Narration"`

	// See JournalLines
	JournalLines []JournalLine `json:"JournalLines"`

	// Date journal was posted – YYYY-MM-DD
	Date string `json:"Date,omitempty"`

	// NoTax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty"`

	// See Manual Journal Status Codes
	Status string `json:"Status,omitempty"`

	// Url link to a source document – shown as “Go to [appName]” in the Xero app
	Url string `json:"Url,omitempty"`

	// Boolean – default is true if not specified
	ShowOnCashBasisReports bool `json:"ShowOnCashBasisReports,omitempty"`

	// Boolean to indicate if a manual journal has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`

	// The Xero identifier for a Manual Journal
	JournalID string `json:"JournalID,omitempty"`
}
