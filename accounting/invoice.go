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

//Invoice is an Accounts Payable or Accounts Recievable document in a Xero organisation
type Invoice struct {
	// See Invoice Types
	Type string `json:"Type" xml:"Type"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// See LineItems
	LineItems []LineItem `json:"LineItems" xml:"LineItems>LineItem"`

	// Date invoice was issued – YYYY-MM-DD. If the Date element is not specified it will default to the current date based on the timezone setting of the organisation
	Date string `json:"DateString,omitempty" xml:"Date,omitempty"`

	// Date invoice is due – YYYY-MM-DD
	DueDate string `json:"DueDateString,omitempty" xml:"DueDate,omitempty"`

	// Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// ACCREC – Unique alpha numeric code identifying invoice (when missing will auto-generate from your Organisation Invoice Settings) (max length = 255)
	InvoiceNumber string `json:"InvoiceNumber,omitempty" xml:"InvoiceNumber,omitempty"`

	// ACCREC only – additional reference number (max length = 255)
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty" xml:"BrandingThemeID,omitempty"`

	// URL link to a source document – shown as “Go to [appName]” in the Xero app
	URL string `json:"Url,omitempty" xml:"Url,omitempty"`

	// The currency that invoice has been raised in (see Currencies)
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// The currency rate for a multicurrency invoice. If no rate is specified, the XE.com day rate is used. (max length = [18].[6])
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// See Invoice Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Boolean to set whether the invoice in the Xero app should be marked as “sent”. This can be set only on invoices that have been approved
	SentToContact bool `json:"SentToContact,omitempty" xml:"SentToContact,omitempty"`

	// Shown on sales invoices (Accounts Receivable) when this has been set
	ExpectedPaymentDate string `json:"ExpectedPaymentDate,omitempty" xml:"ExpectedPaymentDate,omitempty"`

	// Shown on bills (Accounts Payable) when this has been set
	PlannedPaymentDate string `json:"PlannedPaymentDate,omitempty" xml:"PlannedPaymentDate,omitempty"`

	// Total of invoice excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// Total tax on invoice
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// Total of Invoice tax inclusive (i.e. SubTotal + TotalTax). This will be ignored if it doesn’t equal the sum of the LineAmounts
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// Total of discounts applied on the invoice line items
	TotalDiscount float32 `json:"TotalDiscount,omitempty" xml:"-"`

	// Xero generated unique identifier for invoice
	InvoiceID string `json:"InvoiceID,omitempty" xml:"InvoiceID,omitempty"`

	// boolean to indicate if an invoice has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"-"`

	// See Payments
	Payments *[]Payment `json:"Payments,omitempty" xml:"-"`

	// See Prepayments
	Prepayments *[]Prepayment `json:"Prepayments,omitempty" xml:"-"`

	// See Overpayments
	Overpayments *[]Overpayment `json:"Overpayments,omitempty" xml:"-"`

	// Amount remaining to be paid on invoice
	AmountDue float32 `json:"AmountDue,omitempty" xml:"-"`

	// Sum of payments received for invoice
	AmountPaid float32 `json:"AmountPaid,omitempty" xml:"-"`

	// The date the invoice was fully paid. Only returned on fully paid invoices
	FullyPaidOnDate string `json:"FullyPaidOnDate,omitempty" xml:"-"`

	// Sum of all credit notes, over-payments and pre-payments applied to invoice
	AmountCredited float32 `json:"AmountCredited,omitempty" xml:"-"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// Details of credit notes that have been applied to an invoice
	CreditNotes *[]CreditNote `json:"CreditNotes,omitempty" xml:"-"`
}

//Invoices contains a collection of Invoices
type Invoices struct {
	Invoices []Invoice `json:"Invoices" xml:"Invoice"`
}

var (
	dayZero = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
)

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (i *Invoices) convertInvoiceDates() error {
	var err error
	for n := len(i.Invoices) - 1; n >= 0; n-- {
		i.Invoices[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(i.Invoices[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalInvoice(invoiceResponseBytes []byte) (*Invoices, error) {
	var invoiceResponse *Invoices
	err := json.Unmarshal(invoiceResponseBytes, &invoiceResponse)
	if err != nil {
		return nil, err
	}

	err = invoiceResponse.convertInvoiceDates()
	if err != nil {
		return nil, err
	}

	return invoiceResponse, err
}

//CreateInvoice will create invoices given an Invoices struct
func (i *Invoices) CreateInvoice(provider *xero.Provider, session goth.Session) (*Invoices, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(i, "  ", "	")
	if err != nil {
		return nil, err
	}

	invoiceResponseBytes, err := provider.Create(session, "Invoices", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalInvoice(invoiceResponseBytes)
}

//UpdateInvoice will update an invoice given an Invoices struct
//This will only handle single invoice - you cannot update multiple invoices in a single call
func (i *Invoices) UpdateInvoice(provider *xero.Provider, session goth.Session) (*Invoices, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(i, "  ", "	")
	if err != nil {
		return nil, err
	}

	invoiceResponseBytes, err := provider.Update(session, "Invoices/"+i.Invoices[0].InvoiceID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalInvoice(invoiceResponseBytes)
}

//FindInvoicesModifiedSinceWithParams will get all Invoices modified after a specified date.
//These Invoices will not have details like default account codes and tracking categories.
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindInvoicesModifiedSinceWithParams(provider *xero.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Invoices, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	invoiceResponseBytes, err := provider.Find(session, "Invoices", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalInvoice(invoiceResponseBytes)
}

//FindInvoicesModifiedSince will get all Invoices modified after a specified date.
//These Invoices will not have details like default account codes and tracking categories.
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
func FindInvoicesModifiedSince(provider *xero.Provider, session goth.Session, modifiedSince time.Time) (*Invoices, error) {
	return FindInvoicesModifiedSinceWithParams(provider, session, modifiedSince, nil)
}

//FindInvoicesModifiedSinceByPage will get a specified page of Invoices which contains 100 Invoices modified
//after a specified date. Page 1 gives the first 100, page two the next 100 etc etc.
//Paged Invoices contain all the detail of the Invoices whereas if you use FindAllInvoices
//you will only get summarised data e.g. no line items or tracking categories
func FindInvoicesModifiedSinceByPage(provider *xero.Provider, session goth.Session, modifiedSince time.Time, page int) (*Invoices, error) {
	querystringParameters := map[string]string{
		"page": strconv.Itoa(page),
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindInvoicesModifiedSinceWhere will get Invoices which contains 100 Invoices
//that fit the criteria of a supplied where clause.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
func FindInvoicesModifiedSinceWhere(provider *xero.Provider, session goth.Session, modifiedSince time.Time, whereClause string) (*Invoices, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindInvoicesModifiedSinceOrderedBy will get Invoices and are order them by a supplied named element.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
func FindInvoicesModifiedSinceOrderedBy(provider *xero.Provider, session goth.Session, modifiedSince time.Time, orderBy string) (*Invoices, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindInvoicesByPage will get a specified page of Invoices which contains 100 Invoices
//Page 1 gives the first 100, page two the next 100 etc etc.
//paged Invoices contain all the detail of the Invoices whereas if you use FindAllInvoices
//you will only get summarised data e.g. no line items or tracking categories
func FindInvoicesByPage(provider *xero.Provider, session goth.Session, page int) (*Invoices, error) {
	return FindInvoicesModifiedSinceByPage(provider, session, dayZero, page)
}

//FindInvoicesByPageWhere will get a specified page of Invoices which contains 100 Invoices
//that fit the criteria of a supplied where clause. Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged Invoices contain all the detail of the Invoices whereas if you use FindAllInvoices
//you will only get summarised data e.g. no line items or tracking categories
func FindInvoicesByPageWhere(provider *xero.Provider, session goth.Session, page int, whereClause string) (*Invoices, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindInvoicesByPageWhereOrderedBy will get a specified page of Invoices which contains 100 Invoices
//that fit the criteria of a supplied where clause and are ordered by a supplied named element.
//Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged Invoices contain all the detail of the Invoices whereas if you use FindInvoices
//you will only get summarised data e.g. no line items or tracking categories
func FindInvoicesByPageWhereOrderedBy(provider *xero.Provider, session goth.Session, page int, whereClause string, orderBy string) (*Invoices, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
		"order": orderBy,
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindInvoicesOrderedBy will get all Invoices ordered by a supplied named element.
//These Invoices will not have details like line items.
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
func FindInvoicesOrderedBy(provider *xero.Provider, session goth.Session, orderBy string) (*Invoices, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindInvoicesWhere will get all Invoices that fit the criteria of a supplied where clause.
//These Invoices will not have details like line items.
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
func FindInvoicesWhere(provider *xero.Provider, session goth.Session, whereClause string) (*Invoices, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindInvoicesWhereOrderedBy will get all Invoices that fit the criteria of a supplied where clause
//and are ordered by a supplied named element. These Invoices will not have details like line items.
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
func FindInvoicesWhereOrderedBy(provider *xero.Provider, session goth.Session, whereClause string, orderedBy string) (*Invoices, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
		"order": orderedBy,
	}
	return FindInvoicesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindInvoicesWithParams will get all Invoices. These Invoice will not have details like line items.
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindInvoicesWithParams(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*Invoices, error) {
	return FindInvoicesModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindInvoices will get all Invoices. These Invoice will not have details like line items.
//If you need details then use FindInvoicesByPage and get 100 Invoices at a time
func FindInvoices(provider *xero.Provider, session goth.Session) (*Invoices, error) {
	return FindInvoicesModifiedSinceWithParams(provider, session, dayZero, nil)
}

//FindInvoice will get a single invoice - invoiceID can be a GUID for an invoice or an invoice number
func FindInvoice(provider *xero.Provider, session goth.Session, invoiceID string) (*Invoices, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	invoiceResponseBytes, err := provider.Find(session, "Invoices/"+invoiceID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalInvoice(invoiceResponseBytes)
}

//CreateExampleInvoice Creates an Example invoice
func CreateExampleInvoice() *Invoices {
	lineItem := LineItem{
		Description: "Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	invoice := Invoice{
		Type: "ACCREC",
		Contact: Contact{
			Name: "George Costanza",
		},
		Date:            strings.TrimSuffix(today.Format(time.RFC3339), "Z"),
		DueDate:         strings.TrimSuffix(today.Add(720*time.Hour).Format(time.RFC3339), "Z"),
		LineAmountTypes: "Exclusive",
		LineItems:       []LineItem{},
	}

	invoice.LineItems = append(invoice.LineItems, lineItem)

	invoiceCollection := &Invoices{
		Invoices: []Invoice{},
	}

	invoiceCollection.Invoices = append(invoiceCollection.Invoices, invoice)

	return invoiceCollection
}
