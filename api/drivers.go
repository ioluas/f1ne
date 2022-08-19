package api

import (
	"fmt"

	"github.com/ioluas/f1ne/api/types"
	"github.com/sirupsen/logrus"
)

func (c *Client) Drivers() (*types.DriversMRData, error) {
	url := fmt.Sprintf("%s/drivers", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(types.DriversMRData)

	if err := c.get(url, data, true, nil); err != nil {
		return nil, err
	}
	return data, nil
}
