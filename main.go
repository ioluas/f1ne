package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"git.mills.io/prologic/bitcask"
	"github.com/ioluas/f1ne/api"
	"github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

var (
	appPath string
	db      *bitcask.Bitcask
	client  *api.Client
)

func setupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:    true,
		ForceQuote:       true,
		FullTimestamp:    true,
		DisableSorting:   true,
		QuoteEmptyFields: true,
	})
	logrus.SetOutput(os.Stdout)
	var logLvl logrus.Level
	v, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLvl = logrus.ErrorLevel
	} else {
		tmp, err := logrus.ParseLevel(v)
		if err != nil {
			logLvl = logrus.ErrorLevel
		} else {
			logLvl = tmp
		}
	}
	logrus.SetLevel(logLvl)
	logrus.Trace("logger setup complete")
}

func init() {
	setupLogger()

	// Discover app path
	ex, err := os.Executable()
	if err != nil {
		logrus.Fatalf("could not get executable: %v", err)
	}
	appPath = filepath.Dir(ex)
	logrus.WithFields(logrus.Fields{"dir": appPath}).Trace("Got executable directory")

	// Setup bitcask db, this would be used to cache API results so to save on http calls
	db, err = bitcask.Open(appPath+"/db",
		bitcask.WithMaxKeySize(2_048),
		bitcask.WithMaxValueSize(102_400),
		bitcask.WithSync(true),
	)
	if err != nil {
		logrus.Fatalf("failed to open db: %v", err)
	}
	logrus.Trace("db connection open")

	client = api.NewClient(db)
}

func main() {
	// Close db connection on exit
	defer func(db *bitcask.Bitcask) {
		if err := db.Close(); err != nil {
			logrus.Errorf("failed to close db file: %v", err)
			return
		}
		logrus.Trace("closed db")
	}(db)

	// Create new fyne app and window
	a := app.New()
	w := a.NewWindow("f1ne")
	w.Resize(fyne.NewSize(800, 600))

	grid := container.NewGridWithColumns(2)

	// d, _ := client.Drivers()
	// logrus.WithFields(logrus.Fields{"drivers": d.Drivers}).Info("Drivers list")
	// s, _ := client.SeasonsByCircuit("hungaroring")
	// logrus.WithFields(logrus.Fields{"seasons": s.Seasons, "circuit": "hungaroring"}).Info("Seasons list")

	c := container.NewBorder(setupToolbar(grid), nil, nil, nil, grid)

	w.SetContent(c)
	w.ShowAndRun()
}

func setupToolbar(c *fyne.Container) *widget.Toolbar {

	drv := widget.NewToolbarAction(resourceDrivers32Png, func() {
		logrus.Trace("clicked drivers")
		c.RemoveAll()
		c.Add(widget.NewLabel("hmm, this is not so easy!"))
		c.Refresh()
	})

	return widget.NewToolbar(drv, widget.NewToolbarSpacer())
}