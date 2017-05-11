package accounting

type Currency struct {

	// 3 letter alpha code for the currency – see list of currency codes
	Code string `json:"Code,omitempty"`

	// Name of Currency
	Description string `json:"Description,omitempty"`
}
