package api

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (c *Client) Drivers() (*DriversMRData, error) {
	url := fmt.Sprintf("%s/drivers", Endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(DriversMRData)

	err := c.get(url, data)
	if err != nil {
		return nil, err
	}
	return data, nil

	// v, err := c.getKey([]byte(url))
	// if err == nil {
	// 	err = xml.Unmarshal(*v, data)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// } else {
	// 	res, err := c.get(url)
	// 	err = xml.Unmarshal(*res, data)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// return data, nil
}
