// Package main app
package main

import (
	"os"
	"path/filepath"

	"github.com/ioluas/f1ne/api/ergast"
	"github.com/ioluas/f1ne/cache"
	"github.com/ioluas/f1ne/ui"
	"github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

var (
	appPath string
	cacheDb *cache.Cache
	client  *ergast.Client
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
	v, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logrus.SetLevel(logrus.ErrorLevel)
	} else if tmp, err := logrus.ParseLevel(v); err != nil {
		logrus.SetLevel(logrus.ErrorLevel)
	} else {
		logrus.SetLevel(tmp)
	}
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
	cacheDb, err = cache.NewCache(appPath)
	if err != nil {
		logrus.Fatal("failed to start")
	}

	// Setup Ergast api client
	client = ergast.NewClient(cacheDb)
}

func main() {
	// Close cache db connection on exit
	defer func(c *cache.Cache) {
		if err := c.Close(); err != nil {
			logrus.Errorf("failed to close db file: %v", err)
			return
		}
		logrus.Trace("closed db")
	}(cacheDb)

	// Create new Ui and start it
	f1ne := ui.NewApp(client, "f1ne")
	f1ne.Start(nil)
}
