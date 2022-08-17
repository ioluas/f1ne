package api

import "encoding/xml"

const Endpoint = "https://ergast.com/api/f1"

// DriversList

type DriversMRData struct {
	XMLName xml.Name    `xml:"MRData"`
	Series  string      `xml:"series,attr"`
	Url     string      `xml:"url,attr,omitempty"`
	Limit   int16       `xml:"limit,attr"`
	Offset  int16       `xml:"offset,attr"`
	Total   int16       `xml:"total,attr"`
	Drivers DriversList `xml:"DriverTable"`
}

type DriversList struct {
	XMLName xml.Name `xml:"DriverTable"`
	Season  string   `xml:"season,attr"`
	Drivers []Driver `xml:"Driver"`
}

type Driver struct {
	XMLName         xml.Name `xml:"Driver"`
	Id              string   `xml:"driverId,attr"`
	Code            string   `xml:"code,attr"`
	Url             string   `xml:"url,attr"`
	PermanentNumber uint8    `xml:"PermanentNumber"`
	GivenName       string   `xml:"GivenName"`
	FamilyName      string   `xml:"FamilyName"`
	DateOfBirth     string   `xml:"DateOfBirth"`
	Nationality     string   `xml:"Nationality"`
}

// Seasons

type SeasonsMRData struct {
	XMLName xml.Name    `xml:"MRData"`
	Series  string      `xml:"series,attr"`
	Url     string      `xml:"url,attr,omitempty"`
	Limit   int16       `xml:"limit,attr"`
	Offset  int16       `xml:"offset,attr"`
	Total   int16       `xml:"total,attr"`
	Seasons SeasonsList `xml:"SeasonTable"`
}

type SeasonsList struct {
	XMLName   xml.Name `xml:"SeasonTable"`
	Season    string   `xml:"season,attr,omitempty"`
	CircuitId string   `xml:"circuitId,attr,omitempty"`
	Seasons   []Season `xml:"Season"`
}

type Season struct {
	XMLName xml.Name `xml:"Season"`
	Url     string   `xml:"url,attr,omitempty"`
	Value   string   `xml:",chardata"`
}
