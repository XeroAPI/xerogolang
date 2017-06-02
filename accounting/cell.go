package accounting

//Cell is an element within a Row on a Report
type Cell struct {
	//Value is the value within a cell
	Value string `json:"Value,omitempty" xml:"Value,omitempty"`
	//Attributes are Attributes of a cell
	Attributes *[]Attribute `json:"Attributes,omitempty" xml:"Attributes>Attribute,omitempty"`
}
