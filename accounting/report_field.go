package accounting

//ReportField is an element on a report
type ReportField struct {
	//Name is the ID of the Field
	FieldID string `json:"FieldID,omitempty" xml:"FieldID,omitempty"`
	//Description describes the Field
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`
	//Value contains the value of the Field
	Value string `json:"Value,omitempty" xml:"Value,omitempty"`
}
