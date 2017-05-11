package accounting

type BankTransfer struct {

	//
	Amount string `json:"Amount"`

	// The date of the Transfer YYYY-MM-DD
	Date string `json:"Date,omitempty"`

	// The identifier of the Bank Transfer
	BankTransferID string `json:"BankTransferID,omitempty"`

	// The currency rate
	CurrencyRate float32 `json:"CurrencyRate,omitempty"`

	// The Bank Transaction ID for the source account
	FromBankTransactionID string `json:"FromBankTransactionID,omitempty"`

	// The Bank Transaction ID for the destination account
	ToBankTransactionID string `json:"ToBankTransactionID,omitempty"`

	// Boolean to indicate if a Bank Transfer has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`

	// UTC timestamp of creation date of bank transfer
	CreatedDateUTC string `json:"CreatedDateUTC,omitempty"`
}
