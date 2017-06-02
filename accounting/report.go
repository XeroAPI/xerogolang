package accounting

import (
	"encoding/json"
	"strconv"

	xero "github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/helpers"
	"github.com/markbates/goth"
)

type Report struct {
	ReportID     string    `json:"ReportID,omitempty" xml:"ReportID,omitempty"`
	ReportName   string    `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	ReportType   string    `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	ReportTitles *[]string `json:"ReportTitles,omitempty" xml:"ReportTitles>ReportTitle,omitempty"`
	ReportDate   string    `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC string             `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`
	Attributes     *[]ReportAttribute `json:"Attributes,omitempty" xml:"Attributes>Attribute,omitempty"`
	Rows           *[]Row             `json:"Rows,omitempty" xml:"Rows>Row,omitempty"`
}

type Reports struct {
	Reports []Report `json:"Reports" xml:"Report"`
}

//ReportTitle is a title on a report
type ReportTitle struct {
	ReportTitle string `json:"ReportTitle,omitempty" xml:"ReportTitle,omitempty"`
}

//The Xero API returns Dates based on the .Net JSON date format available at the time of development
//We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (r *Reports) convertReportDates() error {
	var err error
	for n := len(r.Reports) - 1; n >= 0; n-- {
		if r.Reports[n].UpdatedDateUTC != "" {
			r.Reports[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(r.Reports[n].UpdatedDateUTC, true)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalReport(reportResponseBytes []byte) (*Reports, error) {
	var reportResponse *Reports
	err := json.Unmarshal(reportResponseBytes, &reportResponse)
	if err != nil {
		return nil, err
	}

	err = reportResponse.convertReportDates()
	if err != nil {
		return nil, err
	}

	return reportResponse, err
}

//Run1099 will run the 1099 Report and marshal the results to a Report Struct
//This Report will only work for US based Organisations
func Run1099(provider *xero.Provider, session goth.Session, reportYear int) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	querystringParameters := map[string]string{
		"reportYear": strconv.Itoa(reportYear),
	}

	reportResponseBytes, err := provider.Find(session, "Reports/TenNinetyNine", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunAgedPayablesByContact will run the Aged Payables By Contact Report and marshal the results to a Report Struct
//Date, FromDate and ToDate can be added as optional paramters as a map
func RunAgedPayablesByContact(provider *xero.Provider, session goth.Session, contactID string, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if querystringParameters != nil {
		querystringParameters["ContactID"] = contactID
	} else {
		querystringParameters = map[string]string{
			"ContactID": contactID,
		}
	}

	reportResponseBytes, err := provider.Find(session, "Reports/AgedPayablesByContact", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunAgedReceivablesByContact will run the Aged Receivables By Contact Report and marshal the results to a Report Struct
//Date, FromDate and ToDate can be added as optional paramters as a map
func RunAgedReceivablesByContact(provider *xero.Provider, session goth.Session, contactID string, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if querystringParameters != nil {
		querystringParameters["ContactID"] = contactID
	} else {
		querystringParameters = map[string]string{
			"ContactID": contactID,
		}
	}

	reportResponseBytes, err := provider.Find(session, "Reports/AgedReceivablesByContact", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunBalanceSheet will run the Balance Sheet Report and marshal the results to a Report Struct
//date, trackingOptionID1, trackingOptionID2, standardLayout, and paymentsOnly can be added as optional paramters as a map
func RunBalanceSheet(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/BalanceSheet", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunBankStatement will run the Bank Statement Report and marshal the results to a Report Struct
//FromDate and ToDate can be added as optional paramters as a map
func RunBankStatement(provider *xero.Provider, session goth.Session, bankAccountID string, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if querystringParameters != nil {
		querystringParameters["bankAccountID"] = bankAccountID
	} else {
		querystringParameters = map[string]string{
			"bankAccountID": bankAccountID,
		}
	}

	reportResponseBytes, err := provider.Find(session, "Reports/BankStatement", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunBankSummary will run the Bank Summary Report and marshal the results to a Report Struct
//FromDate and ToDate can be added as optional paramters as a map
func RunBankSummary(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/BankSummary", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunBASReport will retrieve an individual BAS Report given a reportID and marshal the results to a Report Struct
//Will only work for AU based Organisations
func RunBASReport(provider *xero.Provider, session goth.Session, reportID string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/"+reportID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunBASReports will retrieve all BAS Reports and marshal the results to a Report Struct
//Will only work for AU based Organisations
func RunBASReports(provider *xero.Provider, session goth.Session) (*Reports, error) {
	return RunBASReport(provider, session, "")
}

//RunBudgetSummary will run the Budget Summary Report and marshal the results to a Report Struct
func RunBudgetSummary(provider *xero.Provider, session goth.Session) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/BudgetSummary", additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunExecutiveSummary will run the Executive Summary Report and marshal the results to a Report Struct
//date can be added as an optional paramter as a map
func RunExecutiveSummary(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/ExecutiveSummary", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunGSTReport will retrieve an individual GST Report given a reportID and marshal the results to a Report Struct
//Will only work for NZ based Organisations
func RunGSTReport(provider *xero.Provider, session goth.Session, reportID string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/"+reportID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunGSTReports will retrieve all GST Reports and marshal the results to a Report Struct
//Will only work for NZ based Organisations
func RunGSTReports(provider *xero.Provider, session goth.Session) (*Reports, error) {
	return RunGSTReport(provider, session, "")
}

//RunProfitAndLoss will run the Profit And Loss Report and marshal the results to a Report Struct
//date, trackingCategoryID, trackingOptionID, trackingCategoryID2, trackingOptionID2,
//standardLayout, and paymentsOnly can be added as optional paramters as a map
func RunProfitAndLoss(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/ProfitAndLoss", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

//RunTrialBalance will run the TrialBalance Report and marshal the results to a Report Struct
//date and paymentsOnly can be added as optional paramters as a map
func RunTrialBalance(provider *xero.Provider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(session, "Reports/TrialBalance", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}
