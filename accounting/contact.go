package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//Contact is a debtor/customer or creditor/supplier in a Xero Organisation
type Contact struct {

	// Xero identifier
	ContactID string `json:"ContactID,omitempty" xml:"ContactID,omitempty"`

	// This can be updated via the API only i.e. This field is read only on the Xero contact screen, used to identify contacts in external systems (max length = 50). If the Contact Number is used, this is displayed as Contact Code in the Contacts UI in Xero.
	ContactNumber string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`

	// A user defined account number. This can be updated via the API and the Xero UI (max length = 50)
	AccountNumber string `json:"AccountNumber,omitempty" xml:"AccountNumber,omitempty"`

	// Current status of a contact – see contact status types
	ContactStatus string `json:"ContactStatus,omitempty" xml:"ContactStatus,omitempty"`

	// Full name of contact/organisation (max length = 255)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// First name of contact person (max length = 255)
	FirstName string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`

	// Last name of contact person (max length = 255)
	LastName string `json:"LastName,omitempty" xml:"LastName,omitempty"`

	// Email address of contact person (umlauts not supported) (max length = 255)
	EmailAddress string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`

	// Skype user name of contact
	SkypeUserName string `json:"SkypeUserName,omitempty" xml:"SkypeUserName,omitempty"`

	// See contact persons
	ContactPersons *[]ContactPerson `json:"ContactPersons,omitempty" xml:"ContactPersons>ContactPerson,omitempty"`

	// Bank account number of contact
	BankAccountDetails string `json:"BankAccountDetails,omitempty" xml:"BankAccountDetails,omitempty"`

	// Tax number of contact – this is also known as the ABN (Australia), GST Number (New Zealand), VAT Number (UK) or Tax ID Number (US and global) in the Xero UI depending on which regionalized version of Xero you are using (max length = 50)
	TaxNumber string `json:"TaxNumber,omitempty" xml:"TaxNumber,omitempty"`

	// Default tax type used for contact on AR Contacts
	AccountsReceivableTaxType string `json:"AccountsReceivableTaxType,omitempty" xml:"AccountsReceivableTaxType,omitempty"`

	// Default tax type used for contact on AP Contacts
	AccountsPayableTaxType string `json:"AccountsPayableTaxType,omitempty" xml:"AccountsPayableTaxType,omitempty"`

	// Store certain address types for a contact – see address types
	Addresses *[]Address `json:"Addresses,omitempty" xml:"Addresses>Address,omitempty"`

	// Store certain phone types for a contact – see phone types
	Phones *[]Phone `json:"Phones,omitempty" xml:"Phones>Phone,omitempty"`

	// true or false – Boolean that describes if a contact that has any AP Contacts entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts payable Contact is generated against this contact.
	IsSupplier bool `json:"IsSupplier,omitempty" xml:"IsSupplier,omitempty"`

	// true or false – Boolean that describes if a contact has any AR Contacts entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts receivable Contact is generated against this contact.
	IsCustomer bool `json:"IsCustomer,omitempty" xml:"IsCustomer,omitempty"`

	// Default currency for raising Contacts against contact
	DefaultCurrency string `json:"DefaultCurrency,omitempty" xml:"DefaultCurrency,omitempty"`

	// Store XeroNetworkKey for contacts.
	XeroNetworkKey string `json:"XeroNetworkKey,omitempty" xml:"XeroNetworkKey,omitempty"`

	// The default sales account code for contacts
	SalesDefaultAccountCode string `json:"SalesDefaultAccountCode,omitempty" xml:"SalesDefaultAccountCode,omitempty"`

	// The default purchases account code for contacts
	PurchasesDefaultAccountCode string `json:"PurchasesDefaultAccountCode,omitempty" xml:"PurchasesDefaultAccountCode,omitempty"`

	// The default sales tracking categories for contacts
	SalesTrackingCategories *[]TrackingCategory `json:"SalesTrackingCategories,omitempty" xml:"SalesTrackingCategories>SalesTrackingCategory,omitempty"`

	// The default purchases tracking categories for contacts
	PurchasesTrackingCategories *[]TrackingCategory `json:"PurchasesTrackingCategories,omitempty" xml:"PurchasesTrackingCategories>PurchaseTrackingCategory,omitempty"`

	// The name of the Tracking Category assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories
	TrackingCategoryName string `json:"TrackingCategoryName,omitempty" xml:"TrackingCategoryName,omitempty"`

	// The name of the Tracking Option assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories
	TrackingCategoryOption string `json:"TrackingCategoryOption,omitempty" xml:"TrackingCategoryOption,omitempty"`

	// UTC timestamp of last update to contact
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// Displays which contact groups a contact is included in
	ContactGroups *[]ContactGroup `json:"ContactGroups,omitempty" xml:"ContactGroups>ContactGroup,omitempty"`

	// Website address for contact (read only)
	Website string `json:"Website,omitempty" xml:"-"`

	// batch payment details for contact (read only)
	BatchPayments BatchPayment `json:"BatchPayments,omitempty" xml:"-"`

	// The default discount rate for the contact (read only)
	Discount float32 `json:"Discount,omitempty" xml:"-"`

	// The raw AccountsReceivable(sales Contacts) and AccountsPayable(bills) outstanding and overdue amounts, not converted to base currency (read only)
	Balances Balances `json:"Balances,omitempty" xml:"-"`

	// A boolean to indicate if a contact has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`
}

