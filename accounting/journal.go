package accounting

import (
	"encoding/json"
	"strconv"
	"time"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/helpers"
	"github.com/markbates/goth"
)

//Journal is a record of a financial transaction in Xero
type Journal struct {

	// Xero identifier
	JournalID string `json:"JournalID,omitempty" xml:"JournalID,omitempty"`

	// Date the journal was posted
	JournalDate string `json:"JournalDate,omitempty" xml:"JournalDate,omitempty"`

	// Xero generated journal number
	JournalNumber int `json:"JournalNumber,omitempty" xml:"JournalNumber,omitempty"`

	// Created date UTC format
	CreatedDateUTC string `json:"CreatedDateUTC,omitempty" xml:"CreatedDateUTC,omitempty"`

	//
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// The identifier for the source transaction (e.g. InvoiceID)
	SourceID string `json:"SourceID,omitempty" xml:"SourceID,omitempty"`

	// The journal source type. The type of transaction that created the journal
	SourceType string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`

	// See JournalLines
	JournalLines []JournalLine `json:"JournalLines,omitempty" xml:"JournalLines>JournalLine,omitempty"`
}

//Journals is a collection of Journals
type Journals struct {
	Journals []Journal `json:"Journals,omitempty" xml:"Journal,omitempty"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (j *Journals) convertJournalDates() error {
	var err error
	for n := len(j.Journals) - 1; n >= 0; n-- {
		j.Journals[n].JournalDate, err = helpers.DotNetJSONTimeToRFC3339(j.Journals[n].JournalDate, false)
		if err != nil {
			return err
		}
		j.Journals[n].CreatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(j.Journals[n].CreatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalJournals(journalResponseBytes []byte) (*Journals, error) {
	var journalResponse *Journals
	err := json.Unmarshal(journalResponseBytes, &journalResponse)
	if err != nil {
		return nil, err
	}

	err = journalResponse.convertJournalDates()
	if err != nil {
		return nil, err
	}

	return journalResponse, err
}

//FindJournalsModifiedSinceWithParams will get all journals modified after a specified date.
//A maximum of 100 journals will be returned in any response.
//Use the offset or ModifiedSince filters with multiple API calls to retrieve larger sets of journals.
//Journals are ordered oldest to newest.
//additional querystringParameters such as where, page, order can be added as a map
func FindJournalsModifiedSinceWithParams(provider *xero.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Journals, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	journalResponseBytes, err := provider.Find(session, "Journals", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalJournals(journalResponseBytes)
}

//FindJournalsModifiedSince will get all journals modified after a specified date.
//A maximum of 100 journals will be returned in any response.
//Use the offset or ModifiedSince filters with multiple API calls to retrieve larger sets of journals.
//Journals are ordered oldest to newest.
func FindJournalsModifiedSince(provider *xero.Provider, session goth.Session, modifiedSince time.Time, offset int) (*Journals, error) {
	querystringParameters := map[string]string{
		"offset": strconv.Itoa(offset),
	}
	return FindJournalsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindJournalsModifiedSincePaymentsOnly will get all journals modified after a specified date and are on a cash basis. Journals are returned on accrual basis by default
func FindJournalsModifiedSincePaymentsOnly(provider *xero.Provider, session goth.Session, modifiedSince time.Time, offset int) (*Journals, error) {
	querystringParameters := map[string]string{
		"paymentsOnly": "true",
		"offset":       strconv.Itoa(offset),
	}
	return FindJournalsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindJournalsPaymentsOnly will get all journals on a cash basis. Journals are returned on accrual basis by default
//A maximum of 100 journals will be returned in any response.
//Use the offset or ModifiedSince filters with multiple API calls to retrieve larger sets of journals.
//Journals are ordered oldest to newest.
func FindJournalsPaymentsOnly(provider *xero.Provider, session goth.Session, offset int) (*Journals, error) {
	querystringParameters := map[string]string{
		"paymentsOnly": "true",
		"offset":       strconv.Itoa(offset),
	}
	return FindJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindJournals will get all journals.
//A maximum of 100 journals will be returned in any response.
//Use the offset or ModifiedSince filters with multiple API calls to retrieve larger sets of journals.
//Journals are ordered oldest to newest.
func FindJournals(provider *xero.Provider, session goth.Session, offset int) (*Journals, error) {
	querystringParameters := map[string]string{
		"offset": strconv.Itoa(offset),
	}
	return FindJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindJournal will get a single journal - journalID can be a GUID for an journal or an journal number
func FindJournal(provider *xero.Provider, session goth.Session, journalID string) (*Journals, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	journalResponseBytes, err := provider.Find(session, "Journals/"+journalID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalJournals(journalResponseBytes)
}
