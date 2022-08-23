package ergast

import (
	"fmt"
	"time"

	"github.com/ioluas/f1ne/api/ergast/types"
	"github.com/sirupsen/logrus"
)

// CurrentDriversStandings returns types.DriversStandingsMRData result for `/current/driverStandings`
// or error if any encountered
func (c *Client) CurrentDriversStandings() (*types.DriversStandingsMRData, error) {
	url := fmt.Sprintf("%s/current/driverStandings", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")
	data := new(types.DriversStandingsMRData)
	ttl := time.Hour * 1

	if err := c.get(url, data, true, &ttl); err != nil {
		return nil, err
	}

	return data, nil
}

// CurrentConstructorsStandings returns types.ConstructorsStandingsMRData result for `/current/constructorStandings`
// or error if any encountered
func (c *Client) CurrentConstructorsStandings() (*types.ConstructorsStandingsMRData, error) {
	url := fmt.Sprintf("%s/current/constructorStandings", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")
	data := new(types.ConstructorsStandingsMRData)
	ttl := time.Hour * 1

	if err := c.get(url, data, true, &ttl); err != nil {
		return nil, err
	}

	return data, nil
}
