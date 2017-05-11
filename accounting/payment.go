package accounting

//Payment details payments against invoices and CreditNotes
type Payment struct {

	// Number of invoice or credit note you are applying payment to e.g. INV-4003
	InvoiceNumber string `json:"InvoiceNumber,omitempty" xml:"InvoiceNumber,omitempty"`

	// Number of invoice or credit note you are applying payment to e.g. INV-4003
	CreditNoteNumber string `json:"CreditNoteNumber,omitempty" xml:"CreditNoteNumber,omitempty"`

	// Code of account you are using to make the payment e.g. 001 (note: not all accounts have a code value)
	Code string `json:"Code,omitempty" xml:"Code,omitempty"`

	// Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06
	Date string `json:"Date,omitempty" xml:"Date,omitempty"`

	// Exchange rate when payment is received. Only used for non base currency invoices and credit notes e.g. 0.7500
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The amount of the payment. Must be less than or equal to the outstanding amount owing on the invoice e.g. 200.00
	Amount float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`

	// An optional description for the payment e.g. Direct Debit
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// An optional parameter for the payment. A boolean indicating whether you would like the payment to be created as reconciled when using PUT, or whether a payment has been reconciled when using GET
	IsReconciled string `json:"IsReconciled,omitempty" xml:"IsReconciled,omitempty"`

	// The status of the payment.
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Payment Types.
	PaymentType string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`

	// UTC timestamp of last update to the payment
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`

	// The Xero identifier for an Payment e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	PaymentID string `json:"PaymentID,omitempty" xml:"PaymentID,omitempty"`
}
