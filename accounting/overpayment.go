package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//Overpayment is used when a debtor overpays an invoice
type Overpayment struct {

	// See Overpayment Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// The date the overpayment is created YYYY-MM-DD
	Date string `json:"DateString,omitempty" xml:"Date,omitempty"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// See Overpayment Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Overpayment Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// See Overpayment Line Items
	LineItems []LineItem `json:"LineItems,omitempty" xml:"LineItems,omitempty"`

	// The subtotal of the overpayment excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// The total tax on the overpayment
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// The total of the overpayment (subtotal + total tax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// UTC timestamp of last update to the overpayment
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`

	// Currency used for the overpayment
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// Xero generated unique identifier
	OverpaymentID string `json:"OverpaymentID,omitempty" xml:"OverpaymentID,omitempty"`

	// The currency rate for a multicurrency overpayment. If no rate is specified, the XE.com day rate is used
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The remaining credit balance on the overpayment
	RemainingCredit float32 `json:"RemainingCredit,omitempty" xml:"RemainingCredit,omitempty"`

	// See Allocations
	Allocations []Allocation `json:"Allocations,omitempty" xml:"Allocations,omitempty"`

	// See Payments
	Payments []Payment `json:"Payments,omitempty" xml:"Payments,omitempty"`

	// boolean to indicate if a overpayment has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`
}

//Overpayments is a collection of Overpayments
type Overpayments struct {
	Overpayments []Overpayment `json:"Overpayments" xml:"Overpayment"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (o *Overpayments) convertDates() error {
	var err error
	for n := len(o.Overpayments) - 1; n >= 0; n-- {
		o.Overpayments[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(o.Overpayments[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalOverpayment(overpaymentResponseBytes []byte) (*Overpayments, error) {
	var overpaymentResponse *Overpayments
	err := json.Unmarshal(overpaymentResponseBytes, &overpaymentResponse)
	if err != nil {
		return nil, err
	}

	err = overpaymentResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return overpaymentResponse, err
}

//FindOverpaymentsModifiedSince will get all Overpayments modified after a specified date.
//These Overpayments will not have details like default line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Overpayments at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindOverpaymentsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Overpayments, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	overpaymentResponseBytes, err := provider.Find(session, "Overpayments", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalOverpayment(overpaymentResponseBytes)
}

//FindOverpayments will get all Overpayments. These Overpayment will not have details like line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Overpayments at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindOverpayments(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*Overpayments, error) {
	return FindOverpaymentsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindOverpayment will get a single overpayment - overpaymentID can be a GUID for an overpayment or an overpayment number
func FindOverpayment(provider *xerogolang.Provider, session goth.Session, overpaymentID string) (*Overpayments, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	overpaymentResponseBytes, err := provider.Find(session, "Overpayments/"+overpaymentID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalOverpayment(overpaymentResponseBytes)
}

//Allocate allocates an overpayment - to create an overpayment
//use the bankTransactions endpoint.
func (o *Overpayments) Allocate(provider *xerogolang.Provider, session goth.Session, allocations Allocations) (*Overpayments, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(allocations, "  ", "	")
	if err != nil {
		return nil, err
	}

	overpaymentResponseBytes, err := provider.Create(session, "Overpayments/"+o.Overpayments[0].OverpaymentID+"/Allocations", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalOverpayment(overpaymentResponseBytes)
}
