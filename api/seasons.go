package api

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (c *Client) Seasons() (*SeasonsMRData, error) {
	url := fmt.Sprintf("%s/seasons?limit=100", Endpoint)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(SeasonsMRData)
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

func (c *Client) SeasonsByCircuit(circuit string) (*SeasonsMRData, error) {
	url := fmt.Sprintf("%s/circuits/%s/seasons?limit=100", Endpoint, circuit)
	logrus.WithFields(logrus.Fields{"endpoint": url}).Trace("got endpoint url")

	data := new(SeasonsMRData)
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
