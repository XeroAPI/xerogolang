package accounting

import (
	"encoding/json"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//BrandingTheme applies structure and visuals to an invoice when printed or sent
type BrandingTheme struct {

	// Xero identifier
	BrandingThemeID string `json:"BrandingThemeID,omitempty" xml:"BrandingThemeID,omitempty"`

	// Name of branding theme
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// Integer – ranked order of branding theme. The default branding theme has a value of 0
	SortOrder float64 `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`

	// UTC timestamp of creation date of branding theme
	CreatedDateUTC string `json:"CreatedDateUTC,omitempty" xml:"CreatedDateUTC,omitempty"`
}

//BrandingThemes contains a collection of BrandingThemes
type BrandingThemes struct {
	BrandingThemes []BrandingTheme `json:"BrandingThemes" xml:"BrandingTheme"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (b *BrandingThemes) convertDates() error {
	var err error
	for n := len(b.BrandingThemes) - 1; n >= 0; n-- {
		b.BrandingThemes[n].CreatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(b.BrandingThemes[n].CreatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalBrandingTheme(brandingThemeResponseBytes []byte) (*BrandingThemes, error) {
	var brandingThemeResponse *BrandingThemes
	err := json.Unmarshal(brandingThemeResponseBytes, &brandingThemeResponse)
	if err != nil {
		return nil, err
	}

	err = brandingThemeResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return brandingThemeResponse, err
}

//FindBrandingThemes will get all BrandingThemes.
func FindBrandingThemes(provider *xerogolang.Provider, session goth.Session) (*BrandingThemes, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	brandingThemeResponseBytes, err := provider.Find(session, "BrandingThemes", additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalBrandingTheme(brandingThemeResponseBytes)
}
