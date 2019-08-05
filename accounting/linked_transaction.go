package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//LinkedTransaction can link transactions from an Accounts Payable invoice to an
//accounts recievable invoice - also known as a billable expense
type LinkedTransaction struct {

	// Filter by the SourceTransactionID. Get all the linked transactions created from a particular ACCPAY invoice
	SourceTransactionID string `json:"SourceTransactionID,omitempty" xml:"SourceTransactionID,omitempty"`

	// The line item identifier from the source transaction.
	SourceLineItemID string `json:"SourceLineItemID" xml:"SourceLineItemID"`

	// Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID=4bb34b03-3378-4bb2-a0ed-6345abf3224e&Status=APPROVED.
	ContactID string `json:"ContactID,omitempty" xml:"ContactID,omitempty"`

	// Filter by the TargetTransactionID. Get all the linked transactions allocated to a particular ACCREC invoice
	TargetTransactionID string `json:"TargetTransactionID,omitempty" xml:"TargetTransactionID,omitempty"`

	// The line item identifier from the target transaction. It is possible to link multiple billable expenses to the same TargetLineItemID.
	TargetLineItemID string `json:"TargetLineItemID,omitempty" xml:"TargetLineItemID,omitempty"`

	// The Xero identifier for an Linked Transaction e.g. /LinkedTransactions/297c2dc5-cc47-4afd-8ec8-74990b8761e9
	LinkedTransactionID string `json:"LinkedTransactionID,omitempty" xml:"LinkedTransactionID,omitempty"`

	// Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID=4bb34b03-3378-4bb2-a0ed-6345abf3224e&Status=APPROVED.
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// This will always be BILLABLEEXPENSE. More types may be added in future.
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// The last modified date in UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// The Type of the source tranasction. This will be ACCPAY if the linked transaction was created from an invoice and SPEND if it was created from a bank transaction.
	SourceTransactionTypeCode string `json:"SourceTransactionTypeCode,omitempty" xml:"SourceTransactionTypeCode,omitempty"`
}

//LinkedTransactions is a collection of LinkedTransactions
type LinkedTransactions struct {
	LinkedTransactions []LinkedTransaction `json:"LinkedTransactions" xml:"LinkedTransaction"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (l *LinkedTransactions) convertDates() error {
	var err error
	for n := len(l.LinkedTransactions) - 1; n >= 0; n-- {
		l.LinkedTransactions[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(l.LinkedTransactions[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalLinkedTransaction(linkedTransactionResponseBytes []byte) (*LinkedTransactions, error) {
	var linkedTransactionResponse *LinkedTransactions
	err := json.Unmarshal(linkedTransactionResponseBytes, &linkedTransactionResponse)
	if err != nil {
		return nil, err
	}

	err = linkedTransactionResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return linkedTransactionResponse, err
}

//Create will create LinkedTransactions given an LinkedTransactions struct
func (l *LinkedTransactions) Create(provider *xerogolang.Provider, session goth.Session) (*LinkedTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(l, "  ", "	")
	if err != nil {
		return nil, err
	}

	linkedTransactionResponseBytes, err := provider.Create(session, "LinkedTransactions", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalLinkedTransaction(linkedTransactionResponseBytes)
}

//Update will update an LinkedTransaction given an LinkedTransactions struct
//This will only handle single LinkedTransaction - you cannot update multiple LinkedTransactions in a single call
//LinkedTransactions cannot be modified, only created and deleted.
func (l *LinkedTransactions) Update(provider *xerogolang.Provider, session goth.Session) (*LinkedTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(l, "  ", "	")
	if err != nil {
		return nil, err
	}

	LinkedTransactionResponseBytes, err := provider.Update(session, "LinkedTransactions/"+l.LinkedTransactions[0].LinkedTransactionID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalLinkedTransaction(LinkedTransactionResponseBytes)
}

//FindLinkedTransactionsModifiedSince will get all LinkedTransactions modified after a specified date.
//additional querystringParameters such as page, SourceTransactionID, ContactID,
//Status, and TargetTransactionID can be added as a map
func FindLinkedTransactionsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*LinkedTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	linkedTransactionResponseBytes, err := provider.Find(session, "LinkedTransactions", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalLinkedTransaction(linkedTransactionResponseBytes)
}

//FindLinkedTransactions will get all LinkedTransactions.
//additional querystringParameters such as page, SourceTransactionID, ContactID,
//Status, and TargetTransactionID can be added as a map
func FindLinkedTransactions(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*LinkedTransactions, error) {
	return FindLinkedTransactionsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindLinkedTransaction will get a single LinkedTransaction - LinkedTransactionID must be a GUID for an LinkedTransaction
func FindLinkedTransaction(provider *xerogolang.Provider, session goth.Session, linkedTransactionID string) (*LinkedTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	linkedTransactionResponseBytes, err := provider.Find(session, "LinkedTransactions/"+linkedTransactionID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalLinkedTransaction(linkedTransactionResponseBytes)
}

//RemoveLinkedTransaction will get a single LinkedTransaction - LinkedTransactionID must be a GUID for an LinkedTransaction
func RemoveLinkedTransaction(provider *xerogolang.Provider, session goth.Session, linkedTransactionID string) (*LinkedTransactions, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	linkedTransactionResponseBytes, err := provider.Remove(session, "LinkedTransactions/"+linkedTransactionID, additionalHeaders)
	if err != nil {
		return nil, err
	}

	return unmarshalLinkedTransaction(linkedTransactionResponseBytes)
}

//GenerateExampleLinkedTransaction Creates an Example LinkedTransaction
func GenerateExampleLinkedTransaction(sourceTransactionID string, sourceLineItemID string, contactID string) *LinkedTransactions {
	linkedTransaction := LinkedTransaction{
		SourceTransactionID: sourceTransactionID,
		SourceLineItemID:    sourceLineItemID,
		ContactID:           contactID,
	}

	l := &LinkedTransactions{
		LinkedTransactions: []LinkedTransaction{},
	}

	l.LinkedTransactions = append(l.LinkedTransactions, linkedTransaction)

	return l
}
