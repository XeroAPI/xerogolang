package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"time"

	"math"

	"github.com/XeroAPI/xerogolang"
	"github.com/XeroAPI/xerogolang/accounting"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

var (
	provider           = xerogolang.New(os.Getenv("XERO_KEY"), os.Getenv("XERO_SECRET"), "http://localhost:3000/auth/callback?provider=xero")
	store              = sessions.NewFilesystemStore(os.TempDir(), []byte("xero-example"))
	invoices           = new(accounting.Invoices)
	contacts           = new(accounting.Contacts)
	accounts           = new(accounting.Accounts)
	bankTransactions   = new(accounting.BankTransactions)
	creditNotes        = new(accounting.CreditNotes)
	contactGroups      = new(accounting.ContactGroups)
	currencies         = new(accounting.Currencies)
	items              = new(accounting.Items)
	journals           = new(accounting.Journals)
	manualJournals     = new(accounting.ManualJournals)
	payments           = new(accounting.Payments)
	purchaseOrders     = new(accounting.PurchaseOrders)
	trackingCategories = new(accounting.TrackingCategories)
	taxRates           = new(accounting.TaxRates)
	receipts           = new(accounting.Receipts)
	bankTransfers      = new(accounting.BankTransfers)
)

func init() {
	goth.UseProviders(provider)

	store.MaxLength(math.MaxInt64)

	gothic.Store = store
}

