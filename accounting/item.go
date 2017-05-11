package accounting

type Item struct {

	// User defined item code (max length = 30)
	Code string `json:"Code"`

	// The inventory asset account for the item. The account must be of type INVENTORY. The  COGSAccountCode in PurchaseDetails is also required to create a tracked item
	InventoryAssetAccountCode string `json:"InventoryAssetAccountCode"`

	// The name of the item (max length = 50)
	Name string `json:"Name,omitempty"`

	// Boolean value, defaults to true. When IsSold is true the item will be available on sales transactions in the Xero UI. If IsSold is updated to false then Description and SalesDetails values will be nulled.
	IsSold bool `json:"IsSold,omitempty"`

	// Boolean value, defaults to true. When IsPurchased is true the item is available for purchase transactions in the Xero UI. If IsPurchased is updated to false then PurchaseDescription and PurchaseDetails values will be nulled.
	IsPurchased bool `json:"IsPurchased,omitempty"`

	// The sales description of the item (max length = 4000)
	Description string `json:"Description,omitempty"`

	// The purchase description of the item (max length = 4000)
	PurchaseDescription string `json:"PurchaseDescription,omitempty"`

	// See Purchases & Sales
	PurchaseDetails []Purchase `json:"PurchaseDetails,omitempty"`

	// See Purchases & Sales
	SalesDetails string `json:"SalesDetails,omitempty"`

	// True for items that are tracked as inventory. An item will be tracked as inventory if the InventoryAssetAccountCode and COGSAccountCode are set.
	IsTrackedAsInventory bool `json:"IsTrackedAsInventory,omitempty"`

	// The value of the item on hand. Calculated using average cost accounting.
	TotalCostPool string `json:"TotalCostPool,omitempty"`

	// The quantity of the item on hand
	QuantityOnHand string `json:"QuantityOnHand,omitempty"`

	// Last modified date in UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`

	// The Xero identifier for an Item
	ItemID string `json:"ItemID,omitempty"`
}
