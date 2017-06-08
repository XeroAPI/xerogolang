package accounting

//TaxComponent is a component of tax witjin a TaxRate
type TaxComponent struct {

	// Name of Tax Component
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// Tax Rate (up to 4dp)
	Rate float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`

	// Boolean to describe if Tax rate is compounded.Learn more
	IsCompound bool `json:"IsCompound,omitempty" xml:"IsCompound,omitempty"`

	// Filter by a Tax Type
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`
}
