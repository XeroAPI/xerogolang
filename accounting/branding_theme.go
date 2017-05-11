package accounting

//BrandingTheme applies structure and visuals to an invoice when printed or sent
type BrandingTheme struct {

	// Xero identifier
	BrandingThemeID string `json:"BrandingThemeID,omitempty"`

	// Name of branding theme
	Name string `json:"Name,omitempty"`

	// Integer – ranked order of branding theme. The default branding theme has a value of 0
	SortOrder float64 `json:"SortOrder,omitempty"`

	// UTC timestamp of creation date of branding theme
	CreatedDateUTC string `json:"CreatedDateUTC,omitempty"`
}
