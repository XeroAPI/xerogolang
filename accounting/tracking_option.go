package accounting

//TrackingOption is an option from within a Tracking category
type TrackingOption struct {

	// The Xero identifier for a tracking optione.g. ae777a87-5ef3-4fa0-a4f0-d10e1f13073a
	TrackingOptionID string `json:"TrackingOptionID,omitempty" xml:"TrackingOptionID,omitempty"`

	// The name of the tracking option e.g. Marketing, East (max length = 50)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// The status of a tracking option
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Filter by a tracking categorye.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	TrackingCategoryID string `json:"TrackingCategoryID,omitempty" xml:"TrackingCategoryID,omitempty"`
}
