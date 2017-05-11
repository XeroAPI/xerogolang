package accounting

type Employee struct {

	// The Xero identifier for an employee e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	EmployeeID string `json:"EmployeeID,omitempty"`

	// Current status of an employee – see contact status types
	Status string `json:"Status,omitempty"`

	// First name of an employee (max length = 255)
	FirstName string `json:"FirstName,omitempty"`

	// Last name of an employee (max length = 255)
	LastName string `json:"LastName,omitempty"`
}
