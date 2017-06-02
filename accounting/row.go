package accounting

//Row is a Row on a Report
type Row struct {
	//RowType indicates what type of Row this is
	RowType string `json:"RowType,omitempty" xml:"RowType,omitempty"`
	//Title describes the row
	Title string `json:"Title,omitempty" xml:"Title,omitempty"`
	//Rows is a collection of other rows that can be nested beneath the row
	Rows *[]Row `json:"Rows,omitempty" xml:"Rows>Row,omitempty"`
	//Cells is a collection of cells that can be nested beneath the row
	Cells *[]Cell `json:"Cells,omitempty" xml:"Cells>Cell,omitempty"`
}
