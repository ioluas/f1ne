package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

func (a *F1neUi) setupStandingsUi() *container.Split {
	items, standingTableContainer := []string{"Drivers", "Constructors"}, container.NewMax()
	l := widget.NewList(
		func() int { return len(items) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, co fyne.CanvasObject) {
			object := co.(*widget.Label)
			object.SetText(items[id])
			object.Refresh()
		},
	)
	l.OnSelected = func(id widget.ListItemID) {
		switch items[id] {
		case "Drivers":
			d, err := a.cli.CurrentDriversStandings()
			if err != nil {
				logrus.WithError(err).Error("failed to get current drivers standings")
				dialog.ShowError(err, a.mw)
				break
			}
			// create new 2 grid layout to show the drivers listing on left and details of clicked driver on right
			driversSlice := d.StandingsTable.StandingsList.DriverStandings
			// c := container.NewGridWithColumns(2, list, container.NewMax())
			acc := widget.NewAccordion()
			for _, dr := range driversSlice {
				title := fmt.Sprintf("%s %s", dr.Driver.GivenName, dr.Driver.FamilyName)
				card := widget.NewCard(title, "", nil)
				acc.Append(widget.NewAccordionItem(title, card))
			}
			standingTableContainer.RemoveAll()
			standingTableContainer.Add(acc)
			standingTableContainer.Refresh()

		case "Constructors":
			// @todo.
		}
	}

	split := container.NewHSplit(l, standingTableContainer)
	split.Offset = 0.2

	return split
}
