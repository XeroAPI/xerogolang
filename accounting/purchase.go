package accounting

type Purchase struct {

	// Unit Price of the item. By default UnitPrice is rounded to two decimal places. You can use 4 decimal places by adding the unitdp=4 querystring parameter to your request.
	UnitPrice float32 `json:"UnitPrice,omitempty"`

	// Default account code to be used for purchased/sale. Not applicable to the purchase details of tracked items
	AccountCode string `json:"AccountCode,omitempty"`

	// Cost of goods sold account. Only applicable to the purchase details of tracked items.
	COGSAccountCode string `json:"COGSAccountCode,omitempty"`

	// Used as an override if the default Tax Code for the selected <AccountCode> is not correct – see TaxTypes.
	TaxType string `json:"TaxType,omitempty"`
}
