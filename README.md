# xerogolang (alpha)
This is the Xero Golang SDK for the [Xero API](https://developer.xero.com/).

Currently it only supports the Accounting API.


### Xero App
You'll need to decide which type of Xero app you'll be building [Private](http://developer.xero.com/documentation/auth-and-limits/private-applications/), [Public](http://developer.xero.com/documentation/auth-and-limits/public-applications/), or [Partner](http://developer.xero.com/documentation/auth-and-limits/partner-applications/). Go to [http://app.xero.com](http://app.xero.com) and login with your Xero user account to create an app.

### Download Xero Golang SDK
Download the SDK using the following command:
```text
$ go get github.com/XeroAPI/xerogolang
```


### Configure
The Xero Golang SDK is easily configured using environment variables to configure values unique to your Application.

The following values need to be set:
```text
XERO_KEY=Your_Consumer_Key
XERO_SECRET=Your_Consumer_Secret
XERO_USER_AGENT=Your_application_name
```
You must also set a method - this must be either ["public"](http://developer.xero.com/documentation/auth-and-limits/public-applications/), ["private"](http://developer.xero.com/documentation/auth-and-limits/private-applications/), or ["partner"](http://developer.xero.com/documentation/auth-and-limits/partner-applications/)
```text
XERO_METHOD=public_private_or_partner
```
If you are using the ["private"](http://developer.xero.com/documentation/auth-and-limits/private-applications/) or ["partner"](http://developer.xero.com/documentation/auth-and-limits/partner-applications/) method you'll also need a [Private key](https://developer.xero.com/documentation/api-guides/create-publicprivate-key) path:
```text
XERO_PRIVATE_KEY_PATH=/Path/to/your/privatekey.pem
```

We include an Example App (in this repo) built using [Gorilla](http://www.gorillatoolkit.org/).

### Example App
This repo includes an Example App mentioned above.  The app contains examples of most of the functions available via the API.

To run the example app do the following:
```text
$ cd xerogolang/example
$ go get -v
$ go build
$ ./example
```
Now open up your browser and go to [http://localhost:3000](http://localhost:3000) to see the example.

The example app uses a filesystem store for sessions - you will need to implement your own store when using this SDK within your own app. We recommend [Gorilla Sessions](https://github.com/gorilla/sessions)


**Data Endpoints**

The Xero Golang SDK contains the Accounting package which has helper methods to perform (Create, Find, Update and Remove) actions on each endpoints and structs of each endpoint's response.  If an endpoint does not have one of the methods then that method is not available on the endpoint. E.g. Branding Themes can only use Find methods because you cannot Create, Update, or Remove them via the API.

#### Create
Create can be called on structs that have been populated with data:
```go
c := &Contacts{
  Contacts: []Contact{
    Contact{
      Name: "Cosmo Kramer",
    },
  },
}
r, err := c.Create(provider, session)
```

#### Find
Find is called either to get a single entity given an id:
```go
i, err := accounting.FindInvoice(provider, session, id)
```
all entities from an endpoint:
```go
i, err = accounting.FindInvoices(provider, session, nil)
```
all entities from an endpoint since a specific time:
```go
i, err = accounting.FindInvoicesModifiedSince(provider, session, time.Now().Add(-24*time.Hour), nil)
```
all entities from an endpoint using paging:
```go
querystringParameters := map[string]string{
  "page": 1,
}

i, err = accounting.FindInvoices(provider, session, querystringParameters)
```
all entities from an endpoint that match a given where clause:
```go
querystringParameters := map[string]string{
  "where": "Contact.Name==\"Vanderlay Industries\"",
}

i, err = accounting.FindInvoices(provider, session, querystringParameters)
```
all entities from an endpoint in a particular order:
```go
querystringParameters := map[string]string{
  "order": "DueDate",
}

i, err = accounting.FindInvoices(provider, session, querystringParameters)
```
all entities from an endpoint using a filter:
```go
querystringParameters := map[string]string{
  "Statuses": "DRAFT,SUBMITTED",
}

i, err = accounting.FindInvoices(provider, session, querystringParameters)
```
a combination of all of the above:
```go
querystringParameters := map[string]string{
  "page": 1,
  "where": "Contact.Name==\"Vanderlay Industries\"",
  "order": "DueDate",
  "Statuses": "DRAFT,SUBMITTED",
}

i, err = accounting.FindInvoicesModifiedSince(provider, session, time.Now().Add(-24*time.Hour), querystringParameters)
```

#### Update
Update can be called on a struct containing the data to update.  You can only update one entity at a time though.
```go
a, err := accounts.Update(provider, session)
```

#### Remove
Remove can be called to remove an entity if you provide an ID - it is not provided on all endpoints though.
```go
t, err := RemoveTrackingCategory(provider, session, "trackingCategoryID")
```

## Acknowledgement

The Xero golang SDK is extended from the great oauth work done by [markbates' Goth](https://github.com/markbates/goth) and [mrjones' oauth](https://github.com/mrjones/oauth).  We have added support for Xero a provider directly in goth as well so if for some reason you don't want models and methods you can use goth directly.

The [gorilla web toolkit](https://github.com/gorilla) is used throughout this SDK and comes highly recommended.


## License

This software is published under the [MIT License](http://en.wikipedia.org/wiki/MIT_License).

	Copyright (c) 2016 Xero Limited

	Permission is hereby granted, free of charge, to any person
	obtaining a copy of this software and associated documentation
	files (the "Software"), to deal in the Software without
	restriction, including without limitation the rights to use,
	copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the
	Software is furnished to do so, subject to the following
	conditions:

	The above copyright notice and this permission notice shall be
	included in all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
	OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
	WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
	FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
	OTHER DEALINGS IN THE SOFTWARE.
