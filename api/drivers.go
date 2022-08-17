package api

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (c *Client) Drivers() (*DriversMRData, error) {
	url := fmt.Sprintf("%s/drivers", Endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(DriversMRData)

	if err := c.get(url, data); err != nil {
		return nil, err
	}
	return data, nil
}
