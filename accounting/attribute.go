package accounting

//Attribute is an element within a Cell on a Report
type Attribute struct {
	//value of the Attribute
	Value string `json:"Value,omitempty" xml:"Value,omitempty"`
	//ID of the Attribute
	ID string `json:"Id,omitempty" xml:"Id,omitempty"`
}
