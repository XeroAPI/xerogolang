package accounting

type LinkedTransaction struct {

	// Filter by the SourceTransactionID. Get all the linked transactions created from a particular ACCPAY invoice
	SourceTransactionID string `json:"SourceTransactionID,omitempty"`

	// The line item identifier from the source transaction.
	SourceLineItemID string `json:"SourceLineItemID"`

	// Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID=4bb34b03-3378-4bb2-a0ed-6345abf3224e&Status=APPROVED.
	ContactID string `json:"ContactID,omitempty"`

	// Filter by the TargetTransactionID. Get all the linked transactions allocated to a particular ACCREC invoice
	TargetTransactionID string `json:"TargetTransactionID,omitempty"`

	// The line item identifier from the target transaction. It is possible to link multiple billable expenses to the same TargetLineItemID.
	TargetLineItemID string `json:"TargetLineItemID,omitempty"`

	// The Xero identifier for an Linked Transaction e.g. /LinkedTransactions/297c2dc5-cc47-4afd-8ec8-74990b8761e9
	LinkedTransactionID string `json:"LinkedTransactionID,omitempty"`

	// Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID=4bb34b03-3378-4bb2-a0ed-6345abf3224e&Status=APPROVED.
	Status string `json:"Status,omitempty"`

	// This will always be BILLABLEEXPENSE. More types may be added in future.
	Type_ string `json:"Type,omitempty"`

	// The last modified date in UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`

	// The Type of the source tranasction. This will be ACCPAY if the linked transaction was created from an invoice and SPEND if it was created from a bank transaction.
	SourceTransactionTypeCode string `json:"SourceTransactionTypeCode,omitempty"`
}
