package accounting

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"time"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/helpers"
	"github.com/markbates/goth"
)

//Payment details payments against invoices and CreditNotes
type Payment struct {

	// Number of invoice or credit note you are applying payment to e.g. INV-4003
	Invoice *Invoice `json:"Invoice,omitempty" xml:"Invoice,omitempty"`

	// Number of invoice or credit note you are applying payment to e.g. INV-4003
	CreditNote *CreditNote `json:"CreditNote,omitempty" xml:"CreditNote,omitempty"`

	//Account of payment
	Account *Account `json:"Account,omitempty" xml:"Account,omitempty"`

	// Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06
	Date string `json:"Date,omitempty" xml:"Date,omitempty"`

	// Exchange rate when payment is received. Only used for non base currency invoices and credit notes e.g. 0.7500
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The amount of the payment. Must be less than or equal to the outstanding amount owing on the invoice e.g. 200.00
	Amount float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`

	// An optional description for the payment e.g. Direct Debit
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// An optional parameter for the payment. A boolean indicating whether you would like the payment to be created as reconciled when using PUT, or whether a payment has been reconciled when using GET
	IsReconciled bool `json:"IsReconciled,omitempty" xml:"IsReconciled,omitempty"`

	// The status of the payment.
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See Payment Types.
	PaymentType string `json:"PaymentType,omitempty" xml:"-"`

	// UTC timestamp of last update to the payment
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// The Xero identifier for an Payment e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	PaymentID string `json:"PaymentID,omitempty" xml:"PaymentID,omitempty"`
}

//Payments is a collection of Payments
type Payments struct {
	Payments []Payment `json:"Payments" xml:"Payment"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (p *Payments) convertPaymentDates() error {
	var err error
	for n := len(p.Payments) - 1; n >= 0; n-- {
		p.Payments[n].Date, err = helpers.DotNetJSONTimeToRFC3339(p.Payments[n].Date, false)
		if err != nil {
			return err
		}
		p.Payments[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(p.Payments[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalPayment(paymentResponseBytes []byte) (*Payments, error) {
	var paymentResponse *Payments
	err := json.Unmarshal(paymentResponseBytes, &paymentResponse)
	if err != nil {
		return nil, err
	}

	err = paymentResponse.convertPaymentDates()
	if err != nil {
		return nil, err
	}

	return paymentResponse, err
}

//CreatePayment will create payments given an Payments struct
func (p *Payments) CreatePayment(provider *xero.Provider, session goth.Session) (*Payments, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(p, "  ", "	")
	if err != nil {
		return nil, err
	}

	paymentResponseBytes, err := provider.Create(session, "Payments", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalPayment(paymentResponseBytes)
}

//UpdatePayment will update an payment given an Payments struct
//This will only handle single payment - you cannot update multiple payments in a single call
//Payments cannot be modified, only created and deleted.
func (p *Payments) UpdatePayment(provider *xero.Provider, session goth.Session) (*Payments, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	//we can only update the status on a payment so we must strip out all the other values in order to update it
	paymentToMarshal := Payment{
		Status: p.Payments[0].Status,
	}

	body, err := xml.MarshalIndent(paymentToMarshal, "  ", "	")
	if err != nil {
		return nil, err
	}

	paymentResponseBytes, err := provider.Update(session, "Payments/"+p.Payments[0].PaymentID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalPayment(paymentResponseBytes)
}

//FindPaymentsModifiedSinceWithParams will get all payments modified after a specified date.
//additional querystringParameters such as where, page, order can be added as a map
func FindPaymentsModifiedSinceWithParams(provider *xero.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Payments, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	paymentResponseBytes, err := provider.Find(session, "Payments", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalPayment(paymentResponseBytes)
}

//FindPaymentsModifiedSince will get all payments modified after a specified date.
func FindPaymentsModifiedSince(provider *xero.Provider, session goth.Session, modifiedSince time.Time) (*Payments, error) {
	return FindPaymentsModifiedSinceWithParams(provider, session, modifiedSince, nil)
}

//FindPaymentsModifiedSinceWhere will get all payments modified after a specified date that fit the criteria of a supplied where clause.
func FindPaymentsModifiedSinceWhere(provider *xero.Provider, session goth.Session, modifiedSince time.Time, whereClause string) (*Payments, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindPaymentsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindPaymentsWhere will get payments that fit the criteria of a supplied where clause.
func FindPaymentsWhere(provider *xero.Provider, session goth.Session, whereClause string) (*Payments, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindPaymentsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindPayments will get all payments.
func FindPayments(provider *xero.Provider, session goth.Session) (*Payments, error) {
	return FindPaymentsModifiedSinceWithParams(provider, session, dayZero, nil)
}

//FindPayment will get a single payment - paymentID must be a GUID for an payment
func FindPayment(provider *xero.Provider, session goth.Session, paymentID string) (*Payments, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	paymentResponseBytes, err := provider.Find(session, "Payments/"+paymentID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalPayment(paymentResponseBytes)
}

//RemovePayment will get a single payment - paymentID must be a GUID for an payment
func RemovePayment(provider *xero.Provider, session goth.Session, paymentID string) (*Payments, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	paymentResponseBytes, err := provider.Remove(session, "Payments/"+paymentID, additionalHeaders)
	if err != nil {
		return nil, err
	}

	return unmarshalPayment(paymentResponseBytes)
}

//CreateExamplePayment Creates an Example payment
func CreateExamplePayment(invoiceID string, amount float32) *Payments {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	payment := Payment{
		Date:   strings.TrimSuffix(today.Format(time.RFC3339), "Z"),
		Amount: amount,
		Invoice: &Invoice{
			InvoiceID: invoiceID,
		},
		Account: &Account{
			Code: "200",
		},
	}

	paymentCollection := &Payments{
		Payments: []Payment{},
	}

	paymentCollection.Payments = append(paymentCollection.Payments, payment)

	return paymentCollection
}
