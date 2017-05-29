package accounting

import (
	"encoding/json"
	"encoding/xml"
	"time"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/helpers"
	"github.com/markbates/goth"
)

//Item is something that is sold or purchased.  It can have inventory tracked or not tracked.
type Item struct {

	// User defined item code (max length = 30)
	Code string `json:"Code" xml:"Code"`

	// The inventory asset account for the item. The account must be of type INVENTORY. The  COGSAccountCode in PurchaseDetails is also required to create a tracked item
	InventoryAssetAccountCode string `json:"InventoryAssetAccountCode" xml:"InventoryAssetAccountCode"`

	// The name of the item (max length = 50)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// Boolean value, defaults to true. When IsSold is true the item will be available on sales transactions in the Xero UI. If IsSold is updated to false then Description and SalesDetails values will be nulled.
	IsSold bool `json:"IsSold,omitempty" xml:"IsSold,omitempty"`

	// Boolean value, defaults to true. When IsPurchased is true the item is available for purchase transactions in the Xero UI. If IsPurchased is updated to false then PurchaseDescription and PurchaseDetails values will be nulled.
	IsPurchased bool `json:"IsPurchased,omitempty" xml:"IsPurchased,omitempty"`

	// The sales description of the item (max length = 4000)
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// The purchase description of the item (max length = 4000)
	PurchaseDescription string `json:"PurchaseDescription,omitempty" xml:"PurchaseDescription,omitempty"`

	// See Purchases & Sales
	PurchaseDetails PurchaseAndSaleDetails `json:"PurchaseDetails,omitempty" xml:"PurchaseDetails,omitempty"`

	// See Purchases & Sales
	SalesDetails PurchaseAndSaleDetails `json:"SalesDetails,omitempty" xml:"SalesDetails,omitempty"`

	// True for items that are tracked as inventory. An item will be tracked as inventory if the InventoryAssetAccountCode and COGSAccountCode are set.
	IsTrackedAsInventory bool `json:"IsTrackedAsInventory,omitempty" xml:"-"`

	// The value of the item on hand. Calculated using average cost accounting.
	TotalCostPool float32 `json:"TotalCostPool,omitempty" xml:"-"`

	// The quantity of the item on hand
	QuantityOnHand float32 `json:"QuantityOnHand,omitempty" xml:"-"`

	// Last modified date in UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"-"`

	// The Xero identifier for an Item
	ItemID string `json:"ItemID,omitempty" xml:"ItemID,omitempty"`
}

//Items is a collection of Items
type Items struct {
	Items []Item `json:"Items" xml:"Item"`
}

//PurchaseAndSaleDetails are Elements for Purchases and Sales
type PurchaseAndSaleDetails struct {
	//Unit Price of the item. By default UnitPrice is returned to two decimal places.  You can use 4 decimal places by adding the unitdp=4 querystring parameter to your request.
	UnitPrice float32 `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`

	//Default account code to be used for purchased/sale. Not applicable to the purchase details of tracked items
	AccountCode string `json:"AccountCode,omitempty" xml:"AccountCode,omitempty"`

	//Cost of goods sold account. Only applicable to the purchase details of tracked items.
	COGSAccountCode string `json:"COGSAccountCode,omitempty" xml:"COGSAccountCode,omitempty"`

	//Used as an override if the default Tax Code for the selected AccountCode is not correct - see TaxTypes.
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (i *Items) convertItemDates() error {
	var err error
	for n := len(i.Items) - 1; n >= 0; n-- {
		i.Items[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(i.Items[n].UpdatedDateUTC, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalItem(itemResponseBytes []byte) (*Items, error) {
	var itemResponse *Items
	err := json.Unmarshal(itemResponseBytes, &itemResponse)
	if err != nil {
		return nil, err
	}

	err = itemResponse.convertItemDates()
	if err != nil {
		return nil, err
	}

	return itemResponse, err
}

//CreateItem will create items given an Items struct
func (i *Items) CreateItem(provider *xero.Provider, session goth.Session) (*Items, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(i, "  ", "	")
	if err != nil {
		return nil, err
	}

	itemResponseBytes, err := provider.Create(session, "Items", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalItem(itemResponseBytes)
}

//UpdateItem will update an item given an Items struct
//This will only handle single item - you cannot update multiple items in a single call
func (i *Items) UpdateItem(provider *xero.Provider, session goth.Session) (*Items, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(i, "  ", "	")
	if err != nil {
		return nil, err
	}

	itemResponseBytes, err := provider.Update(session, "Items/"+i.Items[0].ItemID, additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalItem(itemResponseBytes)
}

//FindItemsModifiedSinceWithParams will get all items modified after a specified date.
//additional querystringParameters such as where, page, order can be added as a map
func FindItemsModifiedSinceWithParams(provider *xero.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Items, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	itemResponseBytes, err := provider.Find(session, "Items", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalItem(itemResponseBytes)
}

//FindItemsModifiedSince will get all items modified after a specified date.
func FindItemsModifiedSince(provider *xero.Provider, session goth.Session, modifiedSince time.Time) (*Items, error) {
	return FindItemsModifiedSinceWithParams(provider, session, modifiedSince, nil)
}

//FindItemsModifiedSinceWhere will get all items modified after a specified date that fit the criteria of a supplied where clause.
func FindItemsModifiedSinceWhere(provider *xero.Provider, session goth.Session, modifiedSince time.Time, whereClause string) (*Items, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindItemsModifiedSinceWithParams(provider, session, modifiedSince, querystringParameters)
}

//FindItemsWhere will get items that fit the criteria of a supplied where clause.
func FindItemsWhere(provider *xero.Provider, session goth.Session, whereClause string) (*Items, error) {
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	return FindItemsModifiedSinceWithParams(provider, session, dayZero, querystringParameters)
}

//FindItems will get all items.
func FindItems(provider *xero.Provider, session goth.Session) (*Items, error) {
	return FindItemsModifiedSinceWithParams(provider, session, dayZero, nil)
}

//FindItem will get a single item - itemID must be a GUID for an item
func FindItem(provider *xero.Provider, session goth.Session, itemID string) (*Items, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	itemResponseBytes, err := provider.Find(session, "Items/"+itemID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalItem(itemResponseBytes)
}

//RemoveItem will get a single item - itemID must be a GUID for an item
func RemoveItem(provider *xero.Provider, session goth.Session, itemID string) (*Items, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	itemResponseBytes, err := provider.Remove(session, "Items/"+itemID, additionalHeaders)
	if err != nil {
		return nil, err
	}

	return unmarshalItem(itemResponseBytes)
}

//CreateExampleItem Creates an Example item
func CreateExampleItem() *Items {
	item := Item{
		Code:                "42",
		Name:                "The Executive",
		Description:         "A Beltless Trenchcoat",
		PurchaseDescription: "A Beltless Trenchcoat",
		IsSold:              true,
		IsPurchased:         true,
		PurchaseDetails: PurchaseAndSaleDetails{
			UnitPrice:   140.00,
			AccountCode: "300",
		},
		SalesDetails: PurchaseAndSaleDetails{
			UnitPrice:   300.00,
			AccountCode: "200",
		},
	}

	itemCollection := &Items{
		Items: []Item{},
	}

	itemCollection.Items = append(itemCollection.Items, item)

	return itemCollection
}
