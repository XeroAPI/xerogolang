package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//ExpenseClaim is a collection of receipts from personal spending that needs to be reimbursed by the business
type ExpenseClaim struct {
	//See Users
	User User `json:"User" xml:"User"`

	// See Receipts
	Receipts []Receipt `json:"Receipts" xml:"Receipts>Receipt"`

	// Xero generated unique identifier for an expense claim
	ExpenseClaimID string `json:"ExpenseClaimID,omitempty" xml:"ExpenseClaimID,omitempty"`

	// See Payments
	Payments *[]Payment `json:"Payments,omitempty" xml:"Payments>Payment,omitempty"`

	// Current status of an expense claim – see status types
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// The total of an expense claim being paid
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// The amount due to be paid for an expense claim
	AmountDue float32 `json:"AmountDue,omitempty" xml:"AmountDue,omitempty"`

	// The amount still to pay for an expense claim
	AmountPaid float32 `json:"AmountPaid,omitempty" xml:"AmountPaid,omitempty"`

	// The date when the expense claim is due to be paid YYYY-MM-DD
	PaymentDueDate string `json:"PaymentDueDate,omitempty" xml:"PaymentDueDate,omitempty"`

	// The date the expense claim will be reported in Xero YYYY-MM-DD
	ReportingDate string `json:"ReportingDate,omitempty" xml:"ReportingDate,omitempty"`

	// The Xero identifier for the Receipt e.g. e59a2c7f-1306-4078-a0f3-73537afcbba9
	ReceiptID string `json:"ReceiptID" xml:"ReceiptID"`
}

//ExpenseClaims contains a collection of ExpenseClaims
type ExpenseClaims struct {
	ExpenseClaims []ExpenseClaim `json:"ExpenseClaims" xml:"ExpenseClaim"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (e *ExpenseClaims) convertDates() error {
	var err error
	for n := len(e.ExpenseClaims) - 1; n >= 0; n-- {
		e.ExpenseClaims[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(e.ExpenseClaims[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalExpenseClaim(expenseClaimResponseBytes []byte) (*ExpenseClaims, error) {
	var expenseClaimResponse *ExpenseClaims
	err := json.Unmarshal(expenseClaimResponseBytes, &expenseClaimResponse)
	if err != nil {
		return nil, err
	}

	err = expenseClaimResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return expenseClaimResponse, err
}

//Create will create expenseClaims given an ExpenseClaims struct
func (e *ExpenseClaims) Create(provider *xerogolang.Provider, session goth.Session) (*ExpenseClaims, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(e, "  ", "	")
	if err != nil {
		return nil, err
	}

	expenseClaimResponseBytes, err := provider.Create(session, "ExpenseClaims", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalExpenseClaim(expenseClaimResponseBytes)
}

//Update will update an expenseClaim given an ExpenseClaims struct
//This will only handle single expenseClaim - you cannot update multiple expenseClaims in a single call
func (e *ExpenseClaims) Update(provider *xerogolang.Provider, session goth.Session) (*ExpenseClaims, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	//strip out user details because we only need to send an ID
	u := User{
		UserID: e.ExpenseClaims[0].User.UserID,
	}
	e.ExpenseClaims[0].User = u

	body, err := xml.MarshalIndent(e, "  ", "	")
	if err != nil {
		return nil, err
	}

	expenseClaimResponseBytes, err := provider.Update(session, "ExpenseClaims/"+e.ExpenseClaims[0].ExpenseClaimID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalExpenseClaim(expenseClaimResponseBytes)
}

//FindExpenseClaimsModifiedSince will get all ExpenseClaims modified after a specified date.
//These ExpenseClaims will not have details like default line items by default.
//If you need details then add a 'page' querystringParameter and get 100 ExpenseClaims at a time
//additional querystringParameters such as where and order can be added as a map
func FindExpenseClaimsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*ExpenseClaims, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	expenseClaimResponseBytes, err := provider.Find(session, "ExpenseClaims", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalExpenseClaim(expenseClaimResponseBytes)
}

//FindExpenseClaims will get all ExpenseClaims. These ExpenseClaim will not have details like line items by default.
//If you need details then add a 'page' querystringParameter and get 100 ExpenseClaims at a time
//additional querystringParameters such as where and order can be added as a map
func FindExpenseClaims(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*ExpenseClaims, error) {
	return FindExpenseClaimsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindExpenseClaim will get a single expenseClaim - expenseClaimID can be a GUID for an expenseClaim or an expenseClaim number
func FindExpenseClaim(provider *xerogolang.Provider, session goth.Session, expenseClaimID string) (*ExpenseClaims, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	expenseClaimResponseBytes, err := provider.Find(session, "ExpenseClaims/"+expenseClaimID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalExpenseClaim(expenseClaimResponseBytes)
}

//GenerateExampleExpenseClaim Creates an Example expenseClaim
func GenerateExampleExpenseClaim(userID string, receiptID string) *ExpenseClaims {
	receipt := Receipt{
		ReceiptID: receiptID,
	}

	user := User{
		UserID: userID,
	}

	expenseClaim := ExpenseClaim{
		User:     user,
		Receipts: []Receipt{},
	}

	expenseClaim.Receipts = append(expenseClaim.Receipts, receipt)

	expenseClaimCollection := &ExpenseClaims{
		ExpenseClaims: []ExpenseClaim{},
	}

	expenseClaimCollection.ExpenseClaims = append(expenseClaimCollection.ExpenseClaims, expenseClaim)

	return expenseClaimCollection
}
