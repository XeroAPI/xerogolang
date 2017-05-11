package accounting

type TaxComponent struct {

	// Name of Tax Component
	Name string `json:"Name,omitempty"`

	// Tax Rate (up to 4dp)
	Rate float32 `json:"Rate,omitempty"`

	// Boolean to describe if Tax rate is compounded.Learn more
	IsCompound float32 `json:"IsCompound,omitempty"`

	// Filter by a Tax Type
	TaxType string `json:"TaxType,omitempty"`
}
