package types

import "encoding/xml"

type Result struct {
	Url    string `xml:"url,attr,omitempty"`
	Limit  int16  `xml:"limit,attr"`
	Offset int16  `xml:"offset,attr"`
	Total  int16  `xml:"total,attr"`
}

type Driver struct {
	XMLName         xml.Name `xml:"Driver"`
	Id              string   `xml:"driverId,attr"`
	Code            string   `xml:"code,attr"`
	Url             string   `xml:"url,attr"`
	PermanentNumber uint16   `xml:"PermanentNumber"`
	GivenName       string   `xml:"GivenName"`
	FamilyName      string   `xml:"FamilyName"`
	DateOfBirth     string   `xml:"DateOfBirth"`
	Nationality     string   `xml:"Nationality"`
}

type Season struct {
	XMLName xml.Name `xml:"Season"`
	Url     string   `xml:"url,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Constructor struct {
	XMLName     xml.Name `xml:"Constructor"`
	Id          string   `xml:"constructorId,attr"`
	Url         string   `xml:"url,attr,omitempty"`
	Name        string   `xml:"Name"`
	Nationality string   `xml:"Nationality"`
}
