package api

import (
	"fmt"

	"github.com/ioluas/f1ne/api/types"
	"github.com/sirupsen/logrus"
)

func (c *Client) Seasons() (*types.SeasonsMRData, error) {
	url := fmt.Sprintf("%s/seasons?limit=100", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(types.SeasonsMRData)
	if err := c.get(url, data, true, nil); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) SeasonsByCircuit(circuit string) (*types.SeasonsMRData, error) {
	url := fmt.Sprintf("%s/circuits/%s/seasons?limit=100", c.endpoint, circuit)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(types.SeasonsMRData)
	if err := c.get(url, data, true, nil); err != nil {
		return nil, err
	}
	return data, nil
}
