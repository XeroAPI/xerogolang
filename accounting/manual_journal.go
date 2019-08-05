package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
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
func (m *ManualJournals) convertDates() error {
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

	err = manualJournalResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return manualJournalResponse, err
}

//Create will create manualJournals given an ManualJournals struct
func (m *ManualJournals) Create(provider *xerogolang.Provider, session goth.Session) (*ManualJournals, error) {
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

//Update will update an manualJournal given an ManualJournals struct
//This will only handle single manualJournal - you cannot update multiple manualJournals in a single call
func (m *ManualJournals) Update(provider *xerogolang.Provider, session goth.Session) (*ManualJournals, error) {
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

//FindManualJournalsModifiedSince will get all ManualJournals modified after a specified date.
//These ManualJournals will not have details like line items by default
//If you need details then then add a 'page' querystringParameter and get 100 ManualJournals at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindManualJournalsModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*ManualJournals, error) {
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

//FindManualJournals will get all ManualJournals. These ManualJournal will not have details like line items.
//If you need details then then add a 'page' querystringParameter and get 100 ManualJournals at a time
//additional querystringParameters such as where, page, order can be added as a map
func FindManualJournals(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*ManualJournals, error) {
	return FindManualJournalsModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindManualJournal will get a single manualJournal - manualJournalID can be a GUID for an manualJournal or an manualJournal number
func FindManualJournal(provider *xerogolang.Provider, session goth.Session, manualJournalID string) (*ManualJournals, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	manualJournalResponseBytes, err := provider.Find(session, "ManualJournals/"+manualJournalID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalManualJournal(manualJournalResponseBytes)
}

//GenerateExampleManualJournal Creates an Example manualJournal
func GenerateExampleManualJournal() *ManualJournals {
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

	manualJournal := ManualJournal{
		Narration:       "Missed Importing & Exporting Invoice",
		Date:            helpers.TodayRFC3339(),
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
