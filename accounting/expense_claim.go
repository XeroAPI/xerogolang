package accounting

type ExpenseClaim struct {

	// See Receipts
	Receipts []Receipt `json:"Receipts"`

	// Xero generated unique identifier for an expense claim
	ExpenseClaimID string `json:"ExpenseClaimID,omitempty"`

	// See Payments
	Payments []Payment `json:"Payments,omitempty"`

	// Current status of an expense claim – see status types
	Status string `json:"Status,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`

	// The total of an expense claim being paid
	Total float32 `json:"Total,omitempty"`

	// The amount due to be paid for an expense claim
	AmountDue float32 `json:"AmountDue,omitempty"`

	// The amount still to pay for an expense claim
	AmountPaid float32 `json:"AmountPaid,omitempty"`

	// The date when the expense claim is due to be paid YYYY-MM-DD
	PaymentDueDate string `json:"PaymentDueDate,omitempty"`

	// The date the expense claim will be reported in Xero YYYY-MM-DD
	ReportingDate string `json:"ReportingDate,omitempty"`

	// The Xero identifier for the Receipt e.g. e59a2c7f-1306-4078-a0f3-73537afcbba9
	ReceiptID string `json:"ReceiptID"`
}