//indexHandler dictates what is processed on the index route
func indexHandler(res http.ResponseWriter, req *http.Request) {
	session, err := provider.GetSessionFromStore(req, res)
	if err != nil || session == nil {
		t, _ := template.New("foo").Parse(indexNotConnectedTemplate)
		t.Execute(res, nil)
	} else {
		organisationCollection, err := accounting.FindOrganisation(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(indexConnectedTemplate)
		t.Execute(res, organisationCollection.Organisations[0])
	}
}

//authHandler dictates what is processed on the auth route
func authHandler(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		t, _ := template.New("foo").Parse(connectTemplate)
		t.Execute(res, gothUser)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

//callbackHandler dictates what is processed on the auth/callback route
func callbackHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.New("foo").Parse(connectTemplate)
	t.Execute(res, user)
}

//createHandler dictates what is processed on the create route
func createHandler(res http.ResponseWriter, req *http.Request) {
	session, err := provider.GetSessionFromStore(req, res)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	vars := mux.Vars(req)
	object := vars["object"]
	switch object {
	case "invoice":
		invoices = accounting.GenerateExampleInvoice()
		invoiceCollection, err := invoices.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		invoices = invoiceCollection
		t, _ := template.New("foo").Parse(invoiceTemplate)
		t.Execute(res, invoiceCollection.Invoices[0])
	case "contact":
		contacts = accounting.GenerateExampleContact()
		contactCollection, err := contacts.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		contacts = contactCollection
		t, _ := template.New("foo").Parse(contactTemplate)
		t.Execute(res, contactCollection.Contacts[0])
	case "account":
		accounts = accounting.GenerateExampleAccount()
		accountCollection, err := accounts.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		accounts = accountCollection
		t, _ := template.New("foo").Parse(accountTemplate)
		t.Execute(res, accountCollection.Accounts[0])
	case "banktransaction":
		bankTransactions = accounting.GenerateExampleBankTransaction()
		bankTransactionCollection, err := bankTransactions.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		bankTransactions = bankTransactionCollection
		t, _ := template.New("foo").Parse(bankTransactionTemplate)
		t.Execute(res, bankTransactionCollection.BankTransactions[0])
	case "creditnote":
		creditNotes = accounting.GenerateExampleCreditNote()
		creditNoteCollection, err := creditNotes.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		creditNotes = creditNoteCollection
		t, _ := template.New("foo").Parse(creditNoteTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes[0])
	case "contactgroup":
		contactGroups = accounting.GenerateExampleContactGroup()
		contactGroupCollection, err := contactGroups.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		contactGroups = contactGroupCollection
		t, _ := template.New("foo").Parse(contactGroupTemplate)
		t.Execute(res, contactGroupCollection.ContactGroups[0])
	case "item":
		items = accounting.GenerateExampleItem()
		itemCollection, err := items.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		items = itemCollection
		t, _ := template.New("foo").Parse(itemTemplate)
		t.Execute(res, itemCollection.Items[0])
	case "manualjournal":
		manualJournals = accounting.GenerateExampleManualJournal()
		manualJournalCollection, err := manualJournals.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		manualJournals = manualJournalCollection
		t, _ := template.New("foo").Parse(manualJournalTemplate)
		t.Execute(res, manualJournalCollection.ManualJournals[0])
	case "purchaseorder":
		purchaseOrders = accounting.GenerateExamplePurchaseOrder("")
		purchaseOrderCollection, err := purchaseOrders.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		purchaseOrders = purchaseOrderCollection
		t, _ := template.New("foo").Parse(purchaseOrderTemplate)
		t.Execute(res, purchaseOrderCollection.PurchaseOrders[0])
	case "trackingcategory":
		trackingCategories = accounting.GenerateExampleTrackingCategory()
		trackingCategoryCollection, err := trackingCategories.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		trackingCategories = trackingCategoryCollection
		t, _ := template.New("foo").Parse(trackingCategoryTemplate)
		t.Execute(res, trackingCategoryCollection.TrackingCategories[0])
	case "taxrate":
		taxRates = accounting.GenerateExampleTaxRate()
		taxRateCollection, err := taxRates.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		taxRates = taxRateCollection
		t, _ := template.New("foo").Parse(taxRateTemplate)
		t.Execute(res, taxRateCollection.TaxRates[0])
	case "banktransfer":
		bankTransfers = accounting.GenerateExampleBankTransfer()
		bankTransferCollection, err := bankTransfers.Create(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		bankTransfers = bankTransferCollection
		t, _ := template.New("foo").Parse(bankTransferTemplate)
		t.Execute(res, bankTransferCollection.BankTransfers[0])
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

//disconnectHandler dictates what is processed on the disconnect route
func disconnectHandler(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

//findHandler dictates what is processed on the find route
func findHandler(res http.ResponseWriter, req *http.Request) {
	session, err := provider.GetSessionFromStore(req, res)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	vars := mux.Vars(req)
	object := vars["object"]
	id := vars["id"]
	switch object {
	case "invoice":
		invoiceCollection, err := accounting.FindInvoice(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		invoices = invoiceCollection

		t, _ := template.New("foo").Parse(invoiceTemplate)
		t.Execute(res, invoiceCollection.Invoices[0])
	case "contact":
		contactCollection, err := accounting.FindContact(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		contacts = contactCollection

		t, _ := template.New("foo").Parse(contactTemplate)
		t.Execute(res, contactCollection.Contacts[0])
	case "account":
		accountCollection, err := accounting.FindAccount(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		accounts = accountCollection

		t, _ := template.New("foo").Parse(accountTemplate)
		t.Execute(res, accountCollection.Accounts[0])
	case "banktransaction":
		bankTransactionCollection, err := accounting.FindBankTransaction(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		bankTransactions = bankTransactionCollection

		t, _ := template.New("foo").Parse(bankTransactionTemplate)
		t.Execute(res, bankTransactionCollection.BankTransactions[0])
	case "creditnote":
		creditNoteCollection, err := accounting.FindCreditNote(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		creditNotes = creditNoteCollection

		t, _ := template.New("foo").Parse(creditNoteTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes[0])
	case "contactgroup":
		contactGroupCollection, err := accounting.FindContactGroup(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		contactGroups = contactGroupCollection

		t, _ := template.New("foo").Parse(contactGroupTemplate)
		t.Execute(res, contactGroupCollection.ContactGroups[0])
	case "item":
		itemCollection, err := accounting.FindItem(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		items = itemCollection

		t, _ := template.New("foo").Parse(itemTemplate)
		t.Execute(res, itemCollection.Items[0])
	case "journal":
		journalCollection, err := accounting.FindJournal(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		journals = journalCollection

		t, _ := template.New("foo").Parse(journalTemplate)
		t.Execute(res, journalCollection.Journals[0])
	case "manualjournal":
		manualJournalCollection, err := accounting.FindManualJournal(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		manualJournals = manualJournalCollection

		t, _ := template.New("foo").Parse(manualJournalTemplate)
		t.Execute(res, manualJournalCollection.ManualJournals[0])
	case "payment":
		paymentCollection, err := accounting.FindPayment(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		payments = paymentCollection

		t, _ := template.New("foo").Parse(paymentTemplate)
		t.Execute(res, paymentCollection.Payments[0])
	case "purchaseorder":
		purchaseOrderCollection, err := accounting.FindPurchaseOrder(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		purchaseOrders = purchaseOrderCollection

		t, _ := template.New("foo").Parse(purchaseOrderTemplate)
		t.Execute(res, purchaseOrderCollection.PurchaseOrders[0])
	case "trackingcategory":
		trackingCategoryCollection, err := accounting.FindTrackingCategory(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		trackingCategories = trackingCategoryCollection

		t, _ := template.New("foo").Parse(trackingCategoryTemplate)
		t.Execute(res, trackingCategoryCollection.TrackingCategories[0])
	case "overpayment":
		overpaymentCollection, err := accounting.FindOverpayment(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(overpaymentTemplate)
		t.Execute(res, overpaymentCollection.Overpayments[0])
	case "prepayment":
		prepaymentCollection, err := accounting.FindPrepayment(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(prepaymentTemplate)
		t.Execute(res, prepaymentCollection.Prepayments[0])
	case "agedpayablesbycontact":
		agedPayablesCollection, err := accounting.RunAgedPayablesByContact(provider, session, id, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, agedPayablesCollection.Reports[0])
	case "agedreceivablesbycontact":
		agedReceivablesCollection, err := accounting.RunAgedReceivablesByContact(provider, session, id, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, agedReceivablesCollection.Reports[0])
	case "balancesheet":
		balanceSheetCollection, err := accounting.RunBalanceSheet(provider, session, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, balanceSheetCollection.Reports[0])
	case "banksummary":
		bankSummaryCollection, err := accounting.RunBankSummary(provider, session, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, bankSummaryCollection.Reports[0])
	case "budgetsummary":
		budgetSummaryCollection, err := accounting.RunBudgetSummary(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, budgetSummaryCollection.Reports[0])
	case "executivesummary":
		executiveSummaryCollection, err := accounting.RunExecutiveSummary(provider, session, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, executiveSummaryCollection.Reports[0])
	case "profitandloss":
		profitAndLossCollection, err := accounting.RunProfitAndLoss(provider, session, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, profitAndLossCollection.Reports[0])
	case "trialbalance":
		trialBalanceCollection, err := accounting.RunTrialBalance(provider, session, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(reportTemplate)
		t.Execute(res, trialBalanceCollection.Reports[0])
	case "linkedtransaction":
		linkedTransactionCollection, err := accounting.FindLinkedTransaction(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(linkedTransactionTemplate)
		t.Execute(res, linkedTransactionCollection.LinkedTransactions[0])
	case "user":
		userCollection, err := accounting.FindUser(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, userCollection.Users[0])
	case "expenseclaim":
		expenseClaimCollection, err := accounting.FindExpenseClaim(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(expenseClaimTemplate)
		t.Execute(res, expenseClaimCollection.ExpenseClaims[0])
	case "receipt":
		receiptCollection, err := accounting.FindReceipt(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		receipts = receiptCollection

		t, _ := template.New("foo").Parse(receiptTemplate)
		t.Execute(res, receiptCollection.Receipts[0])
	case "repeatinginvoice":
		repeatingInvoiceCollection, err := accounting.FindRepeatingInvoice(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		t, _ := template.New("foo").Parse(repeatingInvoiceTemplate)
		t.Execute(res, repeatingInvoiceCollection.RepeatingInvoices[0])
	case "banktransfer":
		bankTransferCollection, err := accounting.FindBankTransfer(provider, session, id)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		bankTransfers = bankTransferCollection

		t, _ := template.New("foo").Parse(bankTransferTemplate)
		t.Execute(res, bankTransferCollection.BankTransfers[0])
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

//findAllHandler dictates what is processed on the findall route
func findAllHandler(res http.ResponseWriter, req *http.Request) {
	session, err := provider.GetSessionFromStore(req, res)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	vars := mux.Vars(req)
	object := vars["object"]

	modifiedSince := req.URL.Query().Get("modifiedsince")
	modifiedSince, err = url.QueryUnescape(modifiedSince)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	switch object {
	case "invoices":
		invoiceCollection := new(accounting.Invoices)
		var err error
		if modifiedSince == "" {
			invoiceCollection, err = accounting.FindInvoices(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			invoiceCollection, err = accounting.FindInvoicesModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(invoicesTemplate)
		t.Execute(res, invoiceCollection.Invoices)
	case "contacts":
		contactCollection := new(accounting.Contacts)
		var err error
		if modifiedSince == "" {
			contactCollection, err = accounting.FindContacts(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			contactCollection, err = accounting.FindContactsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactsTemplate)
		t.Execute(res, contactCollection.Contacts)
	case "accounts":
		accountCollection := new(accounting.Accounts)
		var err error
		if modifiedSince == "" {
			accountCollection, err = accounting.FindAccounts(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			accountCollection, err = accounting.FindAccountsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(accountsTemplate)
		t.Execute(res, accountCollection.Accounts)
	case "banktransactions":
		bankTransactionCollection := new(accounting.BankTransactions)
		var err error
		if modifiedSince == "" {
			bankTransactionCollection, err = accounting.FindBankTransactions(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			bankTransactionCollection, err = accounting.FindBankTransactionsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(bankTransactionsTemplate)
		t.Execute(res, bankTransactionCollection.BankTransactions)
	case "creditnotes":
		creditNoteCollection := new(accounting.CreditNotes)
		var err error
		if modifiedSince == "" {
			creditNoteCollection, err = accounting.FindCreditNotes(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			creditNoteCollection, err = accounting.FindCreditNotesModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(creditNotesTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes)
	case "contactgroups":
		contactGroupCollection, err := accounting.FindContactGroups(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactGroupsTemplate)
		t.Execute(res, contactGroupCollection.ContactGroups)
	case "currencies":
		currencyCollection, err := accounting.FindCurrencies(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(currenciesTemplate)
		t.Execute(res, currencyCollection.Currencies)
	case "items":
		itemCollection := new(accounting.Items)
		var err error
		if modifiedSince == "" {
			itemCollection, err = accounting.FindItems(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			itemCollection, err = accounting.FindItemsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(itemsTemplate)
		t.Execute(res, itemCollection.Items)
	case "journals":
		journalCollection := new(accounting.Journals)
		var err error
		querystringParameters := map[string]string{
			"offset": "0",
		}
		if modifiedSince == "" {
			journalCollection, err = accounting.FindJournals(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			journalCollection, err = accounting.FindJournalsModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(journalsTemplate)
		t.Execute(res, journalCollection.Journals)
	case "manualjournals":
		manualJournalCollection := new(accounting.ManualJournals)
		var err error
		if modifiedSince == "" {
			manualJournalCollection, err = accounting.FindManualJournals(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			manualJournalCollection, err = accounting.FindManualJournalsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(manualJournalsTemplate)
		t.Execute(res, manualJournalCollection.ManualJournals)
	case "payments":
		paymentCollection := new(accounting.Payments)
		var err error
		if modifiedSince == "" {
			paymentCollection, err = accounting.FindPayments(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			paymentCollection, err = accounting.FindPaymentsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(paymentsTemplate)
		t.Execute(res, paymentCollection.Payments)
	case "purchaseorders":
		purchaseOrderCollection := new(accounting.PurchaseOrders)
		var err error
		if modifiedSince == "" {
			purchaseOrderCollection, err = accounting.FindPurchaseOrders(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			purchaseOrderCollection, err = accounting.FindPurchaseOrdersModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(purchaseOrdersTemplate)
		t.Execute(res, purchaseOrderCollection.PurchaseOrders)
	case "trackingcategories":
		trackingCategoryCollection, err := accounting.FindTrackingCategories(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(trackingCategoriesTemplate)
		t.Execute(res, trackingCategoryCollection.TrackingCategories)
	case "taxrates":
		taxRateCollection, err := accounting.FindTaxRates(provider, session, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(taxRatesTemplate)
		t.Execute(res, taxRateCollection.TaxRates)
	case "overpayments":
		overpaymentCollection := new(accounting.Overpayments)
		var err error
		if modifiedSince == "" {
			overpaymentCollection, err = accounting.FindOverpayments(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			overpaymentCollection, err = accounting.FindOverpaymentsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(overpaymentsTemplate)
		t.Execute(res, overpaymentCollection.Overpayments)
	case "prepayments":
		prepaymentCollection := new(accounting.Prepayments)
		var err error
		if modifiedSince == "" {
			prepaymentCollection, err = accounting.FindPrepayments(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			prepaymentCollection, err = accounting.FindPrepaymentsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(prepaymentsTemplate)
		t.Execute(res, prepaymentCollection.Prepayments)
	case "linkedtransactions":
		linkedTransactionCollection := new(accounting.LinkedTransactions)
		var err error
		if modifiedSince == "" {
			linkedTransactionCollection, err = accounting.FindLinkedTransactions(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			linkedTransactionCollection, err = accounting.FindLinkedTransactionsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(linkedTransactionsTemplate)
		t.Execute(res, linkedTransactionCollection.LinkedTransactions)
	case "users":
		userCollection := new(accounting.Users)
		var err error
		if modifiedSince == "" {
			userCollection, err = accounting.FindUsers(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			userCollection, err = accounting.FindUsersModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(usersTemplate)
		t.Execute(res, userCollection.Users)
	case "expenseclaims":
		expenseClaimCollection := new(accounting.ExpenseClaims)
		var err error
		if modifiedSince == "" {
			expenseClaimCollection, err = accounting.FindExpenseClaims(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			expenseClaimCollection, err = accounting.FindExpenseClaimsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(expenseClaimsTemplate)
		t.Execute(res, expenseClaimCollection.ExpenseClaims)
	case "receipts":
		receiptCollection := new(accounting.Receipts)
		var err error
		if modifiedSince == "" {
			receiptCollection, err = accounting.FindReceipts(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			receiptCollection, err = accounting.FindReceiptsModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(receiptsTemplate)
		t.Execute(res, receiptCollection.Receipts)
	case "repeatinginvoices":
		repeatingInvoiceCollection, err := accounting.FindRepeatingInvoices(provider, session, nil)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(repeatingInvoicesTemplate)
		t.Execute(res, repeatingInvoiceCollection.RepeatingInvoices)
	case "banktransfers":
		bankTransferCollection := new(accounting.BankTransfers)
		var err error
		if modifiedSince == "" {
			bankTransferCollection, err = accounting.FindBankTransfers(provider, session, nil)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			bankTransferCollection, err = accounting.FindBankTransfersModifiedSince(provider, session, parsedTime, nil)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(bankTransfersTemplate)
		t.Execute(res, bankTransferCollection.BankTransfers)
	case "brandingthemes":
		currencyCollection, err := accounting.FindBrandingThemes(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(brandingThemesTemplate)
		t.Execute(res, currencyCollection.BrandingThemes)
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

//findAllPagedHandler dictates what is processed on the findall/{object}/{page} route
func findAllPagedHandler(res http.ResponseWriter, req *http.Request) {
	session, err := provider.GetSessionFromStore(req, res)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	vars := mux.Vars(req)
	object := vars["object"]
	page := vars["page"]
	querystringParameters := map[string]string{
		"page": page,
	}
	modifiedSince := req.URL.Query().Get("modifiedsince")
	modifiedSince, err = url.QueryUnescape(modifiedSince)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	switch object {
	case "invoices":
		invoiceCollection := new(accounting.Invoices)
		var err error
		if modifiedSince == "" {
			invoiceCollection, err = accounting.FindInvoices(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			invoiceCollection, err = accounting.FindInvoicesModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(invoicesTemplate)
		t.Execute(res, invoiceCollection.Invoices)
	case "contacts":
		contactCollection := new(accounting.Contacts)
		var err error
		if modifiedSince == "" {
			contactCollection, err = accounting.FindContacts(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, err)
				return
			}
			contactCollection, err = accounting.FindContactsModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactsTemplate)
		t.Execute(res, contactCollection.Contacts)
	case "banktransactions":
		bankTransactionCollection := new(accounting.BankTransactions)
		var err error
		if modifiedSince == "" {
			bankTransactionCollection, err = accounting.FindBankTransactions(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			bankTransactionCollection, err = accounting.FindBankTransactionsModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(bankTransactionsTemplate)
		t.Execute(res, bankTransactionCollection.BankTransactions)
	case "creditnotes":
		creditNoteCollection := new(accounting.CreditNotes)
		var err error
		if modifiedSince == "" {
			creditNoteCollection, err = accounting.FindCreditNotes(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			creditNoteCollection, err = accounting.FindCreditNotesModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(creditNotesTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes)
	case "journals":
		journalParameters := map[string]string{
			"offset": page,
		}
		journalCollection := new(accounting.Journals)
		var err error
		if modifiedSince == "" {
			journalCollection, err = accounting.FindJournals(provider, session, journalParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, err)
				return
			}
			journalCollection, err = accounting.FindJournalsModifiedSince(provider, session, parsedTime, journalParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(journalsTemplate)
		t.Execute(res, journalCollection.Journals)
	case "manualjournals":
		manualJournalCollection := new(accounting.ManualJournals)
		var err error
		if modifiedSince == "" {
			manualJournalCollection, err = accounting.FindManualJournals(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			manualJournalCollection, err = accounting.FindManualJournalsModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(manualJournalsTemplate)
		t.Execute(res, manualJournalCollection.ManualJournals)
	case "purchaseorders":
		purchaseOrderCollection := new(accounting.PurchaseOrders)
		var err error
		if modifiedSince == "" {
			purchaseOrderCollection, err = accounting.FindPurchaseOrders(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			purchaseOrderCollection, err = accounting.FindPurchaseOrdersModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(purchaseOrdersTemplate)
		t.Execute(res, purchaseOrderCollection.PurchaseOrders)
	case "overpayments":
		overpaymentCollection := new(accounting.Overpayments)
		var err error
		if modifiedSince == "" {
			overpaymentCollection, err = accounting.FindOverpayments(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			overpaymentCollection, err = accounting.FindOverpaymentsModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(overpaymentsTemplate)
		t.Execute(res, overpaymentCollection.Overpayments)
	case "prepayments":
		prepaymentCollection := new(accounting.Prepayments)
		var err error
		if modifiedSince == "" {
			prepaymentCollection, err = accounting.FindPrepayments(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			prepaymentCollection, err = accounting.FindPrepaymentsModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(prepaymentsTemplate)
		t.Execute(res, prepaymentCollection.Prepayments)
	case "linkedtransactions":
		linkedTransactionCollection := new(accounting.LinkedTransactions)
		var err error
		if modifiedSince == "" {
			linkedTransactionCollection, err = accounting.FindLinkedTransactions(provider, session, querystringParameters)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			linkedTransactionCollection, err = accounting.FindLinkedTransactionsModifiedSince(provider, session, parsedTime, querystringParameters)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(linkedTransactionsTemplate)
		t.Execute(res, linkedTransactionCollection.LinkedTransactions)
	default:
		fmt.Fprintln(res, "Paging not supported on object specified")
		return
	}
}

//findWhereHandler dictates what is processed on the findwhere route
func findWhereHandler(res http.ResponseWriter, req *http.Request) {
	session, err := provider.GetSessionFromStore(req, res)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	vars := mux.Vars(req)
	object := vars["object"]

	whereClause := req.URL.Query().Get("where")
	whereClause, err = url.QueryUnescape(whereClause)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	querystringParameters := map[string]string{
		"where": whereClause,
	}
	switch object {
	case "invoices":
		invoiceCollection, err := accounting.FindInvoices(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(invoicesTemplate)
		t.Execute(res, invoiceCollection.Invoices)
	case "contacts":
		contactCollection, err := accounting.FindContacts(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactsTemplate)
		t.Execute(res, contactCollection.Contacts)
	case "banktransactions":
		bankTransactionCollection, err := accounting.FindBankTransactions(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(bankTransactionsTemplate)
		t.Execute(res, bankTransactionCollection.BankTransactions)
	case "creditnotes":
		creditNoteCollection, err := accounting.FindCreditNotes(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(creditNotesTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes)
	case "items":
		itemCollection, err := accounting.FindItems(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(itemsTemplate)
		t.Execute(res, itemCollection.Items)
	case "manualjournals":
		manualJournalCollection, err := accounting.FindManualJournals(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(manualJournalsTemplate)
		t.Execute(res, manualJournalCollection.ManualJournals)
	case "payments":
		paymentCollection, err := accounting.FindPayments(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(paymentsTemplate)
		t.Execute(res, paymentCollection.Payments)
	case "overpayments":
		overpaymentCollection, err := accounting.FindOverpayments(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(overpaymentsTemplate)
		t.Execute(res, overpaymentCollection.Overpayments)
	case "prepayments":
		prepaymentCollection, err := accounting.FindPrepayments(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(prepaymentsTemplate)
		t.Execute(res, prepaymentCollection.Prepayments)
	case "users":
		userCollection, err := accounting.FindUsers(provider, session, querystringParameters)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(usersTemplate)
		t.Execute(res, userCollection.Users)
	default:
		fmt.Fprintln(res, "Where clauses not available on this entity")
		return
	}
}

//updateHandler dictates what is processed on the update route
func updateHandler(res http.ResponseWriter, req *http.Request) {
	session, err := provider.GetSessionFromStore(req, res)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	vars := mux.Vars(req)
	object := vars["object"]
	id := vars["id"]

	switch object {
	case "invoice":
		if id != invoices.Invoices[0].InvoiceID {
			fmt.Fprintln(res, "Could not update Invoice")
			return
		}
		if invoices.Invoices[0].Status == "DRAFT" {
			invoices.Invoices[0].Status = "SUBMITTED"
		} else if invoices.Invoices[0].Status == "SUBMITTED" {
			invoices.Invoices[0].Status = "DRAFT"
		}

		invoiceCollection, err := invoices.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(invoiceTemplate)
		t.Execute(res, invoiceCollection.Invoices[0])
	case "contact":
		if id != contacts.Contacts[0].ContactID {
			fmt.Fprintln(res, "Could not update Contact")
			return
		}
		if contacts.Contacts[0].EmailAddress == "" || contacts.Contacts[0].EmailAddress == "it@shrinks.com" {
			contacts.Contacts[0].EmailAddress = "serenity@now.com"
		} else if contacts.Contacts[0].EmailAddress == "serenity@now.com" {
			contacts.Contacts[0].EmailAddress = "it@shrinks.com"
		}

		contactCollection, err := contacts.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactTemplate)
		t.Execute(res, contactCollection.Contacts[0])
	case "account":
		if id != accounts.Accounts[0].AccountID {
			fmt.Fprintln(res, "Could not update Account")
			return
		}
		if accounts.Accounts[0].EnablePaymentsToAccount == false {
			accounts.Accounts[0].Status = ""
			accounts.Accounts[0].EnablePaymentsToAccount = true
		} else if accounts.Accounts[0].EnablePaymentsToAccount == true {
			accounts.Accounts[0].Status = ""
			accounts.Accounts[0].EnablePaymentsToAccount = false
		}

		accountCollection, err := accounts.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(accountTemplate)
		t.Execute(res, accountCollection.Accounts[0])
	case "banktransaction":
		if id != bankTransactions.BankTransactions[0].BankTransactionID {
			fmt.Fprintln(res, "Could not update BankTransaction")
			return
		}
		if bankTransactions.BankTransactions[0].Status == "AUTHORISED" {
			bankTransactions.BankTransactions[0].Status = "DELETED"
		}

		bankTransactionCollection, err := bankTransactions.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(bankTransactionTemplate)
		t.Execute(res, bankTransactionCollection.BankTransactions[0])
	case "creditnote":
		if id != creditNotes.CreditNotes[0].CreditNoteID {
			fmt.Fprintln(res, "Could not update CreditNote")
			return
		}
		if creditNotes.CreditNotes[0].Status == "DRAFT" {
			creditNotes.CreditNotes[0].Status = "SUBMITTED"
		} else if creditNotes.CreditNotes[0].Status == "SUBMITTED" {
			creditNotes.CreditNotes[0].Status = "DRAFT"
		}

		creditNoteCollection, err := creditNotes.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(creditNoteTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes[0])
	case "contactgroup":
		if id != contactGroups.ContactGroups[0].ContactGroupID {
			fmt.Fprintln(res, "Could not update ContactGroup")
			return
		}
		if contactGroups.ContactGroups[0].Status == "ACTIVE" {
			contactGroups.ContactGroups[0].Status = "DELETED"
		}

		contactGroupCollection, err := contactGroups.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactGroupTemplate)
		t.Execute(res, contactGroupCollection.ContactGroups[0])
	case "item":
		if id != items.Items[0].ItemID {
			fmt.Fprintln(res, "Could not update Item")
			return
		}
		if items.Items[0].Description == "A Beltless Trenchcoat" {
			items.Items[0].Description = "The beltless trench-coat"
		} else if items.Items[0].Description == "The beltless trench-coat" {
			items.Items[0].Description = "A Beltless Trenchcoat"
		}

		itemCollection, err := items.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(itemTemplate)
		t.Execute(res, itemCollection.Items[0])
	case "manualjournal":
		if id != manualJournals.ManualJournals[0].ManualJournalID {
			fmt.Fprintln(res, "Could not update ManualJournal")
			return
		}
		if manualJournals.ManualJournals[0].Status == "DRAFT" {
			manualJournals.ManualJournals[0].Status = "POSTED"
		} else if manualJournals.ManualJournals[0].Status == "POSTED" {
			manualJournals.ManualJournals[0].Status = "DRAFT"
		}

		manualJournalCollection, err := manualJournals.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(manualJournalTemplate)
		t.Execute(res, manualJournalCollection.ManualJournals[0])
	case "payment":
		if id != payments.Payments[0].PaymentID {
			fmt.Fprintln(res, "Could not update Payment")
			return
		}
		if payments.Payments[0].Status == "AUTHORISED" {
			payments.Payments[0].Status = "DELETED"
		}

		paymentCollection, err := payments.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(paymentTemplate)
		t.Execute(res, paymentCollection.Payments[0])
	case "purchaseorder":
		if id != purchaseOrders.PurchaseOrders[0].PurchaseOrderID {
			fmt.Fprintln(res, "Could not update PurchaseOrder")
			return
		}
		if purchaseOrders.PurchaseOrders[0].Status == "DRAFT" {
			purchaseOrders.PurchaseOrders[0].Status = "SUBMITTED"
		} else if purchaseOrders.PurchaseOrders[0].Status == "SUBMITTED" {
			purchaseOrders.PurchaseOrders[0].Status = "DRAFT"
		}

		purchaseOrderCollection, err := purchaseOrders.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(purchaseOrderTemplate)
		t.Execute(res, purchaseOrderCollection.PurchaseOrders[0])
	case "trackingcategory":
		if id != trackingCategories.TrackingCategories[0].TrackingCategoryID {
			fmt.Fprintln(res, "Could not update TrackingCategory")
			return
		}
		if trackingCategories.TrackingCategories[0].Name == "Person Responsible" {
			trackingCategories.TrackingCategories[0].Name = "Manager"
		} else if trackingCategories.TrackingCategories[0].Name == "Manager" {
			trackingCategories.TrackingCategories[0].Name = "Person Responsible"
		}

		trackingCategoryCollection, err := trackingCategories.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(trackingCategoryTemplate)
		t.Execute(res, trackingCategoryCollection.TrackingCategories[0])
	case "taxrate":
		if id != taxRates.TaxRates[0].Name {
			fmt.Fprintln(res, "Could not update TaxRate")
			return
		}
		if taxRates.TaxRates[0].Status == "ACTIVE" {
			taxRates.TaxRates[0].Status = "DELETED"
		}

		taxRateCollection, err := taxRates.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(taxRateTemplate)
		t.Execute(res, taxRateCollection.TaxRates[0])
	case "receipt":
		if id != receipts.Receipts[0].ReceiptID {
			fmt.Fprintln(res, "Could not update Receipt")
			return
		}
		if receipts.Receipts[0].Reference == "1111" || receipts.Receipts[0].Reference == "" {
			receipts.Receipts[0].Reference = "2222"
		} else if receipts.Receipts[0].Reference == "2222" {
			receipts.Receipts[0].Reference = "1111"
		}

		receiptCollection, err := receipts.Update(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(receiptTemplate)
		t.Execute(res, receiptCollection.Receipts[0])
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	a := r.PathPrefix("/auth").Subrouter()
	// "/auth/"
	a.HandleFunc("/", authHandler).Methods("GET")
	// "/auth/callback"
	a.HandleFunc("/callback", callbackHandler).Methods("GET")
	c := r.PathPrefix("/create").Subrouter()
	// "/create/{object}"
	c.HandleFunc("/{object}", createHandler).Methods("GET")
	//"/disconnect"
	r.HandleFunc("/disconnect", disconnectHandler).Methods("GET")
	f := r.PathPrefix("/find").Subrouter()
	// "/find/{object}/id"
	f.HandleFunc("/{object}/{id}", findHandler).Methods("GET")
	fa := r.PathPrefix("/findall").Subrouter()
	// "/findall/{object}"
	fa.HandleFunc("/{object}", findAllHandler).Methods("GET")
	// "/findall/{object}/{page}"
	fa.HandleFunc("/{object}/{page}", findAllPagedHandler).Methods("GET")
	fw := r.PathPrefix("/findwhere").Subrouter()
	// "/findwhere/{object}"
	fw.HandleFunc("/{object}", findWhereHandler).Methods("GET")
	u := r.PathPrefix("/update").Subrouter()
	// "/update/{object}/id"
	u.HandleFunc("/{object}/{id}", updateHandler).Methods("GET")
	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)
}
