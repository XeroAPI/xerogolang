package accounting

type ExternalLink struct {

	// See External link types
	LinkType string `json:"LinkType,omitempty"`

	// URL for service e.g. http://twitter.com/xeroapi
	URL string `json:"Url,omitempty"`
}
