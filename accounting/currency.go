package accounting

import (
	"encoding/json"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/markbates/goth"
)

//Currency is the local currency set up to be used in Xero
type Currency struct {

	// 3 letter alpha code for the currency – see list of currency codes
	Code string `json:"Code,omitempty" xml:"Code,omitempty"`

	// Name of Currency
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`
}

//Currencies is a collection of Currencies
type Currencies struct {
	Currencies []Currency `json:"Currencies,omitempty" xml:"Currency,omitempty"`
}

func unmarshalCurrencies(currencyResponseBytes []byte) (*Currencies, error) {
	var currencyResponse *Currencies
	err := json.Unmarshal(currencyResponseBytes, &currencyResponse)
	if err != nil {
		return nil, err
	}

	return currencyResponse, err
}

//FindCurrencies will get all currencies
func FindCurrencies(provider *xero.Provider, session goth.Session) (*Currencies, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	currencyResponseBytes, err := provider.Find(session, "Currencies", additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalCurrencies(currencyResponseBytes)
}
