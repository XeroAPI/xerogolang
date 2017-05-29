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

//ManualJournal used by accountants or bookkeepers to work directly with the general ledger. For example, to record accrued expenses or completed work not invoiced.
type ManualJournal struct {

	// Description of journal being posted
	Narration string `json:"Narration" xml:"Narration"`

	// See JournalLines
	JournalLines []ManualJournalLine `json:"JournalLines" xml:"JournalLines>JournalLine"`

	// Date journal was posted – YYYY-MM-DD
	Date string `json:"Date,omitempty" xml:"Date,omitempty"`

	// NoTax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// See Manual Journal Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Url link to a source document – shown as “Go to [appName]” in the Xero app
	URL string `json:"Url,omitempty" xml:"Url,omitempty"`

	// Boolean – default is true if not specified
	ShowOnCashBasisReports bool `json:"ShowOnCashBasisReports,omitempty" xml:"ShowOnCashBasisReports,omitempty"`

	// Boolean to indicate if a manual journal has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"-"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// The Xero identifier for a Manual Journal
	ManualJournalID string `json:"ManualJournalID,omitempty" xml:"ManualJournalID,omitempty"`
}

//ManualJournals is a collection of ManualJournals
type ManualJournals struct {
	ManualJournals []ManualJournal `json:"ManualJournals,omitempty" xml:"ManualJournal,omitempty"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (m *ManualJournals) convertManualJournalDates() error {
	var err error
	for n := len(m.ManualJournals) - 1; n >= 0; n-- {
		m.ManualJournals[n].Date, err = helpers.DotNetJSONTimeToRFC3339(m.ManualJournals[n].Date, false)
		if err != nil {
			return err
		}
		m.ManualJournals[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(m.ManualJournals[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalManualJournal(manualJournalResponseBytes []byte) (*ManualJournals, error) {
	var manualJournalResponse *ManualJournals
	err := json.Unmarshal(manualJournalResponseBytes, &manualJournalResponse)
	if err != nil {
		return nil, err
	}

	err = manualJournalResponse.convertManualJournalDates()
	if err != nil {
		return nil, err
	}

	return manualJournalResponse, err
}

//CreateManualJournal will create manualJournals given an ManualJournals struct
func (m *ManualJournals) CreateManualJournal(provider *xero.Provider, session goth.Session) (*ManualJournals, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(m, "  ", "	")
	if err != nil {
		return nil, err
	}

	manualJournalResponseBytes, err := provider.Create(session, "ManualJournals", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalManualJournal(manualJournalResponseBytes)
}

//UpdateManualJournal will update an manualJournal given an ManualJournals struct
//This will only handle single manualJournal - you cannot update multiple manualJournals in a single call
func (m *ManualJournals) UpdateManualJournal(provider *xero.Provider, session goth.Session) (*ManualJournals, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(m, "  ", "	")
	if err != nil {
		return nil, err
	}

	manualJournalResponseBytes, err := provider.Update(session, "ManualJournals/"+m.ManualJournals[0].ManualJournalID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalManualJournal(manualJournalResponseBytes)
}

//FindManualJournalsModifiedSinceWithParams will get all ManualJournals modified after a specified date.
//These ManualJournals will not have details like default account codes and tracking categories.
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindManualJournalsModifiedSinceWithParams(provider *xero.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*ManualJournals, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	manualJournalResponseBytes, err := provider.Find(session, "ManualJournals", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalManualJournal(manualJournalResponseBytes)
}

//FindManualJournalsModifiedSince will get all ManualJournals modified after a specified date.
//These ManualJournals will not have details like default account codes and tracking categories.
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
func FindManualJournalsModifiedSince(provider *xero.Provider, session goth.Session, modifiedSince time.Time) (*ManualJournals, error) {
	return FindManualJournalsModifiedSinceWithParams(provider, session, modifiedSince, nil)
}

//FindManualJournalsModifiedSinceByPage will get a specified page of ManualJournals which contains 100 ManualJournals modified
//after a specified date. Page 1 gives the first 100, page two the next 100 etc etc.
//Paged ManualJournals contain all the detail of the ManualJournals whereas if you use FindAllManualJournals
//you will only get summarised data e.g. no line items or tracking categories
func FindManualJournalsModifiedSinceByPage(provider *xero.Provider, session goth.Session, modifiedSince time.Time, page int) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"page": strconv.Itoa(page),
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindManualJournalsModifiedSinceWhere will get ManualJournals which contains 100 ManualJournals
//that fit the criteria of a supplied where clause.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
func FindManualJournalsModifiedSinceWhere(provider *xero.Provider, session goth.Session, modifiedSince time.Time, whereClause string) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindManualJournalsModifiedSinceOrderedBy will get ManualJournals and are order them by a supplied named element.
//you will only get summarised data e.g. no line items or tracking categories
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
func FindManualJournalsModifiedSinceOrderedBy(provider *xero.Provider, session goth.Session, modifiedSince time.Time, orderBy string) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindManualJournalsByPage will get a specified page of ManualJournals which contains 100 ManualJournals
//Page 1 gives the first 100, page two the next 100 etc etc.
//paged ManualJournals contain all the detail of the ManualJournals whereas if you use FindAllManualJournals
//you will only get summarised data e.g. no line items or tracking categories
func FindManualJournalsByPage(provider *xero.Provider, session goth.Session, page int) (*ManualJournals, error) {
	return FindManualJournalsModifiedSinceByPage(provider, session, dayZero, page)
}

//FindManualJournalsByPageWhere will get a specified page of ManualJournals which contains 100 ManualJournals
//that fit the criteria of a supplied where clause. Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged ManualJournals contain all the detail of the ManualJournals whereas if you use FindAllManualJournals
//you will only get summarised data e.g. no line items or tracking categories
func FindManualJournalsByPageWhere(provider *xero.Provider, session goth.Session, page int, whereClause string) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindManualJournalsByPageWhereOrderedBy will get a specified page of ManualJournals which contains 100 ManualJournals
//that fit the criteria of a supplied where clause and are ordered by a supplied named element.
//Page 1 gives the first 100, page 2 the next 100 etc etc.
//paged ManualJournals contain all the detail of the ManualJournals whereas if you use FindManualJournals
//you will only get summarised data e.g. no line items or tracking categories
func FindManualJournalsByPageWhereOrderedBy(provider *xero.Provider, session goth.Session, page int, whereClause string, orderBy string) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"page":  strconv.Itoa(page),
		"where": whereClause,
		"order": orderBy,
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindManualJournalsOrderedBy will get all ManualJournals ordered by a supplied named element.
//These ManualJournals will not have details like line items.
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
func FindManualJournalsOrderedBy(provider *xero.Provider, session goth.Session, orderBy string) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"order": orderBy,
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindManualJournalsWhere will get all ManualJournals that fit the criteria of a supplied where clause.
//These ManualJournals will not have details like line items.
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
func FindManualJournalsWhere(provider *xero.Provider, session goth.Session, whereClause string) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindManualJournalsWhereOrderedBy will get all ManualJournals that fit the criteria of a supplied where clause
//and are ordered by a supplied named element. These ManualJournals will not have details like line items.
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
func FindManualJournalsWhereOrderedBy(provider *xero.Provider, session goth.Session, whereClause string, orderedBy string) (*ManualJournals, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
		"order": orderedBy,
	}
	return FindManualJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindManualJournalsWithParams will get all ManualJournals. These ManualJournal will not have details like line items.
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindManualJournalsWithParams(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*ManualJournals, error) {
	return FindManualJournalsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindManualJournals will get all ManualJournals. These ManualJournal will not have details like line items.
//If you need details then use FindManualJournalsByPage and get 100 ManualJournals at a time
func FindManualJournals(provider *xero.Provider, session goth.Session) (*ManualJournals, error) {
	return FindManualJournalsModifiedSinceWithParams(provider, session, dayZero, nil)
}

//FindManualJournal will get a single manualJournal - manualJournalID can be a GUID for an manualJournal or an manualJournal number
func FindManualJournal(provider *xero.Provider, session goth.Session, manualJournalID string) (*ManualJournals, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	manualJournalResponseBytes, err := provider.Find(session, "ManualJournals/"+manualJournalID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalManualJournal(manualJournalResponseBytes)
}

//CreateExampleManualJournal Creates an Example manualJournal
func CreateExampleManualJournal() *ManualJournals {
	lineItem := ManualJournalLine{
		Description: "Importing & Exporting Services",
		LineAmount:  395.00,
		AccountCode: "200",
	}

	lineItem2 := ManualJournalLine{
		Description: "Importing & Exporting Services",
		LineAmount:  -395.00,
		AccountCode: "310",
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	manualJournal := ManualJournal{
		Narration:       "Missed Importing & Exporting Invoice",
		Date:            strings.TrimSuffix(today.Format(time.RFC3339), "Z"),
		LineAmountTypes: "Exclusive",
		Status:          "DRAFT",
		JournalLines:    []ManualJournalLine{},
	}

	manualJournal.JournalLines = append(manualJournal.JournalLines, lineItem)

	manualJournal.JournalLines = append(manualJournal.JournalLines, lineItem2)

	manualJournalCollection := &ManualJournals{
		ManualJournals: []ManualJournal{},
	}

	manualJournalCollection.ManualJournals = append(manualJournalCollection.ManualJournals, manualJournal)

	return manualJournalCollection
}
