package ergast

import (
	"fmt"

	"github.com/ioluas/f1ne/api/ergast/types"
	"github.com/sirupsen/logrus"
)

// Drivers returns types.DriversMRData result for `/drivers` or error if any encountered
func (c *Client) Drivers() (*types.DriversMRData, error) {
	url := fmt.Sprintf("%s/drivers", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(types.DriversMRData)

	if err := c.get(url, data, true, nil); err != nil {
		return nil, err
	}
	return data, nil
}
