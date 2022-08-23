package types

import "encoding/xml"

// ConstructorsStandingsMRData is api constructors standings data model from Ergast
type ConstructorsStandingsMRData struct {
	XMLName        xml.Name                   `xml:"MRData"`
	Series         string                     `xml:"series,attr"`
	StandingsTable ConstructorsStandingsTable `xml:"StandingsTable"`
	Result
}

// ConstructorsStandingsTable is api constructors standings table data model from Ergast
type ConstructorsStandingsTable struct {
	XMLName       xml.Name                 `xml:"StandingsTable"`
	Season        string                   `xml:"season,attr,omitempty"`
	StandingsList ConstructorStandingsList `xml:"StandingsList"`
}

// ConstructorStandingsList is api constructors standings list data model from Ergast
type ConstructorStandingsList struct {
	XMLName             xml.Name              `xml:"StandingsList"`
	Season              string                `xml:"season,attr,omitempty"`
	Round               uint8                 `xml:"round,attr,omitempty"`
	ConstructorStanding []ConstructorStanding `xml:"ConstructorStanding"`
}

// ConstructorStanding is api constructor standing data model from Ergast
type ConstructorStanding struct {
	XMLName      xml.Name    `xml:"ConstructorStanding"`
	Position     uint8       `xml:"position,attr"`
	PositionText string      `xml:"positionText,attr"`
	Points       uint16      `xml:"points,attr"`
	Wins         uint16      `xml:"wins,attr,omitempty"`
	Constructor  Constructor `xml:"Constructor"`
}
