package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//CreditNote an be raised directly against a customer or supplier,
//allowing the customer or supplier to be held in credit until a future invoice or bill is raised
type CreditNote struct {

	// See Credit Note Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// The date the credit note is issued YYYY-MM-DD.
	// If the Date element is not specified then it will default
	// to the current date based on the timezone setting of the organisation
	Date string `json:"DateString,omitempty" xml:"Date,omitempty"`

	// See Credit Note Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Invoice Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// See Invoice Line Items
	LineItems []LineItem `json:"LineItems,omitempty" xml:"LineItems>LineItem,omitempty"`

	// The subtotal of the credit note excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// The total tax on the credit note
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// The total of the Credit Note(subtotal + total tax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// UTC timestamp of last update to the credit note
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// Currency used for the Credit Note
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// Date when credit note was fully paid(UTC format)
	FullyPaidOnDate string `json:"FullyPaidOnDate,omitempty" xml:"-"`

	// Xero generated unique identifier
	CreditNoteID string `json:"CreditNoteID,omitempty" xml:"CreditNoteID,omitempty"`

	// ACCRECCREDIT – Unique alpha numeric code identifying credit note (when missing will auto-generate from your Organisation Invoice Settings)
	CreditNoteNumber string `json:"CreditNoteNumber,omitempty" xml:"CreditNoteNumber,omitempty"`

	// ACCRECCREDIT only – additional reference number
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// boolean to indicate if a credit note has been sent to a contact via the Xero app (currently read only)
	SentToContact bool `json:"SentToContact,omitempty" xml:"SentToContact,omitempty"`

	// The currency rate for a multicurrency invoice. If no rate is specified, the XE.com day rate is used
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The remaining credit balance on the Credit Note
	RemainingCredit float32 `json:"RemainingCredit,omitempty" xml:"-"`

	// See Allocations
	Allocations *[]Allocation `json:"Allocations,omitempty" xml:"-"`

	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty" xml:"BrandingThemeID,omitempty"`

	// boolean to indicate if a credit note has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"-"`
}

//CreditNotes is a collection of CreditNote
type CreditNotes struct {
	CreditNotes []CreditNote `json:"CreditNotes" xml:"CreditNote"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (c *CreditNotes) convertDates() error {
	var err error
	for n := len(c.CreditNotes) - 1; n >= 0; n-- {
		c.CreditNotes[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(c.CreditNotes[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalCreditNote(creditNoteResponseBytes []byte) (*CreditNotes, error) {
	var creditNoteResponse *CreditNotes
	err := json.Unmarshal(creditNoteResponseBytes, &creditNoteResponse)
	if err != nil {
		return nil, err
	}

	err = creditNoteResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return creditNoteResponse, err
}

//Create will create creditNotes given an CreditNotes struct
func (c *CreditNotes) Create(provider *xerogolang.Provider, session goth.Session) (*CreditNotes, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(c, "  ", "	")
	if err != nil {
		return nil, err
	}

	creditNoteResponseBytes, err := provider.Create(session, "CreditNotes", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalCreditNote(creditNoteResponseBytes)
}

//Update will update an creditNote given an CreditNotes struct
//This will only handle single creditNote - you cannot update multiple creditNotes in a single call
func (c *CreditNotes) Update(provider *xerogolang.Provider, session goth.Session) (*CreditNotes, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(c, "  ", "	")
	if err != nil {
		return nil, err
	}

	creditNoteResponseBytes, err := provider.Update(session, "CreditNotes/"+c.CreditNotes[0].CreditNoteID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalCreditNote(creditNoteResponseBytes)
}

//FindCreditNotesModifiedSince will get all Credit Notes modified after a specified date.
//These Credit Notes will not have details like line items by default.
//If you need details then then add a 'page' querystringParameter and get 100 Credit Notes at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindCreditNotesModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*CreditNotes, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	creditNoteResponseBytes, err := provider.Find(session, "CreditNotes", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalCreditNote(creditNoteResponseBytes)
}

//FindCreditNotes will get all CreditNotes. These Credit Notes will not have details like line items by default.
//If you need details then then add a 'page' querystringParameter and get 100 Credit Notes at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindCreditNotes(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*CreditNotes, error) {
	return FindCreditNotesModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindCreditNote will get a single creditNote - creditNoteID can be a GUID for a creditNote or a creditNote number
func FindCreditNote(provider *xerogolang.Provider, session goth.Session, creditNoteID string) (*CreditNotes, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	creditNoteResponseBytes, err := provider.Find(session, "CreditNotes/"+creditNoteID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalCreditNote(creditNoteResponseBytes)
}

//GenerateExampleCreditNote Creates an Example creditNote
func GenerateExampleCreditNote() *CreditNotes {
	lineItem := LineItem{
		Description: "Refund Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	creditNote := CreditNote{
		Type: "ACCRECCREDIT",
		Contact: Contact{
			Name: "George Costanza",
		},
		Date:            helpers.TodayRFC3339(),
		LineAmountTypes: "Exclusive",
		LineItems:       []LineItem{},
	}

	creditNote.LineItems = append(creditNote.LineItems, lineItem)

	creditNoteCollection := &CreditNotes{
		CreditNotes: []CreditNote{},
	}

	creditNoteCollection.CreditNotes = append(creditNoteCollection.CreditNotes, creditNote)

	return creditNoteCollection
}
