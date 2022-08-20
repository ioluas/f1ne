package api

import (
	"fmt"
	"time"

	"github.com/ioluas/f1ne/api/types"
	"github.com/sirupsen/logrus"
)

func (c *Client) CurrentDriversStandings() (*types.DriversStandingsMRData, error) {
	url := fmt.Sprintf("%s/current/driverStandings", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")
	data := new(types.DriversStandingsMRData)
	var ttl *time.Duration
	tmp := time.Hour * 1
	ttl = &tmp

	if err := c.get(url, data, true, ttl); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CurrentConstructorsStandings() (*types.ConstructorsStandingsMRData, error) {
	url := fmt.Sprintf("%s/current/constructorStandings", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")
	data := new(types.ConstructorsStandingsMRData)
	var ttl *time.Duration
	tmp := time.Hour * 1
	ttl = &tmp

	if err := c.get(url, data, true, ttl); err != nil {
		return nil, err
	}

	return data, nil
}
