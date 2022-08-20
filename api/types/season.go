package types

import "encoding/xml"

type SeasonsMRData struct {
	XMLName xml.Name    `xml:"MRData"`
	Series  string      `xml:"series,attr"`
	Seasons SeasonsList `xml:"SeasonTable"`
	Result
}

type SeasonsList struct {
	XMLName   xml.Name `xml:"SeasonTable"`
	Season    string   `xml:"season,attr,omitempty"`
	CircuitId string   `xml:"circuitId,attr,omitempty"`
	Seasons   []Season `xml:"Season"`
}
