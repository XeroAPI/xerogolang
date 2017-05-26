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
func (b *BankTransactions) convertBankTransactionDates() error {
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

	err = bankTransactionResponse.convertBankTransactionDates()
	if err != nil {
		return nil, err
	}

	return bankTransactionResponse, err
}

//CreateBankTransaction will create BankTransactions given an BankTransactions struct
func (b *BankTransactions) CreateBankTransaction(provider *xero.Provider, session goth.Session) (*BankTransactions, error) {
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

//UpdateBankTransaction will update a BankTransaction given a BankTransactions struct
//This will only handle single BankTransaction - you cannot update multiple BankTransactions in a single call
func (b *BankTransactions) UpdateBankTransaction(provider *xero.Provider, session goth.Session) (*BankTransactions, error) {
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

//FindBankTransactionsModifiedSinceWithParams will get all BankTransactions modified after a specified date.
//These BankTransactions will not have details like default account codes and tracking categories.
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindBankTransactionsModifiedSinceWithParams(provider *xero.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*BankTransactions, error) {
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

//FindBankTransactionsModifiedSince will get all BankTransactions modified after a specified date.
//These BankTransactions will not have details like default account codes and tracking categories.
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
func FindBankTransactionsModifiedSince(provider *xero.Provider, session goth.Session, modifiedSince time.Time) (*BankTransactions, error) {
	return FindBankTransactionsModifiedSinceWithParams(provider, session, modifiedSince, nil)
}

//FindBankTransactionsModifiedSinceByPage will get a specified page of BankTransactions which contains 100 BankTransactions modified
//after a specified date. Page 1 gives the first 100, page two the next 100 etc etc.
//Paged BankTransactions contain all the detail of the BankTransactions whereas if you use FindAllBankTransactions
//you will only get summarised data e.g. no line items or tracking categories
func FindBankTransactionsModifiedSinceByPage(provider *xero.Provider, session goth.Session, modifiedSince time.Time, page int) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"page": strconv.Itoa(page),
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindBankTransactionsModifiedSinceWhere will get BankTransactions which contains 100 BankTransactions
//that fit the criteria of a supplied where clause.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
func FindBankTransactionsModifiedSinceWhere(provider *xero.Provider, session goth.Session, modifiedSince time.Time, whereClause string) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindBankTransactionsModifiedSinceOrderedBy will get BankTransactions and are order them by a supplied named element.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
func FindBankTransactionsModifiedSinceOrderedBy(provider *xero.Provider, session goth.Session, modifiedSince time.Time, orderBy string) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindBankTransactionsByPage will get a specified page of BankTransactions which contains 100 BankTransactions
//Page 1 gives the first 100, page two the next 100 etc etc.
//paged BankTransactions contain all the detail of the BankTransactions whereas if you use FindAllBankTransactions
//you will only get summarised data e.g. no line items or tracking categories
func FindBankTransactionsByPage(provider *xero.Provider, session goth.Session, page int) (*BankTransactions, error) {
	return FindBankTransactionsModifiedSinceByPage(provider, session, dayZero, page)
}

//FindBankTransactionsByPageWhere will get a specified page of BankTransactions which contains 100 BankTransactions
//that fit the criteria of a supplied where clause. Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged BankTransactions contain all the detail of the BankTransactions whereas if you use FindAllBankTransactions
//you will only get summarised data e.g. no line items or tracking categories
func FindBankTransactionsByPageWhere(provider *xero.Provider, session goth.Session, page int, whereClause string) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindBankTransactionsByPageWhereOrderedBy will get a specified page of BankTransactions which contains 100 BankTransactions
//that fit the criteria of a supplied where clause and are ordered by a supplied named element.
//Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged BankTransactions contain all the detail of the BankTransactions whereas if you use FindBankTransactions
//you will only get summarised data e.g. no line items or tracking categories
func FindBankTransactionsByPageWhereOrderedBy(provider *xero.Provider, session goth.Session, page int, whereClause string, orderBy string) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
		"order": orderBy,
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindBankTransactionsOrderedBy will get all BankTransactions ordered by a supplied named element.
//These BankTransactions will not have details like line items.
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
func FindBankTransactionsOrderedBy(provider *xero.Provider, session goth.Session, orderBy string) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindBankTransactionsWhere will get all BankTransactions that fit the criteria of a supplied where clause.
//These BankTransactions will not have details like line items.
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
func FindBankTransactionsWhere(provider *xero.Provider, session goth.Session, whereClause string) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindBankTransactionsWhereOrderedBy will get all BankTransactions that fit the criteria of a supplied where clause
//and are ordered by a supplied named element. These BankTransactions will not have details like line items.
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
func FindBankTransactionsWhereOrderedBy(provider *xero.Provider, session goth.Session, whereClause string, orderedBy string) (*BankTransactions, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
		"order": orderedBy,
	}
	return FindBankTransactionsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindBankTransactionsWithParams will get all BankTransactions. These BankTransaction will not have details like line items.
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindBankTransactionsWithParams(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*BankTransactions, error) {
	return FindBankTransactionsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindBankTransactions will get all BankTransactions. These BankTransaction will not have details like line items.
//If you need details then use FindBankTransactionsByPage and get 100 BankTransactions at a time
func FindBankTransactions(provider *xero.Provider, session goth.Session) (*BankTransactions, error) {
	return FindBankTransactionsModifiedSinceWithParams(provider, session, dayZero, nil)
}

//FindBankTransaction will get a single BankTransaction - BankTransactionID can be a GUID for an BankTransaction or an BankTransaction number
func FindBankTransaction(provider *xero.Provider, session goth.Session, bankTransactionID string) (*BankTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	bankTransactionResponseBytes, err := provider.Find(session, "BankTransactions/"+bankTransactionID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransaction(bankTransactionResponseBytes)
}

//CreateExampleBankTransaction Creates an Example bankTransaction
func CreateExampleBankTransaction() *BankTransactions {
	lineItem := LineItem{
		Description: "Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	bankAccount := BankAccount{
		Code: "090",
	}

	bankTransaction := BankTransaction{
		Type: "RECEIVE",
		Contact: Contact{
			Name: "George Costanza",
		},
		Date:        strings.TrimSuffix(today.Format(time.RFC3339), "Z"),
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
