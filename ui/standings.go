package ui

import (
	f1w "github.com/ioluas/f1ne/ui/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// StandingsUI represents Standings application view
type StandingsUI struct {
	items  []string
	hsplit *container.Split
	cnt    *fyne.Container
	list   *widget.List
}

func (a *F1neUI) setupStandingsUI() *StandingsUI {
	standingsUI := &StandingsUI{
		items: []string{"Drivers", "Constructors"},
		cnt:   container.NewMax(),
	}
	standingsUI.list = widget.NewList(
		func() int {
			return len(standingsUI.items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			object := co.(*widget.Label)
			object.SetText(standingsUI.items[id])
			object.Refresh()
		},
	)
	standingsUI.list.OnSelected = setStandingsListSelectionHandler(standingsUI, a)

	standingsUI.hsplit = container.NewHSplit(standingsUI.list, standingsUI.cnt)
	standingsUI.hsplit.Offset = 0.2

	return standingsUI
}

func setStandingsListSelectionHandler(su *StandingsUI, a *F1neUI) func(id widget.ListItemID) {
	return func(id widget.ListItemID) {
		switch su.items[id] {
		case "Drivers":
			d, err := a.cli.CurrentDriversStandings()
			if err != nil {
				logrus.WithError(err).Error("failed to get current drivers standings")
				dialog.ShowError(err, a.mainWindow)
				return
			}
			driversSlice := d.StandingsTable.StandingsList.DriverStandings
			acc := widget.NewAccordion()
			for _, dr := range driversSlice {
				card := f1w.NewDriverCard(&dr.Driver)
				acc.Append(widget.NewAccordionItem(card.Title, card))
			}
			su.cnt.RemoveAll()
			su.cnt.Add(acc)
			su.cnt.Refresh()
			return

		case "Constructors":
			return
		default:
			return
		}
	}
}
