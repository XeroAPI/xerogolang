package accounting

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/helpers"
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
func (c *CreditNotes) convertCreditNoteDates() error {
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

	err = creditNoteResponse.convertCreditNoteDates()
	if err != nil {
		return nil, err
	}

	return creditNoteResponse, err
}

//CreateCreditNote will create creditNotes given an CreditNotes struct
func (c *CreditNotes) CreateCreditNote(provider *xero.Provider, session goth.Session) (*CreditNotes, error) {
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

//UpdateCreditNote will update an creditNote given an CreditNotes struct
//This will only handle single creditNote - you cannot update multiple creditNotes in a single call
func (c *CreditNotes) UpdateCreditNote(provider *xero.Provider, session goth.Session) (*CreditNotes, error) {
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

//FindCreditNotesModifiedSinceWithParams will get all Credit Notes modified after a specified date.
//These Credit Notes will not have details like default account codes and tracking categories.
//If you need details then use FindCreditNotesByPage and get 100 Credit Notes at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindCreditNotesModifiedSinceWithParams(provider *xero.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*CreditNotes, error) {
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

//FindCreditNotesModifiedSince will get all Credit Notes modified after a specified date.
//These Credit Notes will not have details like default account codes and tracking categories.
//If you need details then use FindCreditNotesByPage and get 100 CreditNotes at a time
func FindCreditNotesModifiedSince(provider *xero.Provider, session goth.Session, modifiedSince time.Time) (*CreditNotes, error) {
	return FindCreditNotesModifiedSinceWithParams(provider, session, modifiedSince, nil)
}

//FindCreditNotesModifiedSinceByPage will get a specified page of Credit Notes which contains 100 Credit Notes modified
//after a specified date. Page 1 gives the first 100, page two the next 100 etc etc.
//Paged Credit Notes contain all the detail of the Credit Notes whereas if you use FindAllCreditNotes
//you will only get summarised data e.g. no line items or tracking categories
func FindCreditNotesModifiedSinceByPage(provider *xero.Provider, session goth.Session, modifiedSince time.Time, page int) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"page": strconv.Itoa(page),
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindCreditNotesModifiedSinceWhere will get Credit Notes which contains 100 Credit Notes
//that fit the criteria of a supplied where clause.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindCreditNotesByPage and get 100 Credit Notes at a time
func FindCreditNotesModifiedSinceWhere(provider *xero.Provider, session goth.Session, modifiedSince time.Time, whereClause string) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindCreditNotesModifiedSinceOrderedBy will get Credit Notes and are order them by a supplied named element.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindCreditNotesByPage and get 100 Credit Notes at a time
func FindCreditNotesModifiedSinceOrderedBy(provider *xero.Provider, session goth.Session, modifiedSince time.Time, orderBy string) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindCreditNotesByPage will get a specified page of Credit Notes which contains 100 Credit Notes
//Page 1 gives the first 100, page two the next 100 etc etc.
//paged Credit Notes contain all the detail of the Credit Notes whereas if you use FindAllCreditNotes
//you will only get summarised data e.g. no line items or tracking categories
func FindCreditNotesByPage(provider *xero.Provider, session goth.Session, page int) (*CreditNotes, error) {
	return FindCreditNotesModifiedSinceByPage(provider, session, dayZero, page)
}

//FindCreditNotesByPageWhere will get a specified page of Credit Notes which contains 100 Credit Notes
//that fit the criteria of a supplied where clause. Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged Credit Notes contain all the detail of the Credit Notes whereas if you use FindAllCreditNotes
//you will only get summarised data e.g. no line items or tracking categories
func FindCreditNotesByPageWhere(provider *xero.Provider, session goth.Session, page int, whereClause string) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindCreditNotesByPageWhereOrderedBy will get a specified page of Credit Notes which contains 100 Credit Notes
//that fit the criteria of a supplied where clause and are ordered by a supplied named element.
//Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged Credit Notes contain all the detail of the Credit Notes whereas if you use FindCreditNotes
//you will only get summarised data e.g. no line items or tracking categories
func FindCreditNotesByPageWhereOrderedBy(provider *xero.Provider, session goth.Session, page int, whereClause string, orderBy string) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
		"order": orderBy,
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindCreditNotesOrderedBy will get all Credit Notes ordered by a supplied named element.
//These Credit Notes will not have details like line items.
//If you need details then use FindCreditNotesByPage and get 100 CreditNotes at a time
func FindCreditNotesOrderedBy(provider *xero.Provider, session goth.Session, orderBy string) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindCreditNotesWhere will get all Credit Notes that fit the criteria of a supplied where clause.
//These Credit Notes will not have details like line items.
//If you need details then use FindCreditNotesByPage and get 100 CreditNotes at a time
func FindCreditNotesWhere(provider *xero.Provider, session goth.Session, whereClause string) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindCreditNotesWhereOrderedBy will get all Credit Notes that fit the criteria of a supplied where clause
//and are ordered by a supplied named element. These Credit Notes will not have details like line items.
//If you need details then use FindCreditNotesByPage and get 100 CreditNotes at a time
func FindCreditNotesWhereOrderedBy(provider *xero.Provider, session goth.Session, whereClause string, orderedBy string) (*CreditNotes, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
		"order": orderedBy,
	}
	return FindCreditNotesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindCreditNotesWithParams will get all Credit Notes. These Credit Notes will not have details like line items.
//If you need details then use FindCreditNotesByPage and get 100 CreditNotes at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindCreditNotesWithParams(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*CreditNotes, error) {
	return FindCreditNotesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindCreditNotes will get all CreditNotes. These Credit Notes will not have details like line items.
//If you need details then use FindCreditNotesByPage and get 100 CreditNotes at a time
func FindCreditNotes(provider *xero.Provider, session goth.Session) (*CreditNotes, error) {
	return FindCreditNotesModifiedSinceWithParams(provider, session, dayZero, nil)
}

//FindCreditNote will get a single creditNote - creditNoteID can be a GUID for an creditNote or an creditNote number
func FindCreditNote(provider *xero.Provider, session goth.Session, creditNoteID string) (*CreditNotes, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	creditNoteResponseBytes, err := provider.Find(session, "CreditNotes/"+creditNoteID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalCreditNote(creditNoteResponseBytes)
}

//CreateExampleCreditNote Creates an Example creditNote
func CreateExampleCreditNote() *CreditNotes {
	lineItem := LineItem{
		Description: "Refund Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	creditNote := CreditNote{
		Type: "ACCRECCREDIT",
		Contact: Contact{
			Name: "George Costanza",
		},
		Date:            strings.TrimSuffix(today.Format(time.RFC3339), "Z"),
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
