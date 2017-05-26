package accounting

import (
	"encoding/json"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/helpers"
	"github.com/markbates/goth"
)

//Organisation is information about a Xero organisation
type Organisation struct {

	// Display a unique key used for Xero-to-Xero transactions
	APIKey string `json:"APIKey,omitempty"`

	// Display name of organisation shown in Xero
	Name string `json:"Name,omitempty"`

	// Organisation name shown on Reports
	LegalName string `json:"LegalName,omitempty"`

	// Boolean to describe if organisation is registered with a local tax authority i.e. true, false
	PaysTax bool `json:"PaysTax,omitempty"`

	// See Version Types
	Version string `json:"Version,omitempty"`

	// Organisation Type
	OrganisationType string `json:"OrganisationType,omitempty"`

	// Default currency for organisation. See ISO 4217 Currency Codes
	BaseCurrency string `json:"BaseCurrency,omitempty"`

	// Country code for organisation. See ISO 3166-2 Country Codes
	CountryCode string `json:"CountryCode,omitempty"`

	// Boolean to describe if organisation is a demo company.
	IsDemoCompany bool `json:"IsDemoCompany,omitempty"`

	// Will be set to ACTIVE if you can connect to organisation via the Xero API
	OrganisationStatus string `json:"OrganisationStatus,omitempty"`

	// Shows for New Zealand, Australian and UK organisations
	RegistrationNumber string `json:"RegistrationNumber,omitempty"`

	// Shown if set. Displays in the Xero UI as Tax File Number (AU), GST Number (NZ), VAT Number (UK) and Tax ID Number (US & Global).
	TaxNumber string `json:"TaxNumber,omitempty"`

	// Calendar day e.g. 0-31
	FinancialYearEndDay int `json:"FinancialYearEndDay,omitempty"`

	// Calendar Month e.g. 1-12
	FinancialYearEndMonth int `json:"FinancialYearEndMonth,omitempty"`

	// The accounting basis used for tax returns. See Sales Tax Basis
	SalesTaxBasis string `json:"SalesTaxBasis,omitempty"`

	// The frequency with which tax returns are processed. See Sales Tax Period
	SalesTaxPeriod string `json:"SalesTaxPeriod,omitempty"`

	// The default for LineAmountTypes on sales transactions
	DefaultSalesTax string `json:"DefaultSalesTax,omitempty"`

	// The default for LineAmountTypes on purchase transactions
	DefaultPurchasesTax string `json:"DefaultPurchasesTax,omitempty"`

	// Shown if set. See lock dates
	PeriodLockDate string `json:"PeriodLockDate,omitempty"`

	// Shown if set. See lock dates
	EndOfYearLockDate string `json:"EndOfYearLockDate,omitempty"`

	// Timestamp when the organisation was created in Xero
	CreatedDateUTC string `json:"CreatedDateUTC,omitempty"`

	// Timezone specifications
	Timezone string `json:"Timezone,omitempty"`

	// Organisation Type
	OrganisationEntityType string `json:"OrganisationEntityType,omitempty"`

	// A unique identifier for the organisation. Potential uses.
	ShortCode string `json:"ShortCode,omitempty"`

	// Description of business type as defined in Organisation settings
	LineOfBusiness string `json:"LineOfBusiness,omitempty"`

	// Address details for organisation – see Addresses
	Addresses []Address `json:"Addresses,omitempty"`

	// Phones details for organisation – see Phones
	Phones []Phone `json:"Phones,omitempty"`

	// Organisation profile links for popular services such as Facebook, Twitter, GooglePlus and LinkedIn. You can also add link to your website here. Shown if Organisation settings  is updated in Xero. See ExternalLinks below
	ExternalLinks []ExternalLink `json:"ExternalLinks,omitempty"`
}

//OrganisationCollection contains a collection of Organisations - but there will only ever be one. Like Highlander
type OrganisationCollection struct {
	Organisations []Organisation `json:"Organisations,omitempty"`
}

//FindOrganisation returns details about the Xero organisation you're connected to
func FindOrganisation(provider *xero.Provider, session goth.Session) (*OrganisationCollection, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	organisationResponseBytes, err := provider.Find(session, "Organisation", additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	var organisationResponse *OrganisationCollection
	err = json.Unmarshal(organisationResponseBytes, &organisationResponse)
	if err != nil {
		return nil, err
	}

	organisationResponse.Organisations[0].PeriodLockDate, err = helpers.DotNetJSONTimeToRFC3339(organisationResponse.Organisations[0].PeriodLockDate, false)
	if err != nil {
		return nil, err
	}
	organisationResponse.Organisations[0].CreatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(organisationResponse.Organisations[0].CreatedDateUTC, true)
	if err != nil {
		return nil, err
	}

	return organisationResponse, nil
}
