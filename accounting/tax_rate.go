package accounting

type TaxRate struct {

	// Name of tax rate
	Name string `json:"Name,omitempty"`

	// See Tax Types – can only be used on update calls
	TaxType string `json:"TaxType,omitempty"`

	// See TaxComponents
	TaxComponents []TaxComponent `json:"TaxComponents,omitempty"`

	// See Status Codes
	Status string `json:"Status,omitempty"`

	// See ReportTaxTypes
	ReportTaxType string `json:"ReportTaxType"`

	// Boolean to describe if tax rate can be used for asset accounts i.e. true,false
	CanApplyToAssets float32 `json:"CanApplyToAssets,omitempty"`

	// Boolean to describe if tax rate can be used for equity accounts i.e. true,false
	CanApplyToEquity float32 `json:"CanApplyToEquity,omitempty"`

	// Boolean to describe if tax rate can be used for expense accounts i.e. true,false
	CanApplyToExpenses float32 `json:"CanApplyToExpenses,omitempty"`

	// Boolean to describe if tax rate can be used for liability accounts i.e. true,false
	CanApplyToLiabilities float32 `json:"CanApplyToLiabilities,omitempty"`

	// Boolean to describe if tax rate can be used for revenue accounts i.e. true,false
	CanApplyToRevenue float32 `json:"CanApplyToRevenue,omitempty"`

	// Tax Rate (decimal to 4dp) e.g 12.5000
	DisplayTaxRate float32 `json:"DisplayTaxRate,omitempty"`

	// Effective Tax Rate (decimal to 4dp) e.g 12.5000
	EffectiveRate float32 `json:"EffectiveRate,omitempty"`
}
