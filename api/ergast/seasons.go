package ergast

import (
	"fmt"

	"github.com/ioluas/f1ne/api/ergast/types"
	"github.com/sirupsen/logrus"
)

// Seasons returns types.SeasonsMRData result for `/seasons` or error if any encountered
func (c *Client) Seasons() (*types.SeasonsMRData, error) {
	url := fmt.Sprintf("%s/seasons?limit=100", c.endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(types.SeasonsMRData)
	if err := c.get(url, data, true, nil); err != nil {
		return nil, err
	}
	return data, nil
}

// SeasonsByCircuit returns types.SeasonsMRData result for `/circuits/{circuit}/seasons` or error if any encountered
func (c *Client) SeasonsByCircuit(circuit string) (*types.SeasonsMRData, error) {
	url := fmt.Sprintf("%s/circuits/%s/seasons?limit=100", c.endpoint, circuit)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(types.SeasonsMRData)
	if err := c.get(url, data, true, nil); err != nil {
		return nil, err
	}
	return data, nil
}
