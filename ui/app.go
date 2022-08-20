package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/ioluas/f1ne/api"
)

func SetupAppUi(cli *api.Client) {
	a := app.NewWithID("f1ne")
	w := a.NewWindow("f1ne")
	w.Resize(fyne.NewSize(1_280, 1_024))
	w.SetMaster()

	grid := container.NewMax()
	toolbar := SetupToolbarUi(grid, cli)
	borderedContainer := container.NewBorder(toolbar, nil, nil, nil, grid)
	w.SetContent(borderedContainer)
	w.SetPadded(false)

	w.Show()
	a.Run()
}
