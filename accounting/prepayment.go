package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//Prepayment are payments made before the associated document has been created
type Prepayment struct {

	// See Prepayment Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// The date the prepayment is created YYYY-MM-DD
	Date string `json:"DateString,omitempty" xml:"Date,omitempty"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// See Prepayment Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Prepayment Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// See Prepayment Line Items
	LineItems []LineItem `json:"LineItems,omitempty" xml:"LineItems,omitempty"`

	// The subtotal of the prepayment excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// The total tax on the prepayment
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// The total of the prepayment(subtotal + total tax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// UTC timestamp of last update to the prepayment
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`

	// Currency used for the prepayment
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// Xero generated unique identifier
	PrepaymentID string `json:"PrepaymentID,omitempty" xml:"PrepaymentID,omitempty"`

	// The currency rate for a multicurrency prepayment. If no rate is specified, the XE.com day rate is used
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The remaining credit balance on the prepayment
	RemainingCredit float32 `json:"RemainingCredit,omitempty" xml:"RemainingCredit,omitempty"`

	// See Allocations
	Allocations []Allocation `json:"Allocations,omitempty" xml:"Allocations,omitempty"`

	// boolean to indicate if a prepayment has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`
}

//Prepayments is a collection of Prepayments
type Prepayments struct {
	Prepayments []Prepayment `json:"Prepayments" xml:"Prepayment"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (p *Prepayments) convertDates() error {
	var err error
	for n := len(p.Prepayments) - 1; n >= 0; n-- {
		p.Prepayments[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(p.Prepayments[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalPrepayment(prepaymentResponseBytes []byte) (*Prepayments, error) {
	var prepaymentResponse *Prepayments
	err := json.Unmarshal(prepaymentResponseBytes, &prepaymentResponse)
	if err != nil {
		return nil, err
	}

	err = prepaymentResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return prepaymentResponse, err
}

//FindPrepaymentsModifiedSince will get all Prepayments modified after a specified date.
//These Prepayments will not have details like default line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Prepayments at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindPrepaymentsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Prepayments, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	prepaymentResponseBytes, err := provider.Find(session, "Prepayments", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalPrepayment(prepaymentResponseBytes)
}

//FindPrepayments will get all Prepayments. These Prepayment will not have details like line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Prepayments at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindPrepayments(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*Prepayments, error) {
	return FindPrepaymentsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindPrepayment will get a single prepayment - prepaymentID can be a GUID for an prepayment or an prepayment number
func FindPrepayment(provider *xerogolang.Provider, session goth.Session, prepaymentID string) (*Prepayments, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	prepaymentResponseBytes, err := provider.Find(session, "Prepayments/"+prepaymentID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalPrepayment(prepaymentResponseBytes)
}

//Allocate allocates a prepayment - to create a prepayment
//use the bankTransactions endpoint.
func (p *Prepayments) Allocate(provider *xerogolang.Provider, session goth.Session, allocations Allocations) (*Prepayments, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(allocations, "  ", "	")
	if err != nil {
		return nil, err
	}

	prepaymentResponseBytes, err := provider.Create(session, "Prepayments/"+p.Prepayments[0].PrepaymentID+"/Allocations", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalPrepayment(prepaymentResponseBytes)
}
