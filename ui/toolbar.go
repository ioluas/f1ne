package ui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

func (a *F1neUi) setupToolbarUi() *widget.Toolbar {

	standings := widget.NewToolbarAction(embeddedRscStandings32Png, func() {
		logrus.Trace("clicked standings")
		standingSplit := a.setupStandingsUi()
		a.cg.RemoveAll()
		a.cg.Add(standingSplit)
		a.cg.Refresh()
	})
	drivers := widget.NewToolbarAction(embeddedRscDrivers32Png, func() {
		logrus.Trace("clicked drivers")
	})
	seasons := widget.NewToolbarAction(embeddedRscSeasons32Png, func() {
		logrus.Trace("clicked seasons")
	})

	return widget.NewToolbar(standings, widget.NewToolbarSeparator(), drivers, seasons)
}
