// Package widget provides custom fyne.io widgets used in here
package widget

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	types2 "github.com/ioluas/f1ne/api/ergast/types"
	"github.com/sirupsen/logrus"
)

// DriverCard is extending card widget struct
type DriverCard struct {
	widget.Card
	driver *types2.Driver
}

// var wikiMediaApiClient *http.Client

// NewDriverCard creates and returns new custom driver card widget
func NewDriverCard(driver *types2.Driver) *DriverCard {
	dc := &DriverCard{driver: driver}
	dc.SetTitle(fmt.Sprintf("%s %s", driver.GivenName, driver.FamilyName))
	dc.SetSubTitle(fmt.Sprintf("No. %d\tNationality: %s\tBorn: %s", driver.PermanentNumber, driver.Nationality,
		driver.DateOfBirth))
	u, err := url.Parse(driver.URL)
	if err != nil {
		logrus.WithError(err).Error("failed to parse driver url")
	} else {
		dc.SetContent(widget.NewHyperlinkWithStyle("Wikipedia", u, fyne.TextAlignTrailing, fyne.TextStyle{
			Italic:   true,
			TabWidth: 4,
		}))
	}
	// dc.ScrapeImage()
	return dc
}

// func (dc *DriverCard) ScrapeImage() {
// 	if wikiMediaApiClient == nil {
// 		wikiMediaApiClient = &http.Client{
// 			Timeout: time.Second * 5,
// 		}
// 	}
// 	if dc.driver.GivenName != "Lewis" {
// 		return
// 	}
//
// 	req, err := http.NewRequest(http.MethodGet, "https://en.wikipedia.org/w/api.php", nil)
// 	if err != nil {
// 		logrus.WithError(err).Error("failed to create wikimedia api request")
// 		return
// 	}
// 	q := req.URL.Query()
// 	q.Add("action", "query")
// 	q.Add("prop", "pageimages")
// 	q.Add("format", "xml")
// 	// req.URL.Query().Add("piprop", "original")
// 	q.Add("titles", fmt.Sprintf("%s %s", dc.driver.GivenName, dc.driver.FamilyName))
// 	req.URL.RawQuery = q.Encode()
// 	logrus.WithFields(logrus.Fields{"url": req.URL.String()}).Trace("got endpoint url")
//
// 	res, err := wikiMediaApiClient.Do(req)
// 	if err != nil {
// 		logrus.WithError(err).Error("failed to get wikimedia api response")
// 		return
// 	}
// 	body := res.Body
// 	defer func(body io.ReadCloser) {
// 		_ = body.Close()
// 	}(body)
// 	data := new(types.WikimediaResult)
// 	bytes, err := io.ReadAll(body)
// 	if err != nil {
// 		logrus.WithError(err).Error("failed to read response body")
// 		return
// 	}
// 	err = xml.Unmarshal(bytes, data)
// 	if err != nil {
// 		logrus.WithError(err).Error("failed to unmarshall xml response")
// 		return
// 	}
// 	thumRsc, err := fyne.LoadResourceFromURLString(data.Query.Pages.Pages[0].Original.Source)
// 	logrus.WithFields(logrus.Fields{"url": data.Query.Pages.Pages[0].Original.Source}).Trace("creating url resource")
// 	if err != nil {
// 		logrus.WithError(err).WithFields(logrus.Fields{"url": data.Query.Pages.Pages[0].Original.Source}).Error("failed to load resource from url")
// 		return
// 	}
// 	img := canvas.NewImageFromResource(thumRsc)
// 	img.Resize(fyne.NewSize(float32(data.Query.Pages.Pages[0].Original.Width),
// 		float32(data.Query.Pages.Pages[0].Original.Height)))
// 	dc.Card.SetImage(img)
// }
