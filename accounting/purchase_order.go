package accounting

type PurchaseOrder struct {

	// See LineItems
	LineItems []LineItem `json:"LineItems"`

	// Date purchase order was issued – YYYY-MM-DD. If the Date element is not specified then it will default to the current date based on the timezone setting of the organisation
	Date string `json:"Date,omitempty"`

	// Date the goods are to be delivered – YYYY-MM-DD
	DeliveryDate string `json:"DeliveryDate,omitempty"`

	// Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
	LineAmountTypes string `json:"LineAmountTypes,omitempty"`

	// Unique alpha numeric code identifying purchase order (when missing will auto-generate from your Organisation Invoice Settings)
	PurchaseOrderNumber string `json:"PurchaseOrderNumber,omitempty"`

	// Additional reference number
	Reference string `json:"Reference,omitempty"`

	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty"`

	// The currency that purchase order has been raised in (see Currencies)
	CurrencyCode string `json:"CurrencyCode,omitempty"`

	// See Purchase Order Status Codes
	Status string `json:"Status,omitempty"`

	// Boolean to set whether the purchase order should be marked as “sent”. This can be set only on purchase orders that have been approved or billed
	SentToContact bool `json:"SentToContact,omitempty"`

	// The address the goods are to be delivered to
	DeliveryAddress string `json:"DeliveryAddress,omitempty"`

	// The person that the delivery is going to
	AttentionTo string `json:"AttentionTo,omitempty"`

	// The phone number for the person accepting the delivery
	Telephone string `json:"Telephone,omitempty"`

	// A free text feild for instructions (500 characters max)
	DeliveryInstructions string `json:"DeliveryInstructions,omitempty"`

	// The date the goods are expected to arrive.
	ExpectedArrivalDate string `json:"ExpectedArrivalDate,omitempty"`

	// Xero generated unique identifier for purchase order
	PurchaseOrderID string `json:"PurchaseOrderID,omitempty"`

	// The currency rate for a multicurrency purchase order. As no rate can be specified, the XE.com day rate is used.
	CurrencyRate float32 `json:"CurrencyRate,omitempty"`

	// Total of purchase order excluding taxes
	SubTotal float32 `json:"SubTotal,omitempty"`

	// Total tax on purchase order
	TotalTax float32 `json:"TotalTax,omitempty"`

	// Total of Purchase Order tax inclusive (i.e. SubTotal + TotalTax)
	Total float32 `json:"Total,omitempty"`

	// Total of discounts applied on the purchase order line items
	TotalDiscount float32 `json:"TotalDiscount,omitempty"`

	// boolean to indicate if a purchase order has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`

	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`
}
