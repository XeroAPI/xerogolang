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
	provider = xero.New(os.Getenv("XERO_KEY"), os.Getenv("XERO_SECRET"), "http://localhost:3000/auth/callback?provider=xero")
	store    = sessions.NewFilesystemStore(os.TempDir(), []byte("xero-example"))
	invoices = new(accounting.Invoices)
	contacts = new(accounting.Contacts)
)

func init() {
	goth.UseProviders(provider)

	store.MaxLength(math.MaxInt64)

	gothic.Store = store
}

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

func authHandler(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, gothUser)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func callbackHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(res, user)
}

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
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

func disconnectHandler(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

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
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

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
			invoiceCollection, err = accounting.FindAllInvoices(provider, session)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, err)
				return
			}
			invoiceCollection, err = accounting.FindAllInvoicesModifiedSince(provider, session, parsedTime)
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
			contactCollection, err = accounting.FindAllContacts(provider, session)
		} else {
			parsedTime, parseError := time.Parse(time.RFC3339, modifiedSince)
			if parseError != nil {
				fmt.Fprintln(res, parseError)
				return
			}
			contactCollection, err = accounting.FindAllContactsModifiedSince(provider, session, parsedTime)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactsTemplate)
		t.Execute(res, contactCollection.Contacts)
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

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
			invoiceCollection, err = accounting.FindInvoicesByPageModifiedSince(provider, session, pageInt, parsedTime)
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
			contactCollection, err = accounting.FindContactsByPageModifiedSince(provider, session, pageInt, parsedTime)
		}
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(contactsTemplate)
		t.Execute(res, contactCollection.Contacts)
	default:
		fmt.Fprintln(res, "Unknown type specified")
		return
	}
}

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
