package accounting

//Account represents individual accounts in a Xero organisation
type Account struct {

	// Customer defined alpha numeric account code e.g 200 or SALES (max length = 10)
	Code string `json:"Code,omitempty" xml:"Code,omitempty"`

	// Name of account (max length = 150)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// See Account Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// For bank accounts only (Account Type BANK)
	BankAccountNumber string `json:"BankAccountNumber,omitempty" xml:"BankAccountNumber,omitempty"`

	// Accounts with a status of ACTIVE can be updated to ARCHIVED. See Account Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Description of the Account. Valid for all types of accounts except bank accounts (max length = 4000)
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// For bank accounts only. See Bank Account types
	BankAccountType string `json:"BankAccountType,omitempty" xml:"BankAccountType,omitempty"`

	// For bank accounts only
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// See Tax Types
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`

	// Boolean – describes whether account can have payments applied to it
	EnablePaymentsToAccount bool `json:"EnablePaymentsToAccount,omitempty" xml:"EnablePaymentsToAccount,omitempty"`

	// Boolean – describes whether account code is available for use with expense claims
	ShowInExpenseClaims bool `json:"ShowInExpenseClaims,omitempty" xml:"ShowInExpenseClaims,omitempty"`

	// The Xero identifier for an account – specified as a string following the endpoint name e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9
	AccountID string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`

	// See Account Class Types
	Class string `json:"Class,omitempty" xml:"Class,omitempty"`

	// If this is a system account then this element is returned. See System Account types. Note that non-system accounts may have this element set as either “” or null.
	SystemAccount string `json:"SystemAccount,omitempty" xml:"SystemAccount,omitempty"`

	// Shown if set
	ReportingCode string `json:"ReportingCode,omitempty" xml:"ReportingCode,omitempty"`

	// Shown if set
	ReportingCodeName string `json:"ReportingCodeName,omitempty" xml:"RepostingCodeName,omitempty"`

	// boolean to indicate if an account has an attachment (read only)
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`
}

//AccountResponse contains a collection of Accounts
type AccountResponse struct {
	Accounts []Account `json:"Accounts,omitempty" xml:"Accounts,omitempty"`
}
