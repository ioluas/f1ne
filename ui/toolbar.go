package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/ioluas/f1ne/api"
	"github.com/sirupsen/logrus"
)

func SetupToolbarUi(c *fyne.Container, cli *api.Client) *widget.Toolbar {

	standings := widget.NewToolbarAction(embeddedRscStandings32Png, func() {
		logrus.Trace("clicked standings")
		standingSplit := SetupStandingsUi(cli)
		c.RemoveAll()
		c.Add(standingSplit)
		c.Refresh()
	})
	drivers := widget.NewToolbarAction(embeddedRscDrivers32Png, func() {
		logrus.Trace("clicked drivers")
	})
	seasons := widget.NewToolbarAction(embeddedRscSeasons32Png, func() {
		logrus.Trace("clicked seasons")
	})

	return widget.NewToolbar(standings, widget.NewToolbarSeparator(), drivers, seasons)
}
