package ui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

func (a *F1neUi) setupToolbarUi() *widget.Toolbar {
	standings := widget.NewToolbarAction(embeddedRscStandings32Png, func() {
		logrus.Debug("clicked standings")
		a.contentGrid.RemoveAll()
		standingUi := a.setupStandingsUi()
		a.contentGrid.Add(standingUi.hsplit)
		a.contentGrid.Refresh()
	})
	drivers := widget.NewToolbarAction(embeddedRscDrivers32Png, func() {
		logrus.Debug("clicked drivers")
	})
	seasons := widget.NewToolbarAction(embeddedRscSeasons32Png, func() {
		logrus.Debug("clicked seasons")
	})

	return widget.NewToolbar(standings, widget.NewToolbarSeparator(), drivers, seasons)
}
