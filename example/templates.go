package main

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
<p><a href="/create/item?provider=xero">create item</a></p>
<p><a href="/findall/items?provider=xero">find all items</a></p>
<p><a href="/findall/items?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all items changed since 1 May 2017</a></p>
<p><a href="/findall/journals?provider=xero">find all journals</a></p>
<p><a href="/findall/journals?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all journals changed since 1 May 2017</a></p>
<p><a href="/findall/journals/300?provider=xero">find the 100 journals after Journal 300</a></p>
<p><a href="/create/manualjournal?provider=xero">create manual journal</a></p>
<p><a href="/findall/manualjournals?provider=xero">find all manual journals</a></p>
<p><a href="/findall/manualjournals?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all manual journals changed since 1 May 2017</a></p>
<p><a href="/findall/manualjournals/1?provider=xero">find the first 100 manual journals</a></p>
<p><a href="/findall/payments?provider=xero">find all payments</a></p>
<p><a href="/findall/payments?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all payments changed since 1 May 2017</a></p>
<p><a href="/create/purchaseorder?provider=xero">create purchase order</a></p>
<p><a href="/findall/purchaseorders?provider=xero">find all purchase orders</a></p>
<p><a href="/findall/purchaseorders?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all purchase orders changed since 1 May 2017</a></p>
<p><a href="/findall/purchaseorders/1?provider=xero">find the first 100 purchase orders</a></p>
<p><a href="/create/trackingcategory?provider=xero">create tracking category</a></p>
<p><a href="/findall/trackingcategories?provider=xero">find all tracking categories</a></p>
<p><a href="/create/taxrate?provider=xero">create tax rate</a></p>
<p><a href="/findall/taxrates?provider=xero">find all tax rates</a></p>
<p><a href="/findall/overpayments?provider=xero">find all overpayments</a></p>
<p><a href="/findall/overpayments?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all overpayments changed since 1 May 2017</a></p>
<p><a href="/findall/overpayments/1?provider=xero">find the first 100 overpayments</a></p>
<p><a href="/findall/prepayments?provider=xero">find all prepayments</a></p>
<p><a href="/findall/prepayments?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all prepayments changed since 1 May 2017</a></p>
<p><a href="/findall/prepayments/1?provider=xero">find the first 100 prepayments</a></p>
<p><a href="/find/balancesheet/0?provider=xero">run the balance sheet</a></p>
<p><a href="/find/banksummary/0?provider=xero">run the bank summary</a></p>
<p><a href="/find/budgetsummary/0?provider=xero">run the budget summary</a></p>
<p><a href="/find/executivesummary/0?provider=xero">run the executive summary</a></p>
<p><a href="/find/profitandloss/0?provider=xero">run the profit and loss</a></p>
<p><a href="/find/trialbalance/0?provider=xero">run the trial balance</a></p>
<p><a href="/findall/linkedtransactions?provider=xero">find the first 100 linked transactions</a></p>
<p><a href="/findall/linkedtransactions/2?provider=xero">find the next 100 linkedtransactions</a></p>
<p><a href="/findall/users?provider=xero">find all users</a></p>
<p><a href="/findall/users?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all users changed since 1 May 2017</a></p>
<p><a href="/findall/expenseclaims?provider=xero">find all expense claims</a></p>
<p><a href="/findall/expenseclaims?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all expense claims changed since 1 May 2017</a></p>
<p><a href="/findall/receipts?provider=xero">find all receipts</a></p>
<p><a href="/findall/receipts?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all receipts changed since 1 May 2017</a></p>
<p><a href="/findall/repeatinginvoices?provider=xero">find all repeating invoices</a></p>
`

var connectTemplate = `
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
<p><a href="/create/item?provider=xero">create item</a></p>
<p><a href="/findall/items?provider=xero">find all items</a></p>
<p><a href="/findall/items?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all items changed since 1 May 2017</a></p>
<p><a href="/findall/journals?provider=xero">find all journals</a></p>
<p><a href="/findall/journals?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all journals changed since 1 May 2017</a></p>
<p><a href="/findall/journals/300?provider=xero">find the 100 journals after Journal 300</a></p>
<p><a href="/create/manualjournal?provider=xero">create manual journal</a></p>
<p><a href="/findall/manualjournals?provider=xero">find all manual journals</a></p>
<p><a href="/findall/manualjournals?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all manual journals changed since 1 May 2017</a></p>
<p><a href="/findall/manualjournals/1?provider=xero">find the first 100 manual journals</a></p>
<p><a href="/findall/payments?provider=xero">find all payments</a></p>
<p><a href="/findall/payments?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all payments changed since 1 May 2017</a></p>
<p><a href="/create/purchaseorder?provider=xero">create purchase order</a></p>
<p><a href="/findall/purchaseorders?provider=xero">find all purchase orders</a></p>
<p><a href="/findall/purchaseorders?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all purchase orders changed since 1 May 2017</a></p>
<p><a href="/findall/purchaseorders/1?provider=xero">find the first 100 purchase orders</a></p>
<p><a href="/create/trackingcategory?provider=xero">create tracking category</a></p>
<p><a href="/findall/trackingcategories?provider=xero">find all tracking categories</a></p>
<p><a href="/create/taxrate?provider=xero">create tax rate</a></p>
<p><a href="/findall/taxrates?provider=xero">find all tax rates</a></p>
<p><a href="/findall/overpayments?provider=xero">find all overpayments</a></p>
<p><a href="/findall/overpayments?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all overpayments changed since 1 May 2017</a></p>
<p><a href="/findall/overpayments/1?provider=xero">find the first 100 overpayments</a></p>
<p><a href="/findall/prepayments?provider=xero">find all prepayments</a></p>
<p><a href="/findall/prepayments?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all prepayments changed since 1 May 2017</a></p>
<p><a href="/findall/prepayments/1?provider=xero">find the first 100 prepayments</a></p>
<p><a href="/find/balancesheet/0?provider=xero">run the balance sheet</a></p>
<p><a href="/find/banksummary/0?provider=xero">run the bank summary</a></p>
<p><a href="/find/budgetsummary/0?provider=xero">run the budget summary</a></p>
<p><a href="/find/executivesummary/0?provider=xero">run the executive summary</a></p>
<p><a href="/find/profitandloss/0?provider=xero">run the profit and loss</a></p>
<p><a href="/find/trialbalance/0?provider=xero">run the trial balance</a></p>
<p><a href="/findall/linkedtransactions?provider=xero">find the first 100 linked transactions</a></p>
<p><a href="/findall/linkedtransactions/2?provider=xero">find the next 100 linkedtransactions</a></p>
<p><a href="/findall/users?provider=xero">find all users</a></p>
<p><a href="/findall/users?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all users changed since 1 May 2017</a></p>
<p><a href="/findall/expenseclaims?provider=xero">find all expense claims</a></p>
<p><a href="/findall/expenseclaims?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all expense claims changed since 1 May 2017</a></p>
<p><a href="/findall/receipts?provider=xero">find all receipts</a></p>
<p><a href="/findall/receipts?provider=xero&modifiedsince=2017-05-01T00%3A00%3A00Z">find all receipts changed since 1 May 2017</a></p>
<p><a href="/findall/repeatinginvoices?provider=xero">find all repeating invoices</a></p>
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
<p><a href="/find/agedpayablesbycontact/{{.ContactID}}?provider=xero">run aged payables report for this contact</a></p>
<p><a href="/find/agedreceivablesbycontact/{{.ContactID}}?provider=xero">run aged receivables report for this contact</a></p>
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

var itemTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>Code: {{.Code}}</p>
<p>InventoryAssetAccountCode: {{.InventoryAssetAccountCode}}</p>
<p>Name: {{.Name}}</p>
<p>IsSold: {{.IsSold}}</p>
<p>IsPurchased: {{.IsPurchased}}</p>
<p>Description: {{.Description}}</p>
<p>PurchaseDescription: {{.PurchaseDescription}}</p>
<p>PurchaseDetails:</p>
<p>--------UnitPrice: {{.PurchaseDetails.UnitPrice}}</p>
<p>--------AccountCode: {{.PurchaseDetails.AccountCode}}</p>
<p>--------COGSAccountCode: {{.PurchaseDetails.COGSAccountCode}}</p>
<p>--------TaxType: {{.PurchaseDetails.TaxType}}</p>
<p>SalesDetails:</p>
<p>--------UnitPrice: {{.SalesDetails.UnitPrice}}</p>
<p>--------AccountCode: {{.SalesDetails.AccountCode}}</p>
<p>--------COGSAccountCode: {{.SalesDetails.COGSAccountCode}}</p>
<p>--------TaxType: {{.SalesDetails.TaxType}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/update/item/{{.ItemID}}?provider=xero">update description of this item</a></p>
`

var itemsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>Code: {{.Code}}</p>
<p>InventoryAssetAccountCode: {{.InventoryAssetAccountCode}}</p>
<p>Name: {{.Name}}</p>
<p>IsSold: {{.IsSold}}</p>
<p>IsPurchased: {{.IsPurchased}}</p>
<p>Description: {{.Description}}</p>
<p>PurchaseDescription: {{.PurchaseDescription}}</p>
<p>PurchaseDetails:</p>
<p>--------UnitPrice: {{.PurchaseDetails.UnitPrice}}</p>
<p>--------AccountCode: {{.PurchaseDetails.AccountCode}}</p>
<p>--------COGSAccountCode: {{.PurchaseDetails.COGSAccountCode}}</p>
<p>--------TaxType: {{.PurchaseDetails.TaxType}}</p>
<p>SalesDetails:</p>
<p>--------UnitPrice: {{.SalesDetails.UnitPrice}}</p>
<p>--------AccountCode: {{.SalesDetails.AccountCode}}</p>
<p>--------COGSAccountCode: {{.SalesDetails.COGSAccountCode}}</p>
<p>--------TaxType: {{.SalesDetails.TaxType}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/item/{{.ItemID}}?provider=xero">See details of this item</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var journalTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.JournalID}}</p>
<p>Journal Number: {{.JournalNumber}}</p>
<p>Date: {{.JournalDate}}</p>
{{if .JournalLines}}
<p>Lines: </p>
{{range .JournalLines}}
	<p>--  Description:{{.Description}}  |  Account:{{.AccountCode}}  |  NetAmount:{{.NetAmount}}</p>
{{end}}
{{else}}
	<p>No lines were found</p>
{{end}}
`

var journalsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.JournalID}}</p>
<p>Journal Number: {{.JournalNumber}}</p>
<p>Date: {{.JournalDate}}</p>
<p><a href="/find/journal/{{.JournalID}}?provider=xero">See details of this journal</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var manualJournalTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.ManualJournalID}}</p>
<p>Narration: {{.Narration}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
{{if .JournalLines}}
<p>LineItems: </p>
{{range .JournalLines}}
	<p>--  Description:{{.Description}}  |  Account:{{.AccountCode}}  |  LineTotal:{{.LineAmount}}</p>
{{end}}
{{else}}
	<p>No line items were found</p>
{{end}}
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/update/manualjournal/{{.ManualJournalID}}?provider=xero">update status of this manual journal</a></p>
`

var manualJournalsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.ManualJournalID}}</p>
<p>Narration: {{.Narration}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/manualjournal/{{.ManualJournalID}}?provider=xero">See details of this manual journal</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var paymentTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.PaymentID}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Account: {{.Account.AccountID}}</p>
<p>Contact: {{.Invoice.Contact.Name}}</p>
<p>Invoice: {{.Invoice.InvoiceID}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/update/payment/{{.PaymentID}}?provider=xero">Delete this payment</a></p>
`

var paymentsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.PaymentID}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Account: {{.Account.AccountID}}</p>
<p>Contact: {{.Invoice.Contact.Name}}</p>
<p>Invoice: {{.Invoice.InvoiceID}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/payment/{{.PaymentID}}?provider=xero">See this payment</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
var purchaseOrderTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.PurchaseOrderID}}</p>
<p>PurchaseOrder Number: {{.PurchaseOrderNumber}}</p>
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
<p><a href="/update/purchaseorder/{{.PurchaseOrderID}}?provider=xero">update status of this purchase order</a></p>
`

var purchaseOrdersTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.PurchaseOrderID}}</p>
<p>PurchaseOrder Number: {{.PurchaseOrderNumber}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/purchaseorder/{{.PurchaseOrderID}}?provider=xero">See details of this purchase order</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
var trackingCategoryTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.TrackingCategoryID}}</p>
<p>Name: {{.Name}}</p>
<p>Status: {{.Status}}</p>
{{if .Options}}
<p>Options: </p>
{{range .Options}}
     <p>--  ID: {{.TrackingOptionID}}  |  Name: {{.Name}}</p>
{{end}}
{{else}}
     <p>No contacts were found</p>
{{end}}
<p><a href="/update/trackingcategory/{{.TrackingCategoryID}}?provider=xero">Update name of this tracking category</a></p>
`

var trackingCategoriesTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.TrackingCategoryID}}</p>
<p>Name: {{.Name}}</p>
<p>Status: {{.Status}}</p>
<p><a href="/find/trackingcategory/{{.TrackingCategoryID}}?provider=xero">See details of this tracking category</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
var taxRateTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>Name: {{.Name}}</p>
<p>TaxType: {{.TaxType}}</p>
<p>ReportTaxType: {{.ReportTaxType}}</p>
<p>Status: {{.Status}}</p>
{{if .TaxComponents}}
<p>TaxComponents: </p>
{{range .TaxComponents}}
     <p>--  Name: {{.Name}}   |   Rate:  {{.Rate}}</p>
{{end}}
{{else}}
     <p>No Tax Components were found</p>
{{end}}
<p><a href="/update/taxrate/{{.Name}}?provider=xero">Delete this tax rate</a></p>
`

var taxRatesTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>Name: {{.Name}}</p>
<p>TaxType: {{.TaxType}}</p>
<p>ReportTaxType: {{.ReportTaxType}}</p>
<p>Status: {{.Status}}</p>
{{if .TaxComponents}}
<p>TaxComponents: </p>
{{range .TaxComponents}}
     <p>--  Name: {{.Name}}   |   Rate:  {{.Rate}}</p>
{{end}}
{{else}}
     <p>No Tax Components were found</p>
{{end}}
<p>-----------------------------------------------------</p>
{{end}}
`

var overpaymentTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.OverpaymentID}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
{{if .Allocations}}
<p>Allocations: </p>
{{range .Allocations}}
	<p>--  AppliedAmount:{{.AppliedAmount}}  |  Date:{{.Date}}  |  Invoice:{{.Invoice.InvoiceID}}</p>
{{end}}
{{else}}
	<p>No Allocations were found</p>
{{end}}
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
`

var overpaymentsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.OverpaymentID}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/overpayment/{{.OverpaymentID}}?provider=xero">See details of this overpayment</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
var prepaymentTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.PrepaymentID}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
{{if .Allocations}}
<p>Allocations: </p>
{{range .Allocations}}
	<p>--  AppliedAmount:{{.AppliedAmount}}  |  Date:{{.Date}}  |  Invoice:{{.Invoice.InvoiceID}}</p>
{{end}}
{{else}}
	<p>No Allocations were found</p>
{{end}}
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
`

var prepaymentsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.PrepaymentID}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/prepayment/{{.PrepaymentID}}?provider=xero">See details of this prepayment</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
var reportTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.ReportID}}</p>
<p>Name: {{.ReportName}}</p>
<p>Date: {{.ReportDate}}</p>
<p>Type: {{.ReportType}}</p>
{{if .Attributes}}
<p>Attributes: </p>
{{range .Attributes}}
	<p>--  Name:{{.Name}}  |  Description:{{.Description}}  |  Value:{{.Value}}</p>
{{end}}
{{else}}
	<p></p>
{{end}}
{{if .ReportTitles}}
<p>ReportTitles: </p>
{{range .ReportTitles}}
	<p>--  ReportTitle:{{.}}</p>
{{end}}
{{else}}
	<p>No ReportTitles were found</p>
{{end}}
{{if .Rows}}
<p>Rows: </p>
{{range .Rows}}
	<p>--  Type:{{.RowType}}  |  Title:{{.Title}}</p>
	{{if .Rows}}
	<p>Rows: </p>
	{{range .Rows}}
		<p>--  Type:{{.RowType}}  |  Title:{{.Title}}</p>
		{{if .Cells}}
		<p>Cells: </p>
		{{range .Cells}}
			<p>--  Value:{{.Value}}</p>
			{{if .Attributes}}
			<p>Attributes: </p>
			{{range .Attributes}}
				<p>-- ID:{{.ID}}  |  Value:{{.Value}}</p>
			{{end}}
			{{else}}
				<p></p>
			{{end}}
		{{end}}
		{{else}}
			<p></p>
		{{end}}
	{{end}}
	{{else}}
		<p>No Rows were found</p>
	{{end}}
	{{if .Cells}}
	<p>Cells: </p>
	{{range .Cells}}
		<p>--  Value:{{.Value}}</p>
		{{if .Attributes}}
		<p>Attributes: </p>
		{{range .Attributes}}
			<p>-- ID:{{.ID}}  |  Value:{{.Value}}</p>
		{{end}}
		{{else}}
			<p></p>
		{{end}}
	{{end}}
	{{else}}
		<p></p>
	{{end}}
{{end}}
{{else}}
	<p>No Rows were found</p>
{{end}}
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p>--------------------------------------------------</p>
`

var linkedTransactionTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.LinkedTransactionID}}</p>
<p>ContactId: {{.ContactID}}</p>
<p>SourceTransactionID: {{.SourceTransactionID}}</p>
<p>SourceLineItemID: {{.SourceLineItemID}}</p>
<p>TargetTransactionID: {{.TargetTransactionID}}</p>
<p>TargetLineItemID: {{.TargetLineItemID}}</p>
<p>Status: {{.Status}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
`

var linkedTransactionsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.LinkedTransactionID}}</p>
<p>ContactId: {{.ContactID}}</p>
<p>SourceTransactionID: {{.SourceTransactionID}}</p>
<p>SourceLineItemID: {{.SourceLineItemID}}</p>
<p>TargetTransactionID: {{.TargetTransactionID}}</p>
<p>TargetLineItemID: {{.TargetLineItemID}}</p>
<p>Status: {{.Status}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/linkedtransaction/{{.LinkedTransactionID}}?provider=xero">See details of this linked transaction</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var userTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.UserID}}</p>
<p>EmailAddress: {{.EmailAddress}}</p>
<p>FirstName: {{.FirstName}}</p>
<p>LastName: {{.LastName}}</p>
<p>IsSubscriber: {{.IsSubscriber}}</p>
<p>OrganisationRole: {{.OrganisationRole}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
`

var usersTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.UserID}}</p>
<p>EmailAddress: {{.EmailAddress}}</p>
<p>FirstName: {{.FirstName}}</p>
<p>LastName: {{.LastName}}</p>
<p>IsSubscriber: {{.IsSubscriber}}</p>
<p>OrganisationRole: {{.OrganisationRole}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/user/{{.UserID}}?provider=xero">See details of this user</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var expenseClaimTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.ExpenseClaimID}}</p>
<p>User: {{.User.UserID}}</p>
<p>FirstName: {{.User.FirstName}}</p>
<p>LastName: {{.User.LastName}}</p>
{{if .Receipts}}
<p>Receipts: </p>
{{range .Receipts}}
     <p>--  ID: {{.ReceiptID}}   |   Number:  {{.ReceiptNumber}}   |   Total:  {{.Total}}</p>
{{end}}
{{else}}
     <p>No Tax Receipts were found</p>
{{end}}
<p>AmountDue: {{.AmountDue}}</p>
<p>AmountPaid: {{.AmountPaid}}</p>
<p>Total: {{.Total}}</p>
`

var expenseClaimsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.ExpenseClaimID}}</p>
<p>User: {{.User.UserID}}</p>
<p>FirstName: {{.User.FirstName}}</p>
<p>LastName: {{.User.LastName}}</p>
<p>AmountDue: {{.AmountDue}}</p>
<p>AmountPaid: {{.AmountPaid}}</p>
<p>Total: {{.Total}}</p>
<p><a href="/find/expenseclaim/{{.ExpenseClaimID}}?provider=xero">See details of this expense claim</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var receiptTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.ReceiptID}}</p>
<p>Receipt Number: {{.ReceiptNumber}}</p>
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
<p>Reference: {{.Reference}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/update/receipt/{{.ReceiptID}}?provider=xero">update reference of this receipt</a></p>
`

var receiptsTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.ReceiptID}}</p>
<p>Receipt Number: {{.ReceiptNumber}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Date: {{.Date}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>Reference: {{.Reference}}</p>
<p>UpdatedDate: {{.UpdatedDateUTC}}</p>
<p><a href="/find/receipt/{{.ReceiptID}}?provider=xero">See details of this receipt</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`

var repeatingInvoiceTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
<p>ID: {{.RepeatingInvoiceID}}</p>
<p>Contact: {{.Contact.Name}}</p>
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
<p>StartDate: {{.Schedule.StartDate}}</p>
<p>EndDate: {{.Schedule.EndDate}}</p>
<p>NextScheduledDate: {{.Schedule.NextScheduledDate}}</p>
`

var repeatingInvoicesTemplate = `
<p><a href="/disconnect?provider=xero">logout</a></p>
{{range $index,$element:= .}}
<p>ID: {{.RepeatingInvoiceID}}</p>
<p>Contact: {{.Contact.Name}}</p>
<p>Status: {{.Status}}</p>
<p>Total: {{.Total}}</p>
<p>StartDate: {{.Schedule.StartDate}}</p>
<p>EndDate: {{.Schedule.EndDate}}</p>
<p>NextScheduledDate: {{.Schedule.NextScheduledDate}}</p>
<p><a href="/find/repeatinginvoice/{{.RepeatingInvoiceID}}?provider=xero">See details of this repeating invoice</a></p>
<p>-----------------------------------------------------</p>
{{end}}
`
