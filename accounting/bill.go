package accounting

type Bill struct {

	// Day of Month (0-31)
	Day string `json:"Day,omitempty"`

	// One of the following values OFFOLLOWINGMONTH/DAYSAFTERBILLDATE/OFCURRENTMONTH
	Type string `json:"Type,omitempty"`
}
