package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
)

//PurchaseOrder is the first official offer issued by a buyer to a seller, indicating types, quantities, and agreed prices for products or services
type PurchaseOrder struct {

	// See LineItems
	LineItems []LineItem `json:"LineItems" xml:"LineItems>LineItem"`

	// See Contacts
	Contact Contact `json:"Contact" xml:"Contact"`

	// Date purchase order was issued – YYYY-MM-DD. If the Date element is not specified then it will default to the current date based on the timezone setting of the organisation
	Date string `json:"DateString,omitempty" xml:"Date,omitempty"`

	// Date the goods are to be delivered – YYYY-MM-DD
	DeliveryDate string `json:"DeliveryDateString,omitempty" xml:"DeliveryDate,omitempty"`

	// Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`

	// Unique alpha numeric code identifying purchase order (when missing will auto-generate from your Organisation Invoice Settings)
	PurchaseOrderNumber string `json:"PurchaseOrderNumber,omitempty" xml:"PurchaseOrderNumber,omitempty"`

	// Additional reference number
	Reference string `json:"Reference,omitempty" xml:"Reference,omitempty"`

	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty" xml:"BrandingThemeID,omitempty"`

	// The currency that purchase order has been raised in (see Currencies)
	CurrencyCode string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`

	// See Purchase Order Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// Boolean to set whether the purchase order should be marked as “sent”. This can be set only on purchase orders that have been approved or billed
	SentToContact bool `json:"SentToContact,omitempty" xml:"SentToContact,omitempty"`

	// The address the goods are to be delivered to
	DeliveryAddress string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`

	// The person that the delivery is going to
	AttentionTo string `json:"AttentionTo,omitempty" xml:"AttentionTo,omitempty"`

	// The phone number for the person accepting the delivery
	Telephone string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`

	// A free text feild for instructions (500 characters max)
	DeliveryInstructions string `json:"DeliveryInstructions,omitempty" xml:"DeliveryInstructions,omitempty"`

	// The date the goods are expected to arrive.
	ExpectedArrivalDate string `json:"ExpectedArrivalDate,omitempty" xml:"ExpectedArrivalDate,omitempty"`

	// Xero generated unique identifier for purchase order
	PurchaseOrderID string `json:"PurchaseOrderID,omitempty" xml:"PurchaseOrderID,omitempty"`

	// The currency rate for a multicurrency purchase order. As no rate can be specified, the XE.com day rate is used.
	CurrencyRate float32 `json:"CurrencyRate,omitempty" xml:"CurrencyRate,omitempty"`

	// Total of purchase order excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`

	// Total tax on purchase order
	TotalTax float32 `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`

	// Total of Purchase Order tax inclusive (i.e. SubTotal + TotalTax)
	Total float32 `json:"Total,omitempty" xml:"Total,omitempty"`

	// Total of discounts applied on the purchase order line items
	TotalDiscount float32 `json:"TotalDiscount,omitempty" xml:"TotalDiscount,omitempty"`

	// boolean to indicate if a purchase order has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"-"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`
}

//PurchaseOrders contains a collection of PurchaseOrders
type PurchaseOrders struct {
	PurchaseOrders []PurchaseOrder `json:"PurchaseOrders" xml:"PurchaseOrder"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (p *PurchaseOrders) convertDates() error {
	var err error
	for n := len(p.PurchaseOrders) - 1; n >= 0; n-- {
		p.PurchaseOrders[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(p.PurchaseOrders[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalPurchaseOrder(purchaseOrderResponseBytes []byte) (*PurchaseOrders, error) {
	var purchaseOrderResponse *PurchaseOrders
	err := json.Unmarshal(purchaseOrderResponseBytes, &purchaseOrderResponse)
	if err != nil {
		return nil, err
	}

	err = purchaseOrderResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return purchaseOrderResponse, err
}

//Create will create purchaseOrders given an PurchaseOrders struct
func (p *PurchaseOrders) Create(provider *xerogolang.Provider, session goth.Session) (*PurchaseOrders, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(p, "  ", "	")
	if err != nil {
		return nil, err
	}

	purchaseOrderResponseBytes, err := provider.Create(session, "PurchaseOrders", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalPurchaseOrder(purchaseOrderResponseBytes)
}

//Update will update an purchaseOrder given an PurchaseOrders struct
//This will only handle single purchaseOrder - you cannot update multiple purchaseOrders in a single call
func (p *PurchaseOrders) Update(provider *xerogolang.Provider, session goth.Session) (*PurchaseOrders, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(p, "  ", "	")
	if err != nil {
		return nil, err
	}

	purchaseOrderResponseBytes, err := provider.Update(session, "PurchaseOrders/"+p.PurchaseOrders[0].PurchaseOrderID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalPurchaseOrder(purchaseOrderResponseBytes)
}

//FindPurchaseOrdersModifiedSince will get all PurchaseOrders modified after a specified date.
//Paging is enforced by default. 100 purchase orders are returned per page.
//additional querystringParameters such as page, order, status, DateFrom & DateTo can be added as a map
func FindPurchaseOrdersModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*PurchaseOrders, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	purchaseOrderResponseBytes, err := provider.Find(session, "PurchaseOrders", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalPurchaseOrder(purchaseOrderResponseBytes)
}

//FindPurchaseOrders will get all PurchaseOrders. Paging is enforced by default. 100 purchase orders are returned per page.
//additional querystringParameters such as page, order, status, DateFrom & DateTo can be added as a map
func FindPurchaseOrders(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*PurchaseOrders, error) {
	return FindPurchaseOrdersModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindPurchaseOrder will get a single purchaseOrder - purchaseOrderID can be a GUID for an purchaseOrder or an purchaseOrder number
func FindPurchaseOrder(provider *xerogolang.Provider, session goth.Session, purchaseOrderID string) (*PurchaseOrders, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	purchaseOrderResponseBytes, err := provider.Find(session, "PurchaseOrders/"+purchaseOrderID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalPurchaseOrder(purchaseOrderResponseBytes)
}

//GenerateExamplePurchaseOrder Creates an Example purchaseOrder
func GenerateExamplePurchaseOrder(contactID string) *PurchaseOrders {
	lineItem := LineItem{
		Description: "Importing & Exporting Services",
		Quantity:    1.00,
		UnitAmount:  395.00,
		AccountCode: "200",
	}

	purchaseOrder := PurchaseOrder{
		Contact: Contact{
			ContactID: contactID,
		},
		Date:            helpers.TodayRFC3339(),
		LineAmountTypes: "Exclusive",
		LineItems:       []LineItem{},
	}

	purchaseOrder.LineItems = append(purchaseOrder.LineItems, lineItem)

	purchaseOrderCollection := &PurchaseOrders{
		PurchaseOrders: []PurchaseOrder{},
	}

	purchaseOrderCollection.PurchaseOrders = append(purchaseOrderCollection.PurchaseOrders, purchaseOrder)

	return purchaseOrderCollection
}
