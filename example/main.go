package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"math"

	"github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/accounting"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

var (
	provider         = xero.New(os.Getenv("XERO_KEY"), os.Getenv("XERO_SECRET"), "http://localhost:3000/auth/callback?provider=xero")
	store            = sessions.NewFilesystemStore(os.TempDir(), []byte("xero-example"))
	invoices         = new(accounting.Invoices)
	contacts         = new(accounting.Contacts)
	accounts         = new(accounting.Accounts)
	bankTransactions = new(accounting.BankTransactions)
	creditNotes      = new(accounting.CreditNotes)
	contactGroups    = new(accounting.ContactGroups)
	currencies       = new(accounting.Currencies)
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
		t, _ := template.New("foo").Parse(userTemplate)
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
	t, _ := template.New("foo").Parse(userTemplate)
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
		invoices = accounting.CreateExampleInvoice()
		invoiceCollection, err := invoices.CreateInvoice(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		invoices = invoiceCollection
		t, _ := template.New("foo").Parse(invoiceTemplate)
		t.Execute(res, invoiceCollection.Invoices[0])
	case "contact":
		contacts = accounting.CreateExampleContact()
		contactCollection, err := contacts.CreateContact(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		contacts = contactCollection
		t, _ := template.New("foo").Parse(contactTemplate)
		t.Execute(res, contactCollection.Contacts[0])
	case "account":
		accounts = accounting.CreateExampleAccount()
		accountCollection, err := accounts.CreateAccount(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		accounts = accountCollection
		t, _ := template.New("foo").Parse(accountTemplate)
		t.Execute(res, accountCollection.Accounts[0])
	case "banktransaction":
		bankTransactions = accounting.CreateExampleBankTransaction()
		bankTransactionCollection, err := bankTransactions.CreateBankTransaction(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		bankTransactions = bankTransactionCollection
		t, _ := template.New("foo").Parse(bankTransactionTemplate)
		t.Execute(res, bankTransactionCollection.BankTransactions[0])
	case "creditnote":
		creditNotes = accounting.CreateExampleCreditNote()
		creditNoteCollection, err := creditNotes.CreateCreditNote(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		creditNotes = creditNoteCollection
		t, _ := template.New("foo").Parse(creditNoteTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes[0])
	case "contactgroup":
		contactGroups = accounting.CreateExampleContactGroup()
		contactGroupCollection, err := contactGroups.CreateContactGroup(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		contactGroups = contactGroupCollection
		t, _ := template.New("foo").Parse(contactGroupTemplate)
		t.Execute(res, contactGroupCollection.ContactGroups[0])
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
			invoiceCollection, err = accounting.FindInvoices(provider, session)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			invoiceCollection, err = accounting.FindInvoicesModifiedSince(provider, session, parsedTime)
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
			contactCollection, err = accounting.FindContacts(provider, session)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			contactCollection, err = accounting.FindContactsModifiedSince(provider, session, parsedTime)
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
			accountCollection, err = accounting.FindAccounts(provider, session)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			accountCollection, err = accounting.FindAccountsModifiedSince(provider, session, parsedTime)
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
			bankTransactionCollection, err = accounting.FindBankTransactions(provider, session)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			bankTransactionCollection, err = accounting.FindBankTransactionsModifiedSince(provider, session, parsedTime)
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
			creditNoteCollection, err = accounting.FindCreditNotes(provider, session)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			creditNoteCollection, err = accounting.FindCreditNotesModifiedSince(provider, session, parsedTime)
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
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		fmt.Fprintln(res, err)
		return
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
			invoiceCollection, err = accounting.FindInvoicesByPage(provider, session, pageInt)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			invoiceCollection, err = accounting.FindInvoicesModifiedSinceByPage(provider, session, parsedTime, pageInt)
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
			contactCollection, err = accounting.FindContactsByPage(provider, session, pageInt)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, err)
				return
			}
			contactCollection, err = accounting.FindContactsModifiedSinceByPage(provider, session, parsedTime, pageInt)
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
			bankTransactionCollection, err = accounting.FindBankTransactionsByPage(provider, session, pageInt)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			bankTransactionCollection, err = accounting.FindBankTransactionsModifiedSinceByPage(provider, session, parsedTime, pageInt)
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
			creditNoteCollection, err = accounting.FindCreditNotesByPage(provider, session, pageInt)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			creditNoteCollection, err = accounting.FindCreditNotesModifiedSinceByPage(provider, session, parsedTime, pageInt)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(creditNotesTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes)
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
	switch object {
	case "invoices":
		invoiceCollection := new(accounting.Invoices)
		var err error
		if whereClause == "" {
			invoiceCollection, err = accounting.FindInvoices(provider, session)
		} else {
			invoiceCollection, err = accounting.FindInvoicesWhere(provider, session, whereClause)
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
		if whereClause == "" {
			contactCollection, err = accounting.FindContacts(provider, session)
		} else {
			contactCollection, err = accounting.FindContactsWhere(provider, session, whereClause)
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
		if whereClause == "" {
			bankTransactionCollection, err = accounting.FindBankTransactions(provider, session)
		} else {
			bankTransactionCollection, err = accounting.FindBankTransactionsWhere(provider, session, whereClause)
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
		if whereClause == "" {
			creditNoteCollection, err = accounting.FindCreditNotes(provider, session)
		} else {
			creditNoteCollection, err = accounting.FindCreditNotesWhere(provider, session, whereClause)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(creditNotesTemplate)
		t.Execute(res, creditNoteCollection.CreditNotes)
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

		invoiceCollection, err := invoices.UpdateInvoice(provider, session)
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

		contactCollection, err := contacts.UpdateContact(provider, session)
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

		accountCollection, err := accounts.UpdateAccount(provider, session)
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

		bankTransactionCollection, err := bankTransactions.UpdateBankTransaction(provider, session)
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

		creditNoteCollection, err := creditNotes.UpdateCreditNote(provider, session)
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

		contactGroupCollection, err := contactGroups.UpdateContactGroup(provider, session)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactGroupTemplate)
		t.Execute(res, contactGroupCollection.ContactGroups[0])
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

var indexNotConnectedTemplate = `<p>
		<a href="/auth/?provider=xero">
			<img src="https://developer.xero.com/static/images/documentation/connect_xero_button_blue_2x.png" alt="ConnectToXero">
		</a>
	</p>`

var indexConnectedTemplate = `
<p><a href="/disconnect?provider=xero">Disconnect</a></p>
<p>Connected to: {{.Name}}</p>
<p>Actions:</p>
<p><a href="/create/invoice?provider=xero">create invoice</a></p>
<p><a href="/findall/invoices?provider=xero">find all invoices</a></p>
<p><a href="/findall/invoices?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all invoices changed since 1 May 2017</a></p>
<p><a href="/findall/invoices/1?provider=xero">find the first 100 invoices</a></p>
<p><a href="/create/contact?provider=xero">create contact</a></p>
<p><a href="/findall/contacts?provider=xero">find all contacts</a></p>
<p><a href="/findall/contacts?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all contacts changed since 1 May 2017</a></p>
<p><a href="/findall/contacts/1?provider=xero&page=1">find the first 100 contacts</a></p>
<p><a href="/create/account?provider=xero">create account</a></p>
<p><a href="/findall/accounts?provider=xero">find all accounts</a></p>
<p><a href="/findall/accounts?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all accounts changed since 1 May 2017</a></p>
<p><a href="/create/banktransaction?provider=xero">create bank transaction</a></p>
<p><a href="/findall/banktransactions?provider=xero">find all bank transactions</a></p>
<p><a href="/findall/banktransactions?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all bank transactions changed since 1 May 2017</a></p>
<p><a href="/findall/banktransactions/1?provider=xero">find the first 100 bank transactions</a></p>
<p><a href="/create/creditnote?provider=xero">create credit note</a></p>
<p><a href="/findall/creditnotes?provider=xero">find all credit notes</a></p>
<p><a href="/findall/creditnotes?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all credit notes changed since 1 May 2017</a></p>
<p><a href="/findall/creditnotes/1?provider=xero">find the first 100 credit notes</a></p>
<p><a href="/create/contactgroup?provider=xero">create contact group</a></p>
<p><a href="/findall/contactgroups?provider=xero">find all contact groups</a></p>
<p><a href="/findall/currencies?provider=xero">find all currencies</a></p>
`

var userTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>Connected Successfully!</p>
<p>Method: {{.Email}}</p>
<p>Org Name: {{.Name}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p><a href="/create/invoice?provider=xero">create invoice</a></p>
<p><a href="/findall/invoices?provider=xero">find all invoices</a></p>
<p><a href="/findall/invoices?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all invoices changed since 1 May 2017</a></p>
<p><a href="/findall/invoices/1?provider=xero">find the first 100 invoices</a></p>
<p><a href="/create/contact?provider=xero">create contact</a></p>
<p><a href="/findall/contacts?provider=xero">find all contacts</a></p>
<p><a href="/findall/contacts?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all contacts changed since 1 May 2017</a></p>
<p><a href="/findall/contacts/1?provider=xero">find the first 100 contacts</a></p>
<p><a href="/create/account?provider=xero">create account</a></p>
<p><a href="/findall/accounts?provider=xero">find all accounts</a></p>
<p><a href="/findall/accounts?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all accounts changed since 1 May 2017</a></p>
<p><a href="/create/banktransaction?provider=xero">create bank transaction</a></p>
<p><a href="/findall/banktransactions?provider=xero">find all bank transactions</a></p>
<p><a href="/findall/banktransactions?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all bank transactions changed since 1 May 2017</a></p>
<p><a href="/findall/banktransactions/1?provider=xero">find the first 100 bank transactions</a></p>
<p><a href="/create/creditnote?provider=xero">create credit note</a></p>
<p><a href="/findall/creditnotes?provider=xero">find all credit notes</a></p>
<p><a href="/findall/creditnotes?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all credit notes changed since 1 May 2017</a></p>
<p><a href="/findall/creditnotes/1?provider=xero">find the first 100 credit notes</a></p>
<p><a href="/create/contactgroup?provider=xero">create contact group</a></p>
<p><a href="/findall/contactgroups?provider=xero">find all contact groups</a></p>
<p><a href="/findall/currencies?provider=xero">find all currencies</a></p>
`

var invoiceTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.InvoiceID}}</p>
<p>Invoice Number: {{.InvoiceNumber}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
{{if .LineItems}}
<p>LineItems: </p>
{{range .LineItems}}
	<p>--  Description:{{.Description}}  |  Quantity:{{.Quantity}}  |  LineTotal:{{.LineAmount}}</p>
{{end}}
{{else}}
	<p>No line items were found</p>
{{end}}
<p>Total: {{.Total}}</p>
<p>AmountDue: {{.AmountDue}}</p>
<p>AmountPaid: {{.AmountPaid}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/update/invoice/{{.InvoiceID}}?provider=xero">update status of this invoice</a></p>
`

var invoicesTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.InvoiceID}}</p>
<p>Invoice Number: {{.InvoiceNumber}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>AmountDue: {{.AmountDue}}</p>
<p>AmountPaid: {{.AmountPaid}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/invoice/{{.InvoiceID}}?provider=xero">See details of this invoice</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var contactTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.ContactID}}</p>
<p>Contact Number: {{.ContactNumber}}</p>
<p>Name: {{.Name}}</p>
<p>Status: {{.ContactStatus}}</p>
<p>First Name: {{.FirstName}}</p>
<p>Last Name: {{.LastName}}</p>
<p>Email Address: {{.EmailAddress}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p>SalesDefaultAccountCode: {{.SalesDefaultAccountCode}}</p>
<p>PurchasesDefaultAccountCode: {{.PurchasesDefaultAccountCode}}</p>
<p>TrackingCategoryName: {{.TrackingCategoryName}}</p>
<p>TrackingCategoryOption: {{.TrackingCategoryOption}}</p>
<p>Amount overdue: {{.Balances.AccountsReceivable.Overdue}}</p>
<p><a href="/update/contact/{{.ContactID}}?provider=xero">update email address of this contact</a></p>
<p><a href="/findwhere/invoices?provider=xero&where=Contact.ContactID%20%3D%20Guid%28%22{{.ContactID}}%22%29%0D%0A">see invoices for this contact</a></p>
`

var contactsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.ContactID}}</p>
<p>Contact Number: {{.ContactNumber}}</p>
<p>Name: {{.Name}}</p>
<p>Status: {{.ContactStatus}}</p>
<p>First Name: {{.FirstName}}</p>
<p>Last Name: {{.LastName}}</p>
<p>Email Address: {{.EmailAddress}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/contact/{{.ContactID}}?provider=xero">See details of this Contact</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var accountTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.AccountID}}</p>
<p>Account Code: {{.Code}}</p>
<p>Name: {{.Name}}</p>
<p>Type: {{.Type}}</p>
<p>Status: {{.Status}}</p>
<p>Description: {{.Description}}</p>
<p>Tax Type: {{.TaxType}}</p>
<p>Enable payments: {{.EnablePaymentsToAccount}}</p>
<p>Show In Expense Claims: {{.ShowInExpenseClaims}}</p>
<p><a href="/update/account/{{.AccountID}}?provider=xero">Update enable payments this account</a></p>
`

var accountsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.AccountID}}</p>
<p>Account Code: {{.Code}}</p>
<p>Name: {{.Name}}</p>
<p>Type: {{.Type}}</p>
<p>Status: {{.Status}}</p>
<p>Description: {{.Description}}</p>
<p>Tax Type: {{.TaxType}}</p>
<p>Enable payments: {{.EnablePaymentsToAccount}}</p>
<p>Show In Expense Claims: {{.ShowInExpenseClaims}}</p>
<p><a href="/find/account/{{.AccountID}}?provider=xero">See details of this Account</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
var bankTransactionTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.BankTransactionID}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
{{if .LineItems}}
<p>LineItems: </p>
{{range .LineItems}}
	<p>--  Description:{{.Description}}  |  Quantity:{{.Quantity}}  |  Account:{{.AccountCode}}  |  LineTotal:{{.LineAmount}}</p>
{{end}}
{{else}}
	<p>No line items were found</p>
{{end}}
<p>Bank Account: {{.BankAccount.Code}}</p>
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/update/banktransaction/{{.BankTransactionID}}?provider=xero">update Status of this bank transaction to Deleted</a></p>
`

var bankTransactionsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.BankTransactionID}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Bank Account: {{.BankAccount.Code}}</p>
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/banktransaction/{{.BankTransactionID}}?provider=xero">See details of this bank transaction</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
var creditNoteTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.CreditNoteID}}</p>
<p>CreditNote Number: {{.CreditNoteNumber}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
{{if .LineItems}}
<p>LineItems: </p>
{{range .LineItems}}
	<p>--  Description:{{.Description}}  |  Quantity:{{.Quantity}}  |  LineTotal:{{.LineAmount}}</p>
{{end}}
{{else}}
	<p>No line items were found</p>
{{end}}
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/update/creditnote/{{.CreditNoteID}}?provider=xero">update status of this credit note</a></p>
`

var creditNotesTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.CreditNoteID}}</p>
<p>CreditNote Number: {{.CreditNoteNumber}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/creditnote/{{.CreditNoteID}}?provider=xero">See details of this credit note</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var contactGroupTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.ContactGroupID}}</p>
<p>Name: {{.Name}}</p>
<p>Status: {{.Status}}</p>
{{if .Contacts}}
<p>Contacts: </p>
{{range .Contacts}}
     <p>--  ID: {{.ContactID}}  |  Name: {{.Name}}</p>
{{end}}
{{else}}
     <p>No contacts were found</p>
{{end}}
<p><a href="/update/contactgroup/{{.ContactGroupID}}?provider=xero">Delete this contact group</a></p>
`

var contactGroupsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.ContactGroupID}}</p>
<p>Name: {{.Name}}</p>
<p>Status: {{.Status}}</p>
<p><a href="/find/contactgroup/{{.ContactGroupID}}?provider=xero">See details of this contact group</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var currenciesTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>Code: {{.Code}}</p>
<p>Description: {{.Description}}</p>
<p>-----------------------------------------------------</p>
{{end}}
`
