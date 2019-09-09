package accounting

import (
	"encoding/json"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//BankTransfer is a record of monies transferred from one bank account to another
type HistoryRecord struct {

  // The type of change recorded against the document
  Changes string `json:"Changes,omitempty" xml:"-"`

  // UTC date that the history record was created
	DateUTC string `json:"DateUTC,omitempty" xml:"-"`

	// The user responsible for the change ("System Generated" when the change happens via API)
	User string `json:"User,omitempty" xml:"-"`

	// The Bank Transaction ID for the source account
	Details string `json:"Details" xml:"Details"`
}

//BankTransfers contains a collection of BankTransfers
type HistoryRecords struct {
	HistoryRecords []HistoryRecord `json:"HistoryRecords" xml:"HistoryRecords"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (h *HistoryRecords) convertDates() error {
	var err error
	for n := len(h.HistoryRecords) - 1; n >= 0; n-- {
		h.HistoryRecords[n].DateUTC, err = helpers.DotNetJSONTimeToRFC3339(h.HistoryRecords[n].DateUTC, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func unmarshalHistoryRecord(HistoryRecordResponseBytes []byte) (*HistoryRecords, error) {
	var historyRecordResponse *HistoryRecords
	err := json.Unmarshal(HistoryRecordResponseBytes, &historyRecordResponse)
	if err != nil {
		return nil, err
	}

	err = historyRecordResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return historyRecordResponse, err
}

//Create will create History Records given a HistoryRecords struct and a docType and id
func (h *HistoryRecords) Create(provider xerogolang.IProvider, session goth.Session, docType string, id string) (*HistoryRecords, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	body, err := json.MarshalIndent(h, "  ", "	")
	if err != nil {
		return nil, err
	}

	historyRecordResponseBytes, err := provider.Create(session, docType + "/" + id + "/history", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalHistoryRecord(historyRecordResponseBytes)
}

//FindHistoryAndNotes gets all history items and notes for a given type and ID.
//it is not supported on all endpoints.  See https://developer.xero.com/documentation/api/history-and-notes#SupportedDocs
func FindHistoryAndNotes(provider xerogolang.IProvider, session goth.Session, docType string, id string) (*HistoryRecords, error) {
  additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	historyRecordResponseBytes, err := provider.Find(session, docType + "/" + id + "/history", additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalHistoryRecord(historyRecordResponseBytes)
}

//GenerateExampleBankTransfer Creates an Example bankTransfer
func GenerateExampleHistoryRecord() *HistoryRecords {
	historyRecord := HistoryRecord{
		Details: "Nothing happened",
	}

	historyRecordCollection := &HistoryRecords{
		HistoryRecords: []HistoryRecord{},
	}

	historyRecordCollection.HistoryRecords = append(historyRecordCollection.HistoryRecords, historyRecord)

	return historyRecordCollection
}
