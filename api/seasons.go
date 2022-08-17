package api

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (c *Client) Seasons() (*SeasonsMRData, error) {
	url := fmt.Sprintf("%s/seasons?limit=100", Endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(SeasonsMRData)
	if err := c.get(url, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) SeasonsByCircuit(circuit string) (*SeasonsMRData, error) {
	url := fmt.Sprintf("%s/circuits/%s/seasons?limit=100", Endpoint, circuit)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(SeasonsMRData)
	if err := c.get(url, data); err != nil {
		return nil, err
	}
	return data, nil
}
