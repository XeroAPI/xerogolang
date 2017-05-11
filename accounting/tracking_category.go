package accounting

//TrackingCategory is used to segment data within a Xero organisation
type TrackingCategory struct {

	// The Xero identifier for a tracking categorye.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	TrackingCategoryID string `json:"TrackingCategoryID,omitempty" xml:"TrackingCategoryID,omitempty"`

	// The name of the tracking category e.g. Department, Region (max length = 100)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// The status of a tracking category
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Tracking Options
	Options []TrackingOption `json:"Options,omitempty" xml:"Options,omitempty"`
}
