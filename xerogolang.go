package xerogolang

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"crypto"

	"github.com/XeroAPI/xerogolang/auth"
	"github.com/XeroAPI/xerogolang/helpers"
	"github.com/markbates/goth"
	"github.com/mrjones/oauth"
	"golang.org/x/oauth2"
)

var (
	requestURL      = "https://api.xero.com/oauth/RequestToken"
	authorizeURL    = "https://api.xero.com/oauth/Authorize"
	tokenURL        = "https://api.xero.com/oauth/AccessToken"
	endpointProfile = "https://api.xero.com/api.xro/2.0/"
	//userAgentString should match the name of your Application
	userAgentString = os.Getenv("XERO_USER_AGENT") + " (xerogolang 0.1.3) " + os.Getenv("XERO_KEY")
	//privateKeyFilePath is a file path to your .pem private/public key file
	//You only need this for private and partner Applications
	//more details here: https://developer.xero.com/documentation/api-guides/create-publicprivate-key
	privateKeyFilePath = os.Getenv("XERO_PRIVATE_KEY_PATH")
)

type IProvider interface {
	Find(goth.Session, string, map[string]string, map[string]string) ([]byte, error)
	Create(goth.Session, string, map[string]string, []byte) ([]byte, error)
	Update(goth.Session, string, map[string]string, []byte) ([]byte, error)
	Remove(goth.Session, string, map[string]string) ([]byte, error)
}

// Provider is the implementation of `goth.Provider` for accessing Xero.
type Provider struct {
	ClientKey       string
	Secret          string
	CallbackURL     string
	HTTPClient      *http.Client
	Method          string
	UserAgentString string
	PrivateKey      string
	debug           bool
	consumer        *oauth.Consumer
	providerName    string
}

//newPublicConsumer creates a consumer capable of communicating with a Public application: https://developer.xero.com/documentation/auth-and-limits/public-applications
func (p *Provider) newPublicConsumer(authURL string) *oauth.Consumer {

	var c *oauth.Consumer

	if p.HTTPClient != nil {
		c = oauth.NewCustomHttpClientConsumer(
			p.ClientKey,
			p.Secret,
			oauth.ServiceProvider{
				RequestTokenUrl:   requestURL,
				AuthorizeTokenUrl: authURL,
				AccessTokenUrl:    tokenURL},
			p.HTTPClient,
		)
	} else {
		c = oauth.NewConsumer(
			p.ClientKey,
			p.Secret,
			oauth.ServiceProvider{
				RequestTokenUrl:   requestURL,
				AuthorizeTokenUrl: authURL,
				AccessTokenUrl:    tokenURL},
		)
	}

	c.Debug(p.debug)

	return c
}

//newPartnerConsumer creates a consumer capable of communicating with a Partner application: https://developer.xero.com/documentation/auth-and-limits/partner-applications
func (p *Provider) newPrivateOrPartnerConsumer(authURL string) *oauth.Consumer {
	block, _ := pem.Decode([]byte(p.PrivateKey))

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	var c *oauth.Consumer

	if p.HTTPClient != nil {
		c = oauth.NewCustomRSAConsumer(
			p.ClientKey,
			privateKey,
			crypto.SHA1,
			oauth.ServiceProvider{
				RequestTokenUrl:   requestURL,
				AuthorizeTokenUrl: authURL,
				AccessTokenUrl:    tokenURL},
			p.HTTPClient,
		)
	} else {
		c = oauth.NewRSAConsumer(
			p.ClientKey,
			privateKey,
			oauth.ServiceProvider{
				RequestTokenUrl:   requestURL,
				AuthorizeTokenUrl: authURL,
				AccessTokenUrl:    tokenURL},
		)
	}

	c.Debug(p.debug)

	return c
}

// New creates a new Xero provider, and sets up important connection details.
// You should always call `xero.New` to get a new Provider. Never try to create
// one manually.
func New(clientKey, secret, callbackURL string) *Provider {
	p := &Provider{
		ClientKey:   clientKey,
		Secret:      secret,
		CallbackURL: callbackURL,
		//Method determines how you will connect to Xero.
		//Options are public, private, and partner
		//Use public if this is your first time.
		//More details here: https://developer.xero.com/documentation/getting-started/api-application-types
		Method:          os.Getenv("XERO_METHOD"),
		PrivateKey:      helpers.ReadPrivateKeyFromPath(privateKeyFilePath),
		UserAgentString: userAgentString,
		providerName:    "xero",
	}
	return p
}

