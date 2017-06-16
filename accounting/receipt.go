package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//Receipt is a record of personal money spent that a user would like to be reimbursed for - see also ExpenseClaim
type Receipt struct {
	//See Users
	User User `json:"User" xml:"User"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// Date of receipt – YYYY-MM-DD
	Date string `json:"Date" xml:"Date"`

	// See LineItems
	LineItems []LineItem `json:"LineItems" xml:"LineItems>LineItem"`

	// Additional reference number
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// Total of receipt excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// Total tax on receipt
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// Total of receipt tax inclusive (i.e. SubTotal + TotalTax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// Xero generated unique identifier for receipt
	ReceiptID string `json:"ReceiptID,omitempty" xml:"ReceiptID,omitempty"`

	// Current status of receipt – see status types
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Xero generated sequence number for receipt in current claim for a given user
	ReceiptNumber int `json:"ReceiptNumber,omitempty" xml:"ReceiptNumber,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// boolean to indicate if a receipt has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`

	// URL link to a source document – shown as “Go to [appName]” in the Xero app
	URL string `json:"Url,omitempty" xml:"Url,omitempty"`
}

//Receipts contains a collection of Receipts
type Receipts struct {
	Receipts []Receipt `json:"Receipts" xml:"Receipt"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (r *Receipts) convertDates() error {
	var err error
	for n := len(r.Receipts) - 1; n >= 0; n-- {
		r.Receipts[n].Date, err = helpers.DotNetJSONTimeToRFC3339(r.Receipts[n].Date, false)
		if err != nil {
			return err
		}
		r.Receipts[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(r.Receipts[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalReceipt(receiptResponseBytes []byte) (*Receipts, error) {
	var receiptResponse *Receipts
	err := json.Unmarshal(receiptResponseBytes, &receiptResponse)
	if err != nil {
		return nil, err
	}

	err = receiptResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return receiptResponse, err
}

//Create will create receipts given an Receipts struct
func (r *Receipts) Create(provider *xerogolang.Provider, session goth.Session) (*Receipts, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(r, "  ", "	")
	if err != nil {
		return nil, err
	}

	receiptResponseBytes, err := provider.Create(session, "Receipts", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalReceipt(receiptResponseBytes)
}

//Update will update an receipt given an Receipts struct
//This will only handle single receipt - you cannot update multiple receipts in a single call
func (r *Receipts) Update(provider *xerogolang.Provider, session goth.Session) (*Receipts, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	//strip out user and contact details because we only need to send an ID
	u := User{
		UserID: r.Receipts[0].User.UserID,
	}
	r.Receipts[0].User = u

	c := Contact{
		ContactID: r.Receipts[0].Contact.ContactID,
	}
	r.Receipts[0].Contact = c

	body, err := xml.MarshalIndent(r, "  ", "	")
	if err != nil {
		return nil, err
	}

	receiptResponseBytes, err := provider.Update(session, "Receipts/"+r.Receipts[0].ReceiptID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalReceipt(receiptResponseBytes)
}

//FindReceiptsModifiedSince will get all Receipts modified after a specified date.
//These Receipts will not have details like default line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Receipts at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindReceiptsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Receipts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	receiptResponseBytes, err := provider.Find(session, "Receipts", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReceipt(receiptResponseBytes)
}

//FindReceipts will get all Receipts. These Receipt will not have details like line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Receipts at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindReceipts(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*Receipts, error) {
	return FindReceiptsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindReceipt will get a single receipt - receiptID can be a GUID for an receipt or an receipt number
func FindReceipt(provider *xerogolang.Provider, session goth.Session, receiptID string) (*Receipts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	receiptResponseBytes, err := provider.Find(session, "Receipts/"+receiptID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalReceipt(receiptResponseBytes)
}

//GenerateExampleReceipt Creates an Example receipt
func GenerateExampleReceipt(userID string, contactID string) *Receipts {
	lineItem := LineItem{
		Description: "Lunch at the Dream Cafe",
		Quantity:    1.00,
		UnitAmount:  55.00,
		AccountCode: "400",
	}

	receipt := Receipt{
		User: User{
			UserID: userID,
		},
		Contact: Contact{
			ContactID: contactID,
		},
		Date:            helpers.TodayRFC3339(),
		LineAmountTypes: "Inclusive",
		LineItems:       []LineItem{},
	}

	receipt.LineItems = append(receipt.LineItems, lineItem)

	receiptCollection := &Receipts{
		Receipts: []Receipt{},
	}

	receiptCollection.Receipts = append(receiptCollection.Receipts, receipt)

	return receiptCollection
}
