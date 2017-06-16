package accounting

import (
	"encoding/json"
	"encoding/xml"

	"github.com/XeroAPI/xerogolang"
	"github.com/markbates/goth"
)

//TaxRate is a rate at which an item or service is taxed when sold or purchased
type TaxRate struct {

	// Name of tax rate
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`

	// See Tax Types – can only be used on update calls
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`

	// See TaxComponents
	TaxComponents []TaxComponent `json:"TaxComponents,omitempty" xml:"TaxComponents>TaxComponent,omitempty"`

	// See Status Codes
	Status string `json:"Status,omitempty" xml:"Status,omitempty"`

	// See ReportTaxTypes
	ReportTaxType string `json:"ReportTaxType" xml:"ReportTaxType"`

	// Boolean to describe if tax rate can be used for asset accounts i.e. true,false
	CanApplyToAssets bool `json:"CanApplyToAssets,omitempty" xml:"CanApplyToAssets,omitempty"`

	// Boolean to describe if tax rate can be used for equity accounts i.e. true,false
	CanApplyToEquity bool `json:"CanApplyToEquity,omitempty" xml:"CanApplyToEquity,omitempty"`

	// Boolean to describe if tax rate can be used for expense accounts i.e. true,false
	CanApplyToExpenses bool `json:"CanApplyToExpenses,omitempty" xml:"CanApplyToExpenses,omitempty"`

	// Boolean to describe if tax rate can be used for liability accounts i.e. true,false
	CanApplyToLiabilities bool `json:"CanApplyToLiabilities,omitempty" xml:"CanApplyToLiabilities,omitempty"`

	// Boolean to describe if tax rate can be used for revenue accounts i.e. true,false
	CanApplyToRevenue bool `json:"CanApplyToRevenue,omitempty" xml:"CanApplyToRevenue,omitempty"`

	// Tax Rate (decimal to 4dp) e.g 12.5000
	DisplayTaxRate float32 `json:"DisplayTaxRate,omitempty" xml:"DisplayTaxRate,omitempty"`

	// Effective Tax Rate (decimal to 4dp) e.g 12.5000
	EffectiveRate float32 `json:"EffectiveRate,omitempty" xml:"EffectiveRate,omitempty"`
}

type TaxRates struct {
	TaxRates []TaxRate `json:"TaxRates" xml:"TaxRate"`
}

func unmarshalTaxRate(taxRateResponseBytes []byte) (*TaxRates, error) {
	var taxRateResponse *TaxRates
	err := json.Unmarshal(taxRateResponseBytes, &taxRateResponse)
	if err != nil {
		return nil, err
	}

	return taxRateResponse, err
}

//Create will create taxRates given an TaxRates struct
func (t *TaxRates) Create(provider *xerogolang.Provider, session goth.Session) (*TaxRates, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(t, "  ", "	")
	if err != nil {
		return nil, err
	}

	taxRateResponseBytes, err := provider.Create(session, "TaxRates", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalTaxRate(taxRateResponseBytes)
}

//Update will update an taxRate given an TaxRates struct
//This will only handle a single taxRate - you cannot update multiple taxRates in a single call
func (t *TaxRates) Update(provider *xerogolang.Provider, session goth.Session) (*TaxRates, error) {
	additionalHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/xml",
	}

	body, err := xml.MarshalIndent(t, "  ", "	")
	if err != nil {
		return nil, err
	}

	taxRateResponseBytes, err := provider.Update(session, "TaxRates", additionalHeaders, body)
	if err != nil {
		return nil, err
	}

	return unmarshalTaxRate(taxRateResponseBytes)
}

//FindTaxRates will get all TaxRates.
//additional querystringParameters such as taxType, where and order can be added as a map
func FindTaxRates(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*TaxRates, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	taxRateResponseBytes, err := provider.Find(session, "TaxRates", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalTaxRate(taxRateResponseBytes)
}

//GenerateExampleTaxRate Creates an Example taxRate
func GenerateExampleTaxRate() *TaxRates {
	taxComponent1 := TaxComponent{
		Name:       "State Tax",
		Rate:       7.5,
		IsCompound: false,
	}

	taxComponent2 := TaxComponent{
		Name:       "Local Sales Tax",
		Rate:       0.625,
		IsCompound: false,
	}

	taxRate := TaxRate{
		Name:          "Newman's Tax",
		TaxComponents: []TaxComponent{},
		ReportTaxType: "OUTPUT",
	}

	taxRate.TaxComponents = append(taxRate.TaxComponents, taxComponent1)

	taxRate.TaxComponents = append(taxRate.TaxComponents, taxComponent2)

	taxRateCollection := &TaxRates{
		TaxRates: []TaxRate{},
	}

	taxRateCollection.TaxRates = append(taxRateCollection.TaxRates, taxRate)

	return taxRateCollection
}