// NewNoEnviro creates a new Xero provider without using the environmental set variables
// , and sets up important connection details.
// You should always call `xero.New` to get a new Provider. Never try to create
// one manually.
func NewNoEnviro(clientKey, secret, callbackURL, userAgent, xeroMethod string, privateKey []byte) *Provider {
	// Set variables without using the environment
	userAgentString = userAgent + " (xerogolang 0.1.3) " + clientKey
	privateKeyFilePath = ""

	p := &Provider{
		ClientKey:   clientKey,
		Secret:      secret,
		CallbackURL: callbackURL,
		//Method determines how you will connect to Xero.
		//Options are public, private, and partner
		//Use public if this is your first time.
		//More details here: https://developer.xero.com/documentation/getting-started/api-application-types
		Method:          xeroMethod,
		PrivateKey:      string(privateKey),
		UserAgentString: userAgentString,
		providerName:    "xero",
	}
	return p
}

// New creates a new Xero provider, with a custom http client
func NewCustomHTTPClient(clientKey, secret, callbackURL string, httpClient *http.Client) *Provider {
	p := &Provider{
		ClientKey:   clientKey,
		Secret:      secret,
		CallbackURL: callbackURL,

		Method:          os.Getenv("XERO_METHOD"),
		PrivateKey:      helpers.ReadPrivateKeyFromPath(privateKeyFilePath),
		UserAgentString: userAgentString,
		providerName:    "xero",
		HTTPClient:      httpClient,
	}
	return p
}

// Name is the name used to retrieve this provider later.
func (p *Provider) Name() string {
	return p.providerName
}

// SetName is to update the name of the provider (needed in case of multiple providers of 1 type)
func (p *Provider) SetName(name string) {
	p.providerName = name
}

// Client does pretty much everything
func (p *Provider) Client() *http.Client {
	return goth.HTTPClientWithFallBack(p.HTTPClient)
}

// Debug sets the logging of the OAuth client to verbose.
func (p *Provider) Debug(debug bool) {
	p.debug = debug
}

// BeginAuth asks Xero for an authentication end-point and a request token for a session.
// Xero does not support the "state" variable.
func (p *Provider) BeginAuth(state string) (goth.Session, error) {
	if p.consumer == nil {
		p.initConsumer()
	}

	if p.Method == "private" {
		accessToken := &oauth.AccessToken{
			Token:  p.ClientKey,
			Secret: p.Secret,
		}
		privateSession := &Session{
			AuthURL:            authorizeURL,
			RequestToken:       nil,
			AccessToken:        accessToken,
			AccessTokenExpires: time.Now().UTC().Add(87600 * time.Hour),
		}
		return privateSession, nil
	}
	requestToken, url, err := p.consumer.GetRequestTokenAndUrl(p.CallbackURL)
	if err != nil {
		return nil, err
	}
	session := &Session{
		AuthURL:      url,
		RequestToken: requestToken,
	}
	return session, nil
}

