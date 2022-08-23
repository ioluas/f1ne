// Package types defines data models for Wikimedia API usage in f1ne
package types

import "encoding/xml"

// WikimediaResult is api result data model from wikimedia
type WikimediaResult struct {
	XMLName xml.Name       `xml:"api"`
	Query   WikimediaQuery `xml:"query"`
}

// WikimediaQuery is api query data model from wikimedia
type WikimediaQuery struct {
	XMLName xml.Name       `xml:"query"`
	Pages   WikiMediaPages `xml:"pages"`
}

// WikiMediaPages is api pages data model from wikimedia
type WikiMediaPages struct {
	XMLName xml.Name        `xml:"pages"`
	Pages   []WikimediaPage `xml:"page"`
}

// WikimediaPage is api page data model from wikimedia
type WikimediaPage struct {
	XMLName  xml.Name           `xml:"page"`
	Title    string             `xml:"title,attr,omitempty"`
	PageId   string             `xml:"pageid,attr,omitempty"`
	NS       string             `xml:"ns,attr,omitempty"`
	Idx      string             `xml:"_idx,attr,omitempty"`
	Original WikimediaThumbnail `xml:"thumbnail"`
}

// WikimediaThumbnail is api thumbnail data model from wikimedia
type WikimediaThumbnail struct {
	XMLName xml.Name `xml:"thumbnail"`
	Source  string   `xml:"source,attr,omitempty"`
	Width   uint16   `xml:"width,attr,omitempty"`
	Height  uint16   `xml:"height,attr,omitempty"`
}
