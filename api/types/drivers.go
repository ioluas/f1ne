package types

import "encoding/xml"

type DriversMRData struct {
	XMLName xml.Name    `xml:"MRData"`
	Series  string      `xml:"series,attr"`
	Drivers DriversList `xml:"DriverTable"`
	Result
}

type DriversList struct {
	XMLName xml.Name `xml:"DriverTable"`
	Season  string   `xml:"season,attr"`
	Drivers []Driver `xml:"Driver"`
}
