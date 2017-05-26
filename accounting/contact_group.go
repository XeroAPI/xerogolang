package accounting

import (
	"encoding/json"
	"encoding/xml"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/markbates/goth"
)

//ContactGroup is a way of organising Contacts into groups
type ContactGroup struct {

	// The Name of the contact group. Required when creating a new contact group
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// The Status of a contact group. To delete a contact group update the status to DELETED. Only contact groups with a status of ACTIVE are returned on GETs.
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// The Xero identifier for an contact group – specified as a string following the endpoint name. e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9
	ContactGroupID string `json:"ContactGroupID,omitempty" xml:"ContactGroupID,omitempty"`

	// The ContactID and Name of Contacts in a contact group. Returned on GETs when the ContactGroupID is supplied in the URL.
	Contacts []Contact `json:"Contacts,omitempty" xml:"-"`
}

//ContactGroups is a collection of ContactGroups
type ContactGroups struct {
	ContactGroups []ContactGroup `json:"ContactGroups" xml:"ContactGroup"`
}

func unmarshalContactGroup(contactGroupResponseBytes []byte) (*ContactGroups, error) {
	var contactGroupResponse *ContactGroups
	err := json.Unmarshal(contactGroupResponseBytes, &contactGroupResponse)
	if err != nil {
		return nil, err
	}

	return contactGroupResponse, err
}

//CreateContactGroup will create contactGroups given an ContactGroups struct
func (c *ContactGroups) CreateContactGroup(provider *xero.Provider, session goth.Session) (*ContactGroups, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(c, "  ", "	")
	if err != nil {
		return nil, err
	}

	contactGroupResponseBytes, err := provider.Create(session, "ContactGroups", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalContactGroup(contactGroupResponseBytes)
}

//UpdateContactGroup will update an contactGroup given an ContactGroups struct
//This will only handle single contactGroup - you cannot update multiple contactGroups in a single call
func (c *ContactGroups) UpdateContactGroup(provider *xero.Provider, session goth.Session) (*ContactGroups, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(c, "  ", "	")
	if err != nil {
		return nil, err
	}

	contactGroupResponseBytes, err := provider.Update(session, "ContactGroups/"+c.ContactGroups[0].ContactGroupID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalContactGroup(contactGroupResponseBytes)
}

//FindContactGroups will get all contactGroups
func FindContactGroups(provider *xero.Provider, session goth.Session) (*ContactGroups, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	contactGroupResponseBytes, err := provider.Find(session, "ContactGroups", additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalContactGroup(contactGroupResponseBytes)
}

//FindContactGroup will get a single contactGroup - contactGroupID must be a GUID for an contactGroup
func FindContactGroup(provider *xero.Provider, session goth.Session, contactGroupID string) (*ContactGroups, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	contactGroupResponseBytes, err := provider.Find(session, "ContactGroups/"+contactGroupID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalContactGroup(contactGroupResponseBytes)
}

//RemoveContactGroup will get a single contactGroup - contactGroupID must be a GUID for an contactGroup
func RemoveContactGroup(provider *xero.Provider, session goth.Session, contactGroupID string) (*ContactGroups, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	contactGroupResponseBytes, err := provider.Remove(session, "ContactGroups/"+contactGroupID, additionalHeaders)
	if err != nil {
		return nil, err
	}

	return unmarshalContactGroup(contactGroupResponseBytes)
}

//CreateExampleContactGroup Creates an Example contactGroup
func CreateExampleContactGroup() *ContactGroups {
	contactGroup := ContactGroup{
		Name:   "Festivus Supporters",
		Status: "ACTIVE",
	}

	contactGroupCollection := &ContactGroups{
		ContactGroups: []ContactGroup{},
	}

	contactGroupCollection.ContactGroups = append(contactGroupCollection.ContactGroups, contactGroup)

	return contactGroupCollection
}
