package accounting

type ContactGroup struct {

	// The Name of the contact group. Required when creating a new contact group
	Name string `json:"Name,omitempty"`

	// The Status of a contact group. To delete a contact group update the status to DELETED. Only contact groups with a status of ACTIVE are returned on GETs.
	Status string `json:"Status,omitempty"`

	// The Xero identifier for an contact group – specified as a string following the endpoint name. e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9
	ContactGroupID string `json:"ContactGroupID,omitempty"`

	// The ContactID and Name of Contacts in a contact group. Returned on GETs when the ContactGroupID is supplied in the URL.
	Contacts []Contact `json:"Contacts,omitempty"`
}
