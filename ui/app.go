package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/ioluas/f1ne/api"
)

type F1neUi struct {
	a   fyne.App
	mw  fyne.Window
	cg  *fyne.Container
	cli *api.Client
}

// NewApp returns new F1neUi struct representing application main UI component
func NewApp(cli *api.Client, title string) *F1neUi {
	a := app.NewWithID("ioluas/f1ne")
	mw := a.NewWindow(title)
	mw.SetMaster()
	cg := container.NewMax()

	return &F1neUi{
		a:   a,
		mw:  mw,
		cg:  cg,
		cli: cli,
	}
}

// Start does basic setup of initial UI and shows main window and runs the app
func (a *F1neUi) Start(s fyne.Size) {
	a.mw.Resize(s)

	toolbar := a.setupToolbarUi()
	borderedContainer := container.NewBorder(toolbar, nil, nil, nil, a.cg)
	a.mw.SetContent(borderedContainer)
	a.mw.SetPadded(false)

	a.mw.Show()
	a.a.Run()
}
