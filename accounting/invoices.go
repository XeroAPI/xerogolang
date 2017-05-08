package accounting

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"time"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/model"
	"github.com/gorilla/sessions"
)

//CreateInvoice will create an invoice given an Invoices struct
//This will only handle single invoices at the moment
var CreateInvoice = func(res http.ResponseWriter, req *http.Request, provider *xero.Provider, store sessions.Store, invoices *model.Invoices) (*model.Invoices, error) {

	body, err := xml.MarshalIndent(invoices, "  ", "	")
	if err != nil {
		return nil, err
	}

	session, err := provider.GetSessionFromStore(req, store)
	if err != nil {
		return nil, err
	}

	invoiceResponseBytes, err := provider.Create(session, "Invoices", body)
	if err != nil {
		return nil, err
	}

	var invoiceResponse *model.Invoices
	err = json.Unmarshal(invoiceResponseBytes, &invoiceResponse)
	if err != nil {
		return nil, err
	}
	return invoiceResponse, err
}

//UpdateInvoice will update an invoice given an Invoices struct
//This will only handle single invoices at the moment
var UpdateInvoice = func(res http.ResponseWriter, req *http.Request, provider *xero.Provider, store sessions.Store, invoices *model.Invoices) (*model.Invoices, error) {

	body, err := xml.MarshalIndent(invoices, "  ", "	")
	if err != nil {
		return nil, err
	}

	session, err := provider.GetSessionFromStore(req, store)
	if err != nil {
		return nil, err
	}

	invoiceResponseBytes, err := provider.Update(session, "Invoices/"+invoices.Invoices[0].InvoiceID, body)
	if err != nil {
		return nil, err
	}

	var invoiceResponse *model.Invoices
	err = json.Unmarshal(invoiceResponseBytes, &invoiceResponse)
	if err != nil {
		return nil, err
	}
	return invoiceResponse, err
}

//CreateExampleInvoice Creates an Example invoice
func CreateExampleInvoice() *model.Invoices {
	lineItem := model.LineItem{
		Description: "Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	now := time.Now()

	invoice := model.Invoice{
		Type: "ACCREC",
		Contact: model.Contact{
			Name: "George Costanza",
		},
		Date:            string(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format(time.RFC3339)),
		DueDate:         string(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(720 * time.Hour).Format(time.RFC3339)),
		LineAmountTypes: "Exclusive",
		LineItems:       []model.LineItem{},
	}

	invoice.LineItems = append(invoice.LineItems, lineItem)

	invoiceCollection := &model.Invoices{
		Invoices: []model.Invoice{},
	}

	invoiceCollection.Invoices = append(invoiceCollection.Invoices, invoice)

	return invoiceCollection
}
