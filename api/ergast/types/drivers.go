package types

import "encoding/xml"

// DriversMRData is api drivers result data model from Ergast
type DriversMRData struct {
	XMLName xml.Name    `xml:"MRData"`
	Series  string      `xml:"series,attr"`
	Drivers DriversList `xml:"DriverTable"`
	Result
}

// DriversList is api drivers list data model from Ergast
type DriversList struct {
	XMLName xml.Name `xml:"DriverTable"`
	Season  string   `xml:"season,attr"`
	Drivers []Driver `xml:"Driver"`
}
