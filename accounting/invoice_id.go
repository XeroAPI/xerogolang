package accounting

//InvoiceID should only be used when you only need to return an Invoice ID and/or InvoiceNumber
type InvoiceID struct {
	InvoiceID     string `json:"InvoiceID,omitempty" xml:"InvoiceID,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty" xml:"InvoiceNumber,omitempty"`
}
