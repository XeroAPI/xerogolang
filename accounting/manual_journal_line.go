package accounting

//ManualJournalLine is a line on a Manual Journal
type ManualJournalLine struct {
	// See Accounts
	AccountCode string `json:"AccountCode" xml:"AccountCode"`

	// The description from the source transaction line item. Only returned if populated.
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// Net amount of journal line. This will be a positive value for a debit and negative for a credit
	LineAmount float32 `json:"LineAmount" xml:"LineAmount"`

	// The calculated tax amount based on the TaxType and LineAmount
	TaxAmount float32 `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`

	// Used as an override if the default Tax Code for the selected <AccountCode> is not correct – see TaxTypes.
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`

	// Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 <TrackingCategory> elements.
	Tracking []TrackingCategory `json:"Tracking,omitempty" xml:"Tracking>TrackingCategory,omitempty"`
}
