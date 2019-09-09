package xerogolang

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
	"golang.org/x/oauth2"
)

var (
	oauth2AuthURL  = "https:///login.xero.com/identity/connect/authorize"
	oauth2TokenURL = "https://identity.xero.com/connect/token"
	// oauth2TokenURL = "https://oauth-proxy.omniboost.io"
)

func init() {
	// oauth2.RegisterBrokenAuthHeaderProvider("login.xero.com")
	// oauth2.RegisterBrokenAuthHeaderProvider("identity.xero.com")
}

// New creates a new Xero provider, and sets up important connection details.
// You should always call `xero.New` to get a new Provider. Never try to create
// one manually.
func NewOauth2(clientID, clientSecret string, token *oauth2.Token) *Oauth2Provider {
	p := &Oauth2Provider{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		CallbackURL:  "",
		Token:        token,
		//Method determines how you will connect to Xero.
		//Use public if this is your first time.
		//Options are public, private, and partner
		//More details here: https://developer.xero.com/documentation/getting-started/api-application-types
		UserAgentString: userAgentString,
		providerName:    "xero",
		TenantID:        "",
	}
	p.config = newOauth2Config(p, []string{})
	return p
}

// Oauth2Provider is the implementation of `goth.Oauth2Provider` for accessing Xero.
type Oauth2Provider struct {
	ClientID        string
	ClientSecret    string
	CallbackURL     string
	Token           *oauth2.Token
	HTTPClient      *http.Client
	Method          string
	UserAgentString string
	debug           bool
	config          *oauth2.Config
	providerName    string
	TenantID        string
}

//Find retrieves the requested data from an endpoint to be unmarshaled into the appropriate data type
func (p *Oauth2Provider) Find(session goth.Session, endpoint string, additionalHeaders map[string]string, querystringParameters map[string]string) ([]byte, error) {
	var querystring string
	if querystringParameters != nil {
		for key, value := range querystringParameters {
			escapedValue := url.QueryEscape(value)
			querystring = querystring + "&" + key + "=" + escapedValue
		}
		querystring = strings.TrimPrefix(querystring, "&")
		querystring = "?" + querystring
	}

	request, err := http.NewRequest("GET", endpointProfile+endpoint+querystring, nil)
	if err != nil {
		return nil, err
	}

	return p.processRequest(request, session, additionalHeaders)
}

//Create sends data to an endpoint and returns a response to be unmarshaled into the appropriate data type
func (p *Oauth2Provider) Create(session goth.Session, endpoint string, additionalHeaders map[string]string, body []byte) ([]byte, error) {
	bodyReader := bytes.NewReader(body)

	request, err := http.NewRequest("PUT", endpointProfile+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	return p.processRequest(request, session, additionalHeaders)
}

//Update sends data to an endpoint and returns a response to be unmarshaled into the appropriate data type
func (p *Oauth2Provider) Update(session goth.Session, endpoint string, additionalHeaders map[string]string, body []byte) ([]byte, error) {
	bodyReader := bytes.NewReader(body)

	request, err := http.NewRequest("POST", endpointProfile+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	return p.processRequest(request, session, additionalHeaders)
}

//Remove deletes the specified data from an endpoint
func (p *Oauth2Provider) Remove(session goth.Session, endpoint string, additionalHeaders map[string]string) ([]byte, error) {
	request, err := http.NewRequest("DELETE", endpointProfile+endpoint, nil)
	if err != nil {
		return nil, err
	}

	return p.processRequest(request, session, additionalHeaders)
}

// Client does pretty much everything
func (p *Oauth2Provider) Client() *http.Client {
	return p.config.Client(oauth2.NoContext, p.Token)
}

func (p *Oauth2Provider) Config() *oauth2.Config {
	return p.config
}

// Debug sets the logging of the OAuth client to verbose.
func (p *Oauth2Provider) Debug(debug bool) {
	p.debug = debug
}

//processRequest processes a request prior to it being sent to the API
func (p *Oauth2Provider) processRequest(request *http.Request, session goth.Session, additionalHeaders map[string]string) ([]byte, error) {
	request.Header.Add("User-Agent", p.UserAgentString)
	request.Header.Add("Xero-tenant-id", p.TenantID)
	for key, value := range additionalHeaders {
		request.Header.Add(key, value)
	}

	if p.debug {
		b, err := httputil.DumpRequest(request, true)
		if err != nil {
			return nil, err
		}
		log.Println(string(b))
	}

	var err error
	var response *http.Response

	response, err = p.Client().Do(request)

	if p.debug && response != nil {
		b, err := httputil.DumpResponse(response, true)
		if err != nil {
			return nil, err
		}
		log.Println(string(b))
	}

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			helpers.ReaderToString(response.Body),
		)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read response: %s", err.Error())
	}
	if responseBytes == nil {
		return nil, fmt.Errorf("Received no response: %s", err.Error())
	}
	return responseBytes, nil
}

func newOauth2Config(provider *Oauth2Provider, scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     provider.ClientID,
		ClientSecret: provider.ClientSecret,
		RedirectURL:  provider.CallbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  oauth2AuthURL,
			TokenURL: oauth2TokenURL,
		},
		Scopes: []string{},
	}

	if len(scopes) > 0 {
		for _, scope := range scopes {
			c.Scopes = append(c.Scopes, scope)
		}
	}
	return c
}
