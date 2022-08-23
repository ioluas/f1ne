package types

import "encoding/xml"

// DriversStandingsMRData is api drivers standings data model from Ergast
type DriversStandingsMRData struct {
	XMLName        xml.Name              `xml:"MRData"`
	Series         string                `xml:"series,attr"`
	StandingsTable DriversStandingsTable `xml:"StandingsTable"`
	Result
}

// DriversStandingsTable is api drivers standings table data model from Ergast
type DriversStandingsTable struct {
	XMLName       xml.Name             `xml:"StandingsTable"`
	Season        string               `xml:"season,attr,omitempty"`
	StandingsList DriversStandingsList `xml:"StandingsList"`
}

// DriversStandingsList is api drivers standings list data model from Ergast
type DriversStandingsList struct {
	XMLName         xml.Name         `xml:"StandingsList"`
	Season          string           `xml:"season,attr,omitempty"`
	Round           uint8            `xml:"round,attr,omitempty"`
	DriverStandings []DriverStanding `xml:"DriverStanding"`
}

// DriverStanding is api driver standing data model from Ergast
type DriverStanding struct {
	XMLName      xml.Name    `xml:"DriverStanding"`
	Position     uint8       `xml:"position,attr"`
	PositionText string      `xml:"positionText,attr"`
	Points       uint16      `xml:"points,attr"`
	Wins         uint16      `xml:"wins,attr,omitempty"`
	Driver       Driver      `xml:"Driver"`
	Constructor  Constructor `xml:"Constructor"`
}
