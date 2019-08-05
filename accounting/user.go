package accounting

import (
	"encoding/json"
	"time"

	"github.com/XeroAPI/xerogolang"
	"github.com/markbates/goth"
)

//User represents a login that has been granted access to the Xero organisation you're connected to
type User struct {

	// Xero identifier
	UserID string `json:"UserID,omitempty" xml:"UserID,omitempty"`

	// Email address of user
	EmailAddress string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`

	// First name of user
	FirstName string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`

	// Last name of user
	LastName string `json:"LastName,omitempty" xml:"LastName,omitempty"`

	// Timestamp of last change to user
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`

	// Boolean to indicate if user is the subscriber
	IsSubscriber bool `json:"IsSubscriber,omitempty" xml:"IsSubscriber,omitempty"`

	// User role (see Types)
	OrganisationRole string `json:"OrganisationRole,omitempty" xml:"OrganisationRole,omitempty"`
}

//Users is a collection of Users
type Users struct {
	Users []User `json:"Users,omitempty" xml:"User,omitempty"`
}

func unmarshalUsers(userResponseBytes []byte) (*Users, error) {
	var userResponse *Users
	err := json.Unmarshal(userResponseBytes, &userResponse)
	if err != nil {
		return nil, err
	}

	return userResponse, err
}

//FindUsersModifiedSince will get all users modified after a specified date
//additional querystringParameters such as where and order can be added as a map
func FindUsersModifiedSince(provider *xerogolang.Provider, session goth.Session, modifiedSince time.Time, querystringParameters map[string]string) (*Users, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if !modifiedSince.Equal(dayZero) {
		additionalHeaders["If-Modified-Since"] = modifiedSince.Format(time.RFC3339)
	}

	userResponseBytes, err := provider.Find(session, "Users", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalUsers(userResponseBytes)
}

//FindUsers will get all users
//additional querystringParameters such as where and order can be added as a map
func FindUsers(provider *xerogolang.Provider, session goth.Session, querystringParameters map[string]string) (*Users, error) {
	return FindUsersModifiedSince(provider, session, dayZero, querystringParameters)
}

//FindUser will get a single user - UserID must be a GUID for a user
func FindUser(provider *xerogolang.Provider, session goth.Session, userID string) (*Users, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	userResponseBytes, err := provider.Find(session, "Users/"+userID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalUsers(userResponseBytes)
}
