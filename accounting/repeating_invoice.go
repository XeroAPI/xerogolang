package accounting

import (
	"encoding/json"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//RepeatingInvoice a template for the invoices you send and bills you receive regularly
type RepeatingInvoice struct {

	// See Invoice Types
	Type string `json:"Type,omitempty" xml:"Type,omitempty"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// See LineItems
	LineItems []LineItem `json:"LineItems,omitempty" xml:"LineItems>LineItem,omitempty"`

	// Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// ACCREC only – additional reference number
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty" xml:"BrandingThemeID,omitempty"`

	// The currency that invoice has been raised in (see Currencies)
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// One of the following : DRAFT or AUTHORISED – See Invoice Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Total of invoice excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// Total tax on invoice
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// Total of Invoice tax inclusive (i.e. SubTotal + TotalTax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// Xero generated unique identifier for repeating invoice template
	RepeatingInvoiceID string `json:"RepeatingInvoiceID,omitempty" xml:"RepeatingInvoiceID,omitempty"`

	// boolean to indicate if an invoice has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`

	//Specifies when the repeating invoice will be created
	Schedule Schedule `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

//RepeatingInvoices is a collection of RepeatingInvoices
type RepeatingInvoices struct {
	RepeatingInvoices []RepeatingInvoice `json:"RepeatingInvoices,omitempty" xml:"RepeatingInvoice,omitempty"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (i *RepeatingInvoices) convertDates() error {
	var err error
	for n := len(i.RepeatingInvoices) - 1; n >= 0; n-- {
		i.RepeatingInvoices[n].Schedule.EndDate, err = helpers.DotNetJSONTimeToRFC3339(i.RepeatingInvoices[n].Schedule.EndDate, false)
		if err != nil {
			return err
		}

		i.RepeatingInvoices[n].Schedule.StartDate, err = helpers.DotNetJSONTimeToRFC3339(i.RepeatingInvoices[n].Schedule.StartDate, false)
		if err != nil {
			return err
		}

		i.RepeatingInvoices[n].Schedule.NextScheduledDate, err = helpers.DotNetJSONTimeToRFC3339(i.RepeatingInvoices[n].Schedule.NextScheduledDate, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalRepeatingInvoices(repeatingInvoiceResponseBytes []byte) (*RepeatingInvoices, error) {
	var repeatingInvoiceResponse *RepeatingInvoices
	err := json.Unmarshal(repeatingInvoiceResponseBytes, &repeatingInvoiceResponse)
	if err != nil {
		return nil, err
	}

	err = repeatingInvoiceResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return repeatingInvoiceResponse, err
}

//FindRepeatingInvoices will get all repeatingInvoices
//additional querystringParameters such as where and order can be added as a map
func FindRepeatingInvoices(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*RepeatingInvoices, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	repeatingInvoiceResponseBytes, err := provider.Find(session, "RepeatingInvoices", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalRepeatingInvoices(repeatingInvoiceResponseBytes)
}

//FindRepeatingInvoice will get a single repeatingInvoice - RepeatingInvoiceID must be a GUID for a repeatingInvoice
func FindRepeatingInvoice(provider *xerogolang.Provider, session goth.Session, repeatingInvoiceID string) (*RepeatingInvoices, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	repeatingInvoiceResponseBytes, err := provider.Find(session, "RepeatingInvoices/"+repeatingInvoiceID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalRepeatingInvoices(repeatingInvoiceResponseBytes)
}
