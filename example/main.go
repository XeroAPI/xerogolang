package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"log"
	"math"

	"github.com/TheRegan/Xero-Golang"
	"github.com/TheRegan/Xero-Golang/accounting"
	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

var (
	store = sessions.NewFilesystemStore(os.TempDir(), []byte("xero-example"))
)

func init() {
	store.MaxLength(math.MaxInt64)

	gothic.Store = store
}

func main() {
	provider := xero.New(os.Getenv("XERO_KEY"), os.Getenv("XERO_SECRET"), "http://localhost:3000/auth/callback?provider=xero")

	goth.UseProviders(
		provider,
	)

	p := pat.New()
	p.Get("/auth/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, user)
	})

	p.Get("/logout", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.Get("/auth", func(res http.ResponseWriter, req *http.Request) {
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(userTemplate)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	p.Get("/createinvoice", func(res http.ResponseWriter, req *http.Request) {
		invoices := accounting.CreateExampleInvoice()
		invoiceCollection, err := accounting.CreateInvoice(res, req, provider, store, invoices)

		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(invoiceTemplate)
		t.Execute(res, invoiceCollection.Invoices[0])
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, nil)
	})
	log.Fatal(http.ListenAndServe(":3000", p))
}

var indexTemplate = `<p><a href="/auth/?provider=xero">Connect to Xero</a></p>`

var userTemplate = `
<p><a href="/logout?provider=xero">logout</a></p>
<p>Method: {{.Email}}</p>
<p>Org Name: {{.Name}}</p>
<p>Legal Name: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>Type: {{.Description}}</p>
<p>ShortCode: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p><a href="/createinvoice?provider=xero">create invoice</a></p>
`
var invoiceTemplate = `
<p><a href="/logout?provider=xero">logout</a></p>
<p>ID: {{.InvoiceID}}</p>
<p>Invoice Number: {{.InvoiceNumber}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>AmountDue: {{.AmountDue}}</p>
<p>AmountPaid: {{.AmountPaid}}</p>
`
