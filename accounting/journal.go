package accounting

type Journal struct {

	// Xero identifier
	JournalID string `json:"JournalID,omitempty"`

	// Date the journal was posted
	JournalDate string `json:"JournalDate,omitempty"`

	// Xero generated journal number
	JournalNumber string `json:"JournalNumber,omitempty"`

	// Created date UTC format
	CreatedDateUTC string `json:"CreatedDateUTC,omitempty"`

	//
	Reference string `json:"Reference,omitempty"`

	// The identifier for the source transaction (e.g. InvoiceID)
	SourceID string `json:"SourceID,omitempty"`

	// The journal source type. The type of transaction that created the journal
	SourceType string `json:"SourceType,omitempty"`

	// See JournalLines
	JournalLines []JournalLine `json:"JournalLines,omitempty"`
}
