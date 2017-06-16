package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//BankTransaction is a bank transaction
type BankTransaction struct {

	// See Bank Transaction Types
	Type string `json:"Type" xml:"Type"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// See LineItems
	LineItems []LineItem `json:"LineItems" xml:"LineItems>LineItem"`

	// Boolean to show if transaction is reconciled
	IsReconciled bool `json:"IsReconciled,omitempty" xml:"IsReconciled,omitempty"`

	// Date of transaction – YYYY-MM-DD
	Date string `json:"DateString,omitempty" xml:"Date,omitempty"`

	// Reference for the transaction. Only supported for SPEND and RECEIVE transactions.
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// The currency that bank transaction has been raised in (see Currencies). Setting currency is only supported on overpayments.
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// Exchange rate to base currency when money is spent or received. e.g. 0.7500 Only used for bank transactions in non base currency. If this isn’t specified for non base currency accounts then either the user-defined rate (preference) or the XE.com day rate will be used. Setting currency is only supported on overpayments.
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// URL link to a source document – shown as “Go to App Name”
	URL string `json:"Url,omitempty" xml:"Url,omitempty"`

	// See Bank Transaction Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// Total of bank transaction excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// Total tax on bank transaction
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// Total of bank transaction tax inclusive
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// Xero generated unique identifier for bank transaction
	BankTransactionID string `json:"BankTransactionID,omitempty" xml:"BankTransactionID,omitempty"`

	// Xero Bank Account
	BankAccount BankAccount `json:"BankAccount,omitempty" xml:"BankAccount,omitempty"`

	// Xero generated unique identifier for a Prepayment. This will be returned on BankTransactions with a Type of SPEND-PREPAYMENT or RECEIVE-PREPAYMENT
	PrepaymentID string `json:"PrepaymentID,omitempty" xml:"-"`

	// Xero generated unique identifier for an Overpayment. This will be returned on BankTransactions with a Type of SPEND-OVERPAYMENT or RECEIVE-OVERPAYMENT
	OverpaymentID string `json:"OverpaymentID,omitempty" xml:"-"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// Boolean to indicate if a bank transaction has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"-"`
}

//BankTransactions contains a collection of BankTransactions
type BankTransactions struct {
	BankTransactions []BankTransaction `json:"BankTransactions" xml:"BankTransaction"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (b *BankTransactions) convertDates() error {
	var err error
	for n := len(b.BankTransactions) - 1; n >= 0; n-- {
		b.BankTransactions[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(b.BankTransactions[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalBankTransaction(bankTransactionResponseBytes []byte) (*BankTransactions, error) {
	var bankTransactionResponse *BankTransactions
	err := json.Unmarshal(bankTransactionResponseBytes, &bankTransactionResponse)
	if err != nil {
		return nil, err
	}

	err = bankTransactionResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return bankTransactionResponse, err
}

//Create will create BankTransactions given an BankTransactions struct
func (b *BankTransactions) Create(provider *xerogolang.Provider, session goth.Session) (*BankTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(b, "  ", "	")
	if err != nil {
		return nil, err
	}

	bankTransactionResponseBytes, err := provider.Create(session, "BankTransactions", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransaction(bankTransactionResponseBytes)
}

//Update will update a BankTransaction given a BankTransactions struct
//This will only handle single BankTransaction - you cannot update multiple BankTransactions in a single call
func (b *BankTransactions) Update(provider *xerogolang.Provider, session goth.Session) (*BankTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(b, "  ", "	")
	if err != nil {
		return nil, err
	}

	bankTransactionResponseBytes, err := provider.Update(session, "BankTransactions/"+b.BankTransactions[0].BankTransactionID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransaction(bankTransactionResponseBytes)
}

//FindBankTransactionsModifiedSince will get all BankTransactions modified after a specified date.
//These BankTransactions will not have details like default account codes and tracking categories by default.
//If you need details then then add a 'page' querystringParameter and get 100 BankTransactions at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindBankTransactionsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*BankTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	bankTransactionResponseBytes, err := provider.Find(session, "BankTransactions", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransaction(bankTransactionResponseBytes)
}

//FindBankTransactions will get all BankTransactions. These BankTransaction will not have details like line items by default.
//If you need details then then add a 'page' querystringParameter and get 100 BankTransactions at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindBankTransactions(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*BankTransactions, error) {
	return FindBankTransactionsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindBankTransaction will get a single BankTransaction - BankTransactionID can be a GUID for an BankTransaction or an BankTransaction number
func FindBankTransaction(provider *xerogolang.Provider, session goth.Session, bankTransactionID string) (*BankTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	bankTransactionResponseBytes, err := provider.Find(session, "BankTransactions/"+bankTransactionID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransaction(bankTransactionResponseBytes)
}

//GenerateExampleBankTransaction Creates an Example bankTransaction
func GenerateExampleBankTransaction() *BankTransactions {
	lineItem := LineItem{
		Description: "Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	bankAccount := BankAccount{
		Code: "090",
	}

	bankTransaction := BankTransaction{
		Type: "RECEIVE",
		Contact: Contact{
			Name: "George Costanza",
		},
		Date:        helpers.TodayRFC3339(),
		LineItems:   []LineItem{},
		BankAccount: bankAccount,
	}

	bankTransaction.LineItems = append(bankTransaction.LineItems, lineItem)

	bankTransactionCollection := &BankTransactions{
		BankTransactions: []BankTransaction{},
	}

	bankTransactionCollection.BankTransactions = append(bankTransactionCollection.BankTransactions, bankTransaction)

	return bankTransactionCollection
}
