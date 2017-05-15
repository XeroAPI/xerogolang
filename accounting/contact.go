package accounting

type Contact struct {

	// Xero identifier
	ContactID string `json:"ContactID,omitempty" xml:"ContactID,omitempty"`

	// This can be updated via the API only i.e. This field is read only on the Xero contact screen, used to identify contacts in external systems (max length = 50). If the Contact Number is used, this is displayed as Contact Code in the Contacts UI in Xero.
	ContactNumber string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`

	// A user defined account number. This can be updated via the API and the Xero UI (max length = 50)
	AccountNumber string `json:"AccountNumber,omitempty" xml:"AccountNumber,omitempty"`

	// Current status of a contact – see contact status types
	ContactStatus string `json:"ContactStatus,omitempty" xml:"ContactStatus,omitempty"`

	// Full name of contact/organisation (max length = 255)
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// First name of contact person (max length = 255)
	FirstName string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`

	// Last name of contact person (max length = 255)
	LastName string `json:"LastName,omitempty" xml:"LastName,omitempty"`

	// Email address of contact person (umlauts not supported) (max length = 255)
	EmailAddress string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`

	// Skype user name of contact
	SkypeUserName string `json:"SkypeUserName,omitempty" xml:"SkypeUserName,omitempty"`

	// See contact persons
	ContactPersons []ContactPerson `json:"ContactPersons,omitempty" xml:"ContactPersons,omitempty"`

	// Bank account number of contact
	BankAccountDetails string `json:"BankAccountDetails,omitempty" xml:"BankAccountDetails,omitempty"`

	// Tax number of contact – this is also known as the ABN (Australia), GST Number (New Zealand), VAT Number (UK) or Tax ID Number (US and global) in the Xero UI depending on which regionalized version of Xero you are using (max length = 50)
	TaxNumber string `json:"TaxNumber,omitempty" xml:"TaxNumber,omitempty"`

	// Default tax type used for contact on AR invoices
	AccountsReceivableTaxType string `json:"AccountsReceivableTaxType,omitempty" xml:"AccountsReceivableTaxType,omitempty"`

	// Default tax type used for contact on AP invoices
	AccountsPayableTaxType string `json:"AccountsPayableTaxType,omitempty" xml:"AccountsPayableTaxType,omitempty"`

	// Store certain address types for a contact – see address types
	Addresses []Address `json:"Addresses,omitempty" xml:"Addresses,omitempty"`

	// Store certain phone types for a contact – see phone types
	Phones []Phone `json:"Phones,omitempty" xml:"Phones,omitempty"`

	// true or false – Boolean that describes if a contact that has any AP invoices entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts payable invoice is generated against this contact.
	IsSupplier bool `json:"IsSupplier,omitempty" xml:"IsSupplier,omitempty"`

	// true or false – Boolean that describes if a contact has any AR invoices entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts receivable invoice is generated against this contact.
	IsCustomer bool `json:"IsCustomer,omitempty" xml:"IsCustomer,omitempty"`

	// Default currency for raising invoices against contact
	DefaultCurrency string `json:"DefaultCurrency,omitempty" xml:"DefaultCurrency,omitempty"`

	// Store XeroNetworkKey for contacts.
	XeroNetworkKey string `json:"XeroNetworkKey,omitempty" xml:"XeroNetworkKey,omitempty"`

	// The default sales account code for contacts
	SalesDefaultAccountCode string `json:"SalesDefaultAccountCode,omitempty" xml:"SalesDefaultAccountCode,omitempty"`

	// The default purchases account code for contacts
	PurchasesDefaultAccountCode string `json:"PurchasesDefaultAccountCode,omitempty" xml:"PurchasesDefaultAccountCode,omitempty"`

	// The default sales tracking categories for contacts
	SalesTrackingCategories []TrackingCategory `json:"SalesTrackingCategories,omitempty" xml:"SalesTrackingCategories,omitempty"`

	// The default purchases tracking categories for contacts
	PurchasesTrackingCategories []TrackingCategory `json:"PurchasesTrackingCategories,omitempty" xml:"PurchasesTrackingCategories,omitempty"`

	// The name of the Tracking Category assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories
	TrackingCategoryName string `json:"TrackingCategoryName,omitempty" xml:"TrackingCategoryName,omitempty"`

	// The name of the Tracking Option assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories
	TrackingCategoryOption string `json:"TrackingCategoryOption,omitempty" xml:"TrackingCategoryOption,omitempty"`

	// UTC timestamp of last update to contact
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`

	// Displays which contact groups a contact is included in
	ContactGroups []ContactGroup `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`

	// Website address for contact (read only)
	Website string `json:"Website,omitempty" xml:"Website,omitempty"`

	// batch payment details for contact (read only)
	BatchPayments BatchPayment `json:"BatchPayments,omitempty" xml:"BatchPayments,omitempty"`

	// The default discount rate for the contact (read only)
	Discount float32 `json:"Discount,omitempty" xml:"Discount,omitempty"`

	// The raw AccountsReceivable(sales invoices) and AccountsPayable(bills) outstanding and overdue amounts, not converted to base currency (read only)
	Balances string `json:"Balances,omitempty" xml:"Balances,omitempty"`

	// A boolean to indicate if a contact has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty" xml:"HasAttachments,omitempty"`
}
