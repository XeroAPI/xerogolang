package accounting

import (
	"encoding/json"
	"encoding/xml"

	"github.com/XeroAPI/xerogolang"
	"github.com/markbates/goth"
)

//TrackingCategory is used to segment data within a Xero organisation
type TrackingCategory struct {

	// The Xero identifier for a tracking categorye.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	TrackingCategoryID string `json:"TrackingCategoryID,omitempty" xml:"TrackingCategoryID,omitempty"`

	// The name of the tracking category e.g. Department, Region (max length = 100)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// The status of a tracking category
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Tracking Options
	Options []TrackingOption `json:"Options,omitempty" xml:"-"`
}

//TrackingCategories is a collection of TrackingCategories
type TrackingCategories struct {
	TrackingCategories []TrackingCategory `json:"TrackingCategories" xml:"TrackingCategory"`
}

func unmarshalTrackingCategory(trackingCategoryResponseBytes []byte) (*TrackingCategories, error) {
	var trackingCategoryResponse *TrackingCategories
	err := json.Unmarshal(trackingCategoryResponseBytes, &trackingCategoryResponse)
	if err != nil {
		return nil, err
	}

	return trackingCategoryResponse, err
}

//Create will create trackingCategories given an TrackingCategories struct
func (t *TrackingCategories) Create(provider *xerogolang.Provider, session goth.Session) (*TrackingCategories, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(t, "  ", "	")
	if err != nil {
		return nil, err
	}

	trackingCategoryResponseBytes, err := provider.Create(session, "TrackingCategories", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalTrackingCategory(trackingCategoryResponseBytes)
}

//Update will update an trackingCategory given an TrackingCategories struct
//This will only handle single trackingCategory - you cannot update multiple trackingCategories in a single call
func (t *TrackingCategories) Update(provider *xerogolang.Provider, session goth.Session) (*TrackingCategories, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(t, "  ", "	")
	if err != nil {
		return nil, err
	}

	trackingCategoryResponseBytes, err := provider.Update(session, "TrackingCategories/"+t.TrackingCategories[0].TrackingCategoryID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalTrackingCategory(trackingCategoryResponseBytes)
}

//FindTrackingCategories will get all trackingCategories
func FindTrackingCategories(provider *xerogolang.Provider, session goth.Session) (*TrackingCategories, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	trackingCategoryResponseBytes, err := provider.Find(session, "TrackingCategories", additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalTrackingCategory(trackingCategoryResponseBytes)
}

//FindTrackingCategory will get a single trackingCategory - trackingCategoryID must be a GUID for an trackingCategory
func FindTrackingCategory(provider *xerogolang.Provider, session goth.Session, trackingCategoryID string) (*TrackingCategories, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	trackingCategoryResponseBytes, err := provider.Find(session, "TrackingCategories/"+trackingCategoryID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalTrackingCategory(trackingCategoryResponseBytes)
}

//RemoveTrackingCategory will get a single trackingCategory - trackingCategoryID must be a GUID for an trackingCategory
func RemoveTrackingCategory(provider *xerogolang.Provider, session goth.Session, trackingCategoryID string) (*TrackingCategories, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	trackingCategoryResponseBytes, err := provider.Remove(session, "TrackingCategories/"+trackingCategoryID, additionalHeaders)
	if err != nil {
		return nil, err
	}

	return unmarshalTrackingCategory(trackingCategoryResponseBytes)
}

//GenerateExampleTrackingCategory Generates an Example trackingCategory
func GenerateExampleTrackingCategory() *TrackingCategories {
	trackingCategory := TrackingCategory{
		Name:    "Person Responsible",
		Status:  "ACTIVE",
		Options: []TrackingOption{},
	}

	trackingCategoryCollection := &TrackingCategories{
		TrackingCategories: []TrackingCategory{},
	}

	trackingCategoryCollection.TrackingCategories = append(trackingCategoryCollection.TrackingCategories, trackingCategory)

	return trackingCategoryCollection
}
