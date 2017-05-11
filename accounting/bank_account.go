package accounting

type BankAccount struct {

	// The Account Code of the Bank Account
	Code string `json:"Code,omitempty"`

	// The ID of the Bank Account
	AccountID string `json:"AccountID,omitempty"`

	// The Name Bank Account
	Name string `json:"Name,omitempty"`
}