//Contacts contains a collection of Contacts
type Contacts struct {
	Contacts []Contact `json:"Contacts" xml:"Contact"`
}

//Balances are the raw AccountsReceivable(sales invoices) and AccountsPayable(bills)
//outstanding and overdue amounts, not converted to base currency
type Balances struct {
	AccountsReceivable Balance `json:"AccountsReceivable,omitempty" xml:"AccountsReceivable,omitempty"`
	AccountsPayable    Balance `json:"AccountsPayable,omitempty" xml:"AccountsPayable,omitempty"`
}

//Balance is the raw AccountsReceivable(sales invoices) and AccountsPayable(bills)
//outstanding and overdue amounts, not converted to base currency
type Balance struct {
	Outstanding float32 `json:"Oustanding,omitempty" xml:"Oustanding,omitempty"`
	Overdue     float32 `json:"Overdue,omitempty" xml:"Overdue,omitempty"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (c *Contacts) convertDates() error {
	var err error
	for n := len(c.Contacts) - 1; n >= 0; n-- {
		c.Contacts[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(c.Contacts[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalContact(contactResponseBytes []byte) (*Contacts, error) {
	var contactResponse *Contacts
	err := json.Unmarshal(contactResponseBytes, &contactResponse)
	if err != nil {
		return nil, err
	}

	err = contactResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return contactResponse, err
}

//Create will create Contacts given an Contacts struct
func (c *Contacts) Create(provider *xerogolang.Provider, session goth.Session) (*Contacts, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(c, "  ", "	")
	if err != nil {
		return nil, err
	}

	contactResponseBytes, err := provider.Create(session, "Contacts", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalContact(contactResponseBytes)
}

//Update will update a Contact given a Contacts struct
//This will only handle single Contact - you cannot update multiple Contacts in a single call
func (c *Contacts) Update(provider *xerogolang.Provider, session goth.Session) (*Contacts, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(c, "  ", "	")
	if err != nil {
		return nil, err
	}

	contactResponseBytes, err := provider.Update(session, "Contacts/"+c.Contacts[0].ContactID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalContact(contactResponseBytes)
}

//FindContactsModifiedSince will get all Contacts modified after a specified date.
//These Contacts will not have details like default account codes and tracking categories.
//If you need details then then add a 'page' querystringParameter and get 100 Contacts at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindContactsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Contacts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	contactResponseBytes, err := provider.Find(session, "Contacts", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalContact(contactResponseBytes)
}

//FindContacts will get all Contacts. These Contact will not have details like default accounts.
//If you need details then then add a 'page' querystringParameter and get 100 Contacts at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindContacts(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*Contacts, error) {
	return FindContactsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindContact will get a single Contact - ContactID can be a GUID for an Contact or an Contact number
func FindContact(provider *xerogolang.Provider, session goth.Session, contactID string) (*Contacts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	contactResponseBytes, err := provider.Find(session, "Contacts/"+contactID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalContact(contactResponseBytes)
}

//AddToContactGroup will add a collection of Contacts to a supplied contactGroupID
func (c *Contacts) AddToContactGroup(provider *xerogolang.Provider, session goth.Session, contactGroupID string) (*Contacts, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	//We only want to send ContactID's or the endpoint will return a 400 so we need to strip out all the other contact info
	contactsToAdd := []Contact{}
	for _, contact := range c.Contacts {
		contactToAdd := Contact{
			ContactID: contact.ContactID,
		}
		contactsToAdd = append(contactsToAdd, contactToAdd)
	}

	body, err := xml.MarshalIndent(contactsToAdd, "  ", "	")
	if err != nil {
		return nil, err
	}

	contactResponseBytes, err := provider.Update(session, "ContactGroups/"+contactGroupID+"/Contacts", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalContact(contactResponseBytes)
}

//RemoveFromContactGroup will remove a Contact from a supplied contactGroupID - must be done one at a time.
func (c *Contacts) RemoveFromContactGroup(provider *xerogolang.Provider, session goth.Session, contactGroupID string) (*Contacts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	contactResponseBytes, err := provider.Remove(session, "ContactGroups/"+contactGroupID+"/Contacts/"+c.Contacts[0].ContactID, additionalHeaders)
	if err != nil {
		return nil, err
	}

	return unmarshalContact(contactResponseBytes)
}

//GenerateExampleContact Creates an Example contact
func GenerateExampleContact() *Contacts {
	contact := Contact{
		Name: "Cosmo Kramer",
	}

	contactCollection := &Contacts{
		Contacts: []Contact{},
	}

	contactCollection.Contacts = append(contactCollection.Contacts, contact)

	return contactCollection
}
