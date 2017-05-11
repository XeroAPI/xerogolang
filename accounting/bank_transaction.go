package accounting

//BankTransaction is a bank transaction
type BankTransaction struct {

	// See Bank Transaction Types
	Type string `json:"Type"`

	// See LineItems
	Lineitems []LineItem `json:"Lineitems"`

	// Boolean to show if transaction is reconciled
	IsReconciled bool `json:"IsReconciled,omitempty"`

	// Date of transaction – YYYY-MM-DD
	Date string `json:"Date,omitempty"`

	// Reference for the transaction. Only supported for SPEND and RECEIVE transactions.
	Reference string `json:"Reference,omitempty"`

	// The currency that bank transaction has been raised in (see Currencies). Setting currency is only supported on overpayments.
	CurrencyCode string `json:"CurrencyCode,omitempty"`

	// Exchange rate to base currency when money is spent or received. e.g. 0.7500 Only used for bank transactions in non base currency. If this isn’t specified for non base currency accounts then either the user-defined rate (preference) or the XE.com day rate will be used. Setting currency is only supported on overpayments.
	CurrencyRate float32 `json:"CurrencyRate,omitempty"`

	// URL link to a source document – shown as “Go to App Name”
	URL string `json:"Url,omitempty"`

	// See Bank Transaction Status Codes
	Status string `json:"Status,omitempty"`

	// Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty"`

	// Total of bank transaction excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty"`

	// Total tax on bank transaction
	TotalTax float32 `json:"TotalTax,omitempty"`

	// Total of bank transaction tax inclusive
	Total float32 `json:"Total,omitempty"`

	// Xero generated unique identifier for bank transaction
	BankTransactionID string `json:"BankTransactionID,omitempty"`

	// Xero generated unique identifier for a Prepayment. This will be returned on BankTransactions with a Type of SPEND-PREPAYMENT or RECEIVE-PREPAYMENT
	PrepaymentID string `json:"PrepaymentID,omitempty"`

	// Xero generated unique identifier for an Overpayment. This will be returned on BankTransactions with a Type of SPEND-OVERPAYMENT or RECEIVE-OVERPAYMENT
	OverpaymentID string `json:"OverpaymentID,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`

	// Boolean to indicate if a bank transaction has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`
}
