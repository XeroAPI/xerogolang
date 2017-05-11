package accounting

//Allocation allocated an overpayment or Prepayment to an Invoice
type Allocation struct {

	// the amount being applied to the invoice
	AppliedAmount float32 `json:"AppliedAmount,omitempty" xml:"AppliedAmount,omitempty"`

	// the date the prepayment is applied YYYY-MM-DD (read-only). This will be the latter of the invoice date and the prepayment date.
	Date string `json:"Date,omitempty" xml:"Date,omitempty"`
}
