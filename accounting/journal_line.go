package accounting

type JournalLine struct {

	// total for line. Debits are positive, credits are negative value
	LineAmount string `json:"LineAmount"`

	// See Accounts
	AccountCode string `json:"AccountCode"`

	// Description for journal line
	Description string `json:"Description,omitempty"`

	// Used as an override if the default Tax Code for the selected <AccountCode> is not correct – see TaxTypes.
	TaxType string `json:"TaxType,omitempty"`

	// Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 <TrackingCategory> elements.
	Tracking []TrackingCategory `json:"Tracking,omitempty"`

	// The calculated tax amount based on the TaxType and LineAmount
	TaxAmount float32 `json:"TaxAmount,omitempty"`
}
