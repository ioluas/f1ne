package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ioluas/f1ne/api"
)

func SetupStandingsUi(c *api.Client) *container.Split {
	list := []string{"Drivers", "Constructors"}
	l := widget.NewList(
		func() int { return len(list) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).Text = list[id]
			co.Refresh()
		},
	)
	l.OnSelected = func(id widget.ListItemID) {
		if list[id] == "Drivers" {
			body := ""
			d, err := c.CurrentDriversStandings()
			if err != nil {
				body = err.Error()
			} else {
				for _, ds := range d.StandingsTable.StandingsList.DriverStandings {
					body += fmt.Sprintf("Name: %s Nationality: %s Id: %s\n", ds.Driver.GivenName+" "+ds.Driver.FamilyName,
						ds.Driver.Nationality, ds.Driver.Id)
				}
			}
			return
		}
		if list[id] == "Constructors" {

			return
		}
	}

	standingTableContainer := container.NewMax()

	split := container.NewHSplit(l, standingTableContainer)
	split.Offset = 0.2

	return split
}
