package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//BankTransfer is a record of monies transferred from one bank account to another
type BankTransfer struct {

	//
	Amount float32 `json:"Amount" xml:"Amount"`

	// The date of the Transfer YYYY-MM-DD
	Date string `json:"Date,omitempty" xml:"Date,omitempty"`

	// The identifier of the Bank Transfer
	BankTransferID string `json:"BankTransferID,omitempty" xml:"BankTransferID,omitempty"`

	// The currency rate
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// The Bank Transaction ID for the source account
	FromBankTransactionID string `json:"FromBankTransactionID,omitempty" xml:"FromBankTransactionID,omitempty"`

	// The Bank Transaction ID for the destination account
	ToBankTransactionID string `json:"ToBankTransactionID,omitempty" xml:"ToBankTransactionID,omitempty"`

	// Boolean to indicate if a Bank Transfer has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`

	// UTC timestamp of creation date of bank transfer
	CreatedDateUTC string `json:"CreatedDateUTC,omitempty" xml:"CreatedDateUTC,omitempty"`

	// The source BankAccount
	FromBankAccount BankAccount `json:"FromBankAccount,omitempty" xml:"FromBankAccount,omitempty"`

	// The destination BankAccount
	ToBankAccount BankAccount `json:"ToBankAccount,omitempty" xml:"ToBankAccount,omitempty"`
}

//BankTransfers contains a collection of BankTransfers
type BankTransfers struct {
	BankTransfers []BankTransfer `json:"BankTransfers" xml:"BankTransfer"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (b *BankTransfers) convertDates() error {
	var err error
	for n := len(b.BankTransfers) - 1; n >= 0; n-- {
		b.BankTransfers[n].Date, err = helpers.DotNetJSONTimeToRFC3339(b.BankTransfers[n].Date, false)
		if err != nil {
			return err
		}
		b.BankTransfers[n].CreatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(b.BankTransfers[n].CreatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalBankTransfer(bankTransferResponseBytes []byte) (*BankTransfers, error) {
	var bankTransferResponse *BankTransfers
	err := json.Unmarshal(bankTransferResponseBytes, &bankTransferResponse)
	if err != nil {
		return nil, err
	}

	err = bankTransferResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return bankTransferResponse, err
}

//Create will create bankTransfers given a BankTransfers struct
func (b *BankTransfers) Create(provider *xerogolang.Provider, session goth.Session) (*BankTransfers, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(b, "  ", "	")
	if err != nil {
		return nil, err
	}

	bankTransferResponseBytes, err := provider.Create(session, "BankTransfers", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransfer(bankTransferResponseBytes)
}

//FindBankTransfersModifiedSince will get all BankTransfers modified after a specified date.
//These BankTransfers will not have details like default line items by default.
//If you need details then add a 'page' querystringParameter and get 100 BankTransfers at a time
//additional querystringParameters such as where and order can be added as a map
func FindBankTransfersModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*BankTransfers, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	bankTransferResponseBytes, err := provider.Find(session, "BankTransfers", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransfer(bankTransferResponseBytes)
}

//FindBankTransfers will get all BankTransfers. These BankTransfer will not have details like line items by default.
//If you need details then add a 'page' querystringParameter and get 100 BankTransfers at a time
//additional querystringParameters such as where and order can be added as a map
func FindBankTransfers(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*BankTransfers, error) {
	return FindBankTransfersModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindBankTransfer will get a single bankTransfer - bankTransferID can be a GUID for an bankTransfer or an bankTransfer number
func FindBankTransfer(provider *xerogolang.Provider, session goth.Session, bankTransferID string) (*BankTransfers, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	bankTransferResponseBytes, err := provider.Find(session, "BankTransfers/"+bankTransferID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalBankTransfer(bankTransferResponseBytes)
}

//GenerateExampleBankTransfer Creates an Example bankTransfer
func GenerateExampleBankTransfer() *BankTransfers {
	bankTransfer := BankTransfer{
		FromBankAccount: BankAccount{
			Code: "090",
		},
		ToBankAccount: BankAccount{
			Code: "091",
		},
		Amount: 100.00,
	}

	bankTransferCollection := &BankTransfers{
		BankTransfers: []BankTransfer{},
	}

	bankTransferCollection.BankTransfers = append(bankTransferCollection.BankTransfers, bankTransfer)

	return bankTransferCollection
}