//processRequest processes a request prior to it being sent to the API
func (p *Provider) processRequest(request *http.Request, session goth.Session, additionalHeaders map[string]string) ([]byte, error) {
	sess := session.(*Session)

	if p.consumer == nil {
		p.initConsumer()
	}

	if sess.AccessToken == nil {
		// data is not yet retrieved since accessToken is still empty
		return nil, fmt.Errorf("%s cannot process request without accessToken", p.providerName)
	}

	request.Header.Add("User-Agent", p.UserAgentString)
	for key, value := range additionalHeaders {
		request.Header.Add(key, value)
	}

	var err error
	var response *http.Response

	if p.HTTPClient == nil {

		client, _ := p.consumer.MakeHttpClient(sess.AccessToken)

		response, err = client.Do(request)

	} else {

		transport, _ := p.consumer.MakeRoundTripper(sess.AccessToken)

		response, err = transport.RoundTrip(request)
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

//Find retrieves the requested data from an endpoint to be unmarshaled into the appropriate data type
func (p *Provider) Find(session goth.Session, endpoint string, additionalHeaders map[string]string, querystringParameters map[string]string) ([]byte, error) {
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
func (p *Provider) Create(session goth.Session, endpoint string, additionalHeaders map[string]string, body []byte) ([]byte, error) {
	bodyReader := bytes.NewReader(body)

	request, err := http.NewRequest("PUT", endpointProfile+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	return p.processRequest(request, session, additionalHeaders)
}

//Update sends data to an endpoint and returns a response to be unmarshaled into the appropriate data type
func (p *Provider) Update(session goth.Session, endpoint string, additionalHeaders map[string]string, body []byte) ([]byte, error) {
	bodyReader := bytes.NewReader(body)

	request, err := http.NewRequest("POST", endpointProfile+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	return p.processRequest(request, session, additionalHeaders)
}

//Remove deletes the specified data from an endpoint
func (p *Provider) Remove(session goth.Session, endpoint string, additionalHeaders map[string]string) ([]byte, error) {
	request, err := http.NewRequest("DELETE", endpointProfile+endpoint, nil)
	if err != nil {
		return nil, err
	}

	return p.processRequest(request, session, additionalHeaders)
}

//Organisation is the expected response from the Organisation endpoint - this is not a complete schema
//and should only be used by FetchUser
type Organisation struct {
	// Display name of organisation shown in Xero
	Name string `json:"Name,omitempty"`

	// Organisation name shown on Reports
	LegalName string `json:"LegalName,omitempty"`

	// Organisation Type
	OrganisationType string `json:"OrganisationType,omitempty"`

	// Country code for organisation. See ISO 3166-2 Country Codes
	CountryCode string `json:"CountryCode,omitempty"`

	// A unique identifier for the organisation.
	ShortCode string `json:"ShortCode,omitempty"`
}

//OrganisationCollection is the Total response from the Xero API
type OrganisationCollection struct {
	Organisations []Organisation `json:"Organisations,omitempty"`
}

// FetchUser will go to Xero and access basic information about the user.
func (p *Provider) FetchUser(session goth.Session) (goth.User, error) {
	sess := session.(*Session)
	user := goth.User{
		Provider: p.Name(),
	}
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}
	responseBytes, err := p.Find(sess, "Organisation", additionalHeaders, nil)
	if err != nil {
		return user, err
	}
	var organisationCollection OrganisationCollection
	err = json.Unmarshal(responseBytes, &organisationCollection)
	if err != nil {
		return user, fmt.Errorf("Could not unmarshal response: %s", err.Error())
	}

	user.Name = organisationCollection.Organisations[0].Name
	user.NickName = organisationCollection.Organisations[0].LegalName
	user.Location = organisationCollection.Organisations[0].CountryCode
	user.Description = organisationCollection.Organisations[0].OrganisationType
	user.UserID = organisationCollection.Organisations[0].ShortCode

	user.AccessToken = sess.AccessToken.Token
	user.AccessTokenSecret = sess.AccessToken.Secret
	user.ExpiresAt = sess.AccessTokenExpires
	user.Email = p.Method
	return user, err
}

//RefreshOAuth1Token should be used instead of RefeshToken which is not compliant with the Oauth1.0a standard
func (p *Provider) RefreshOAuth1Token(session *Session) error {
	if p.consumer == nil {
		p.initConsumer()
	}
	if session.AccessToken == nil {
		return fmt.Errorf("Could not refresh token as last valid accessToken was not found")
	}
	newAccessToken, err := p.consumer.RefreshToken(session.AccessToken)
	if err != nil {
		return err
	}
	session.AccessToken = newAccessToken
	session.AccessTokenExpires = time.Now().UTC().Add(30 * time.Minute)
	return nil
}

//RefreshToken refresh token is not provided by the Xero Public or Private Application -
//only the Partner Application and you must use RefreshOAuth1Token instead
func (p *Provider) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	return nil, errors.New("Refresh token is only provided by Xero for Partner Applications")
}

//RefreshTokenAvailable refresh token is not provided by the Xero Public or Private Application -
//only the Partner Application and you must use RefreshOAuth1Token instead
func (p *Provider) RefreshTokenAvailable() bool {
	return false
}

//GetSessionFromStore returns a session for a given a request and a response
//This is an exaple of how you could get a session from a store - as long as you're
//supplying a goth.Session to the interactors it will work though so feel free to use your
//own method
func (p *Provider) GetSessionFromStore(request *http.Request, response http.ResponseWriter) (goth.Session, error) {
	sessionMarshalled, _ := auth.Store.Get(request, "xero"+auth.SessionName)
	value := sessionMarshalled.Values["xero"]
	if value == nil {
		return nil, errors.New("could not find a matching session for this request")
	}
	session, err := p.UnmarshalSession(value.(string))
	if err != nil {
		return nil, errors.New("could not unmarshal session for this request")
	}
	sess := session.(*Session)
	if sess.AccessTokenExpires.Before(time.Now().UTC().Add(5 * time.Minute)) {
		if p.Method == "partner" {
			p.RefreshOAuth1Token(sess)
			sessionMarshalled.Values["xero"] = sess.Marshal()
			err = sessionMarshalled.Save(request, response)
			return session, err
		}
		return nil, errors.New("access token has expired - please reconnect")
	}
	return session, err
}

func (p *Provider) initConsumer() {
	switch p.Method {
	case "private":
		p.consumer = p.newPrivateOrPartnerConsumer(authorizeURL)
	case "public":
		p.consumer = p.newPublicConsumer(authorizeURL)
	case "partner":
		p.consumer = p.newPrivateOrPartnerConsumer(authorizeURL)
	default:
		p.consumer = p.newPublicConsumer(authorizeURL)
	}
}
