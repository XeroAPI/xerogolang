package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//Account represents individual accounts in a Xero organisation
type Account struct {

	// Customer defined alpha numeric account code e.g 200 or SALES (max length = 10)
	Code string `json:"Code,omitempty" xml:"Code,omitempty"`

	// Name of account (max length = 150)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// See Account Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// For bank accounts only (Account Type BANK)
	BankAccountNumber string `json:"BankAccountNumber,omitempty" xml:"BankAccountNumber,omitempty"`

	// Accounts with a status of ACTIVE can be updated to ARCHIVED. See Account Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Description of the Account. Valid for all types of accounts except bank accounts (max length = 4000)
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// For bank accounts only. See Bank Account types
	BankAccountType string `json:"BankAccountType,omitempty" xml:"BankAccountType,omitempty"`

	// For bank accounts only
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// See Tax Types
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`

	// Boolean – describes whether account can have payments applied to it
	EnablePaymentsToAccount bool `json:"EnablePaymentsToAccount,omitempty" xml:"EnablePaymentsToAccount,omitempty"`

	// Boolean – describes whether account code is available for use with expense claims
	ShowInExpenseClaims bool `json:"ShowInExpenseClaims,omitempty" xml:"ShowInExpenseClaims,omitempty"`

	// The Xero identifier for an account – specified as a string following the endpoint name e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9
	AccountID string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`

	// See Account Class Types
	Class string `json:"Class,omitempty" xml:"-"`

	// If this is a system account then this element is returned. See System Account types. Note that non-system accounts may have this element set as either “” or null.
	SystemAccount string `json:"SystemAccount,omitempty" xml:"-"`

	// Shown if set
	ReportingCode string `json:"ReportingCode,omitempty" xml:"-"`

	// Shown if set
	ReportingCodeName string `json:"ReportingCodeName,omitempty" xml:"-"`

	// boolean to indicate if an account has an attachment (read only)
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"-"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`
}

//Accounts contains a collection of Accounts
type Accounts struct {
	Accounts []Account `json:"Accounts,omitempty" xml:"Account,omitempty"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (a *Accounts) convertDates() error {
	var err error
	for n := len(a.Accounts) - 1; n >= 0; n-- {
		a.Accounts[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(a.Accounts[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalAccount(accountResponseBytes []byte) (*Accounts, error) {
	var accountResponse *Accounts
	err := json.Unmarshal(accountResponseBytes, &accountResponse)
	if err != nil {
		return nil, err
	}

	err = accountResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return accountResponse, err
}

//Create will create accounts given an Accounts struct
func (a *Accounts) Create(provider *xerogolang.Provider, session goth.Session) (*Accounts, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(a, "  ", "	")
	if err != nil {
		return nil, err
	}

	accountResponseBytes, err := provider.Create(session, "Accounts", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalAccount(accountResponseBytes)
}

//Update will update an account given an Accounts struct
//This will only handle single account - you cannot update multiple accounts in a single call
func (a *Accounts) Update(provider *xerogolang.Provider, session goth.Session) (*Accounts, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(a, "  ", "	")
	if err != nil {
		return nil, err
	}

	accountResponseBytes, err := provider.Update(session, "Accounts/"+a.Accounts[0].AccountID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalAccount(accountResponseBytes)
}

//FindAccountsModifiedSince will get all accounts modified after a specified date.
//additional querystringParameters such as where and order can be added as a map
func FindAccountsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Accounts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	accountResponseBytes, err := provider.Find(session, "Accounts", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalAccount(accountResponseBytes)
}

//FindAccounts will get all accounts. These account will not have details like line items.
//additional querystringParameters such as where and order can be added as a map
func FindAccounts(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*Accounts, error) {
	return FindAccountsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindAccount will get a single account - accountID must be a GUID for an account
func FindAccount(provider *xerogolang.Provider, session goth.Session, accountID string) (*Accounts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	accountResponseBytes, err := provider.Find(session, "Accounts/"+accountID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalAccount(accountResponseBytes)
}

//RemoveAccount will get a single account - accountID must be a GUID for an account
func RemoveAccount(provider *xerogolang.Provider, session goth.Session, accountID string) (*Accounts, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	accountResponseBytes, err := provider.Remove(session, "Accounts/"+accountID, additionalHeaders)
	if err != nil {
		return nil, err
	}

	return unmarshalAccount(accountResponseBytes)
}

//GenerateExampleAccount Creates an Example account
func GenerateExampleAccount() *Accounts {
	account := Account{
		Code:                    "9999",
		Name:                    "Import/Exports",
		Type:                    "SALES",
		Status:                  "ACTIVE",
		Description:             "Proceeds from importing/exporting latex",
		TaxType:                 "OUTPUT2",
		EnablePaymentsToAccount: false,
		ShowInExpenseClaims:     false,
	}

	accountCollection := &Accounts{
		Accounts: []Account{},
	}

	accountCollection.Accounts = append(accountCollection.Accounts, account)

	return accountCollection
}
