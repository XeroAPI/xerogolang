package accounting

type User struct {

	// Xero identifier
	UserID string `json:"UserID,omitempty"`

	// Email address of user
	EmailAddress string `json:"EmailAddress,omitempty"`

	// First name of user
	FirstName string `json:"FirstName,omitempty"`

	// Last name of user
	LastName string `json:"LastName,omitempty"`

	// Timestamp of last change to user
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`

	// Boolean to indicate if user is the subscriber
	IsSubscriber bool `json:"IsSubscriber,omitempty"`

	// User role (see Types)
	OrganisationRole string `json:"OrganisationRole,omitempty"`
}
