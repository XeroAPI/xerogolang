package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
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
	CurrencyRate float64 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// See Invoice Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Boolean to set whether the invoice in the Xero app should be marked as “sent”. This can be set only on invoices that have been approved
	SentToContact bool `json:"SentToContact,omitempty" xml:"SentToContact,omitempty"`

	// Shown on sales invoices (Accounts Receivable) when this has been set
	ExpectedPaymentDate string `json:"ExpectedPaymentDate,omitempty" xml:"ExpectedPaymentDate,omitempty"`

	// Shown on bills (Accounts Payable) when this has been set
	PlannedPaymentDate string `json:"PlannedPaymentDate,omitempty" xml:"PlannedPaymentDate,omitempty"`

	// Total of invoice excluding taxes
	SubTotal float64 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// Total tax on invoice
	TotalTax float64 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// Total of Invoice tax inclusive (i.e. SubTotal + TotalTax). This will be ignored if it doesn’t equal the sum of the LineAmounts
	Total float64 `json:"Total,omitempty" xml:"Total,omitempty"`

	// Total of discounts applied on the invoice line items
	TotalDiscount float64 `json:"TotalDiscount,omitempty" xml:"-"`

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
	AmountDue float64 `json:"AmountDue,omitempty" xml:"-"`

	// Sum of payments received for invoice
	AmountPaid float64 `json:"AmountPaid,omitempty" xml:"-"`

	// The date the invoice was fully paid. Only returned on fully paid invoices
	FullyPaidOnDate string `json:"FullyPaidOnDate,omitempty" xml:"-"`

	// Sum of all credit notes, over-payments and pre-payments applied to invoice
	AmountCredited float64 `json:"AmountCredited,omitempty" xml:"-"`

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
func (i *Invoices) convertDates() error {
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

	err = invoiceResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return invoiceResponse, err
}

//Create will create invoices given an Invoices struct
func (i *Invoices) Create(provider xerogolang.IProvider, session goth.Session) (*Invoices, error) {
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

//Update will update an invoice given an Invoices struct
//This will only handle single invoice - you cannot update multiple invoices in a single call
func (i *Invoices) Update(provider xerogolang.IProvider, session goth.Session) (*Invoices, error) {
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

//FindInvoicesModifiedSince will get all Invoices modified after a specified date.
//These Invoices will not have details like default line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Invoices at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindInvoicesModifiedSince(provider xerogolang.IProvider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Invoices, error) {
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

//FindInvoices will get all Invoices. These Invoice will not have details like line items by default.
//If you need details then add a 'page' querystringParameter and get 100 Invoices at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindInvoices(provider xerogolang.IProvider, session goth.Session, querystringParameters map[string]string) (*Invoices, error) {
	return FindInvoicesModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindInvoice will get a single invoice - invoiceID can be a GUID for an invoice or an invoice number
func FindInvoice(provider xerogolang.IProvider, session goth.Session, invoiceID string) (*Invoices, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	invoiceResponseBytes, err := provider.Find(session, "Invoices/"+invoiceID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalInvoice(invoiceResponseBytes)
}

//GenerateExampleInvoice Creates an Example invoice
func GenerateExampleInvoice() *Invoices {
	lineItem := LineItem{
		Description: "Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	invoice := Invoice{
		Type: "ACCREC",
		Contact: Contact{
			Name: "George Costanza",
		},
		Date:            helpers.TodayRFC3339(),
		DueDate:         helpers.FormatDate(time.Now().Add(720 * time.Hour)),
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
