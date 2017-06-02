package accounting

//ReportAttribute is an element on a report
type ReportAttribute struct {
	//Name is the name of the Attribute
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`
	//Description describes the Attribute
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`
	//Value contains the value of the Attribute
	Value string `json:"Value,omitempty" xml:"Value,omitempty"`
}
