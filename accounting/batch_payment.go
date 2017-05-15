package accounting

type BatchPayment struct {

	// A user defined bank account number.
	BankAccountNumber string `json:"BankAccountNumber,omitempty" xml:"BankAccountNumber,omitempty"`

	// Full name of bank account
	BankAccountName string `json:"BankAccountName,omitempty" xml:"BankAccountName,omitempty"`

	// Details of the Batch payment
	Details string `json:"Details,omitempty" xml:"Details,omitempty"`

	// Code of the Batch payment
	Code string `json:"Code,omitempty" xml:"Code,omitempty"`

	// Reference of the Batch payment
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`
}
