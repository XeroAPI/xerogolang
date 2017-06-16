package accounting

import (
	"encoding/xml"

	"github.com/XeroAPI/xerogolang"
	"github.com/markbates/goth"
)

//TrackingOption is an option from within a Tracking category
type TrackingOption struct {

	// The Xero identifier for a tracking optione.g. ae777a87-5ef3-4fa0-a4f0-d10e1f13073a
	TrackingOptionID string `json:"TrackingOptionID,omitempty" xml:"TrackingOptionID,omitempty"`

	// The name of the tracking option e.g. Marketing, East (max length = 50)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// The status of a tracking option
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Filter by a tracking categorye.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	TrackingCategoryID string `json:"TrackingCategoryID,omitempty" xml:"-"`
}

//Options is a collection of TrackingOptions
type Options struct {
	Options []TrackingOption `json:"Options,omitempty" xml:"Option,omitempty"`
}

//Add will add tracking options to the TrackingCategory Specified on the first option
//All options should belong to the same Tracking Category
func (o *Options) Add(provider *xerogolang.Provider, session goth.Session) (*TrackingCategories, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(o, "  ", "	")
	if err != nil {
		return nil, err
	}

	trackingCategoryResponseBytes, err := provider.Create(session, "TrackingCategories/"+o.Options[0].TrackingCategoryID+"/Options", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalTrackingCategory(trackingCategoryResponseBytes)
}

//Update will update a given tracking option
func (t *TrackingOption) Update(provider *xerogolang.Provider, session goth.Session) (*TrackingCategories, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(t, "  ", "	")
	if err != nil {
		return nil, err
	}

	trackingCategoryResponseBytes, err := provider.Update(session, "TrackingCategories/"+t.TrackingCategoryID+"/Options/"+t.TrackingOptionID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalTrackingCategory(trackingCategoryResponseBytes)
}
