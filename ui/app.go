// Package ui encapsulates f1ne ui components
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/ioluas/f1ne/api/ergast"
)

// F1neUi is the main app Ui struct
type F1neUi struct {
	app         fyne.App
	mainWindow  fyne.Window
	contentGrid *fyne.Container
	cli         *ergast.Client

	standingsUi *StandingsUi
}

// NewApp returns new F1neUi struct representing application main UI component
func NewApp(cli *ergast.Client, title string) *F1neUi {
	a := app.NewWithID("ioluas/f1ne")
	mw := a.NewWindow(title)
	mw.SetMaster()

	return &F1neUi{
		app:         a,
		mainWindow:  mw,
		contentGrid: container.NewMax(),
		cli:         cli,
	}
}

// Start does basic setup of initial UI and shows main window and runs the app
func (a *F1neUi) Start(s *fyne.Size) {
	if s == nil {
		s = &fyne.Size{
			Width:  1_200,
			Height: 800,
		}
	}
	a.mainWindow.Resize(*s)

	toolbar := a.setupToolbarUi()
	borderedContainer := container.NewBorder(toolbar, nil, nil, nil, a.contentGrid)
	a.mainWindow.SetContent(borderedContainer)
	a.mainWindow.SetPadded(false)

	a.mainWindow.Show()
	a.app.Run()
}
