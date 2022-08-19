package types

import "encoding/xml"

type ConstructorsStandingsMRData struct {
	XMLName        xml.Name                   `xml:"MRData"`
	Series         string                     `xml:"series,attr"`
	StandingsTable ConstructorsStandingsTable `xml:"StandingsTable"`
	Result
}

type ConstructorsStandingsTable struct {
	XMLName       xml.Name                 `xml:"StandingsTable"`
	Season        string                   `xml:"season,attr,omitempty"`
	StandingsList ConstructorStandingsList `xml:"StandingsList"`
}

type ConstructorStandingsList struct {
	XMLName             xml.Name              `xml:"StandingsList"`
	Season              string                `xml:"season,attr,omitempty"`
	Round               uint8                 `xml:"round,attr,omitempty"`
	ConstructorStanding []ConstructorStanding `xml:"ConstructorStanding"`
}

type ConstructorStanding struct {
	XMLName      xml.Name    `xml:"ConstructorStanding"`
	Position     uint8       `xml:"position,attr"`
	PositionText string      `xml:"positionText,attr"`
	Points       uint16      `xml:"points,attr"`
	Wins         uint16      `xml:"wins,attr,omitempty"`
	Constructor  Constructor `xml:"Constructor"`
}
