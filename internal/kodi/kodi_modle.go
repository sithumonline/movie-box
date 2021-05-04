package kodi

import "encoding/xml"

type Movie struct {
	XMLName       xml.Name `xml:"movie"`
	Title         string   `xml:"title"`
	OriginalTitle string   `xml:"originaltitle"`
	UserRating    float32  `xml:"userrating"`
	Plot          string   `xml:"plot"`
	Runtime       int      `xml:"runtime"`
	Thumb         string   `xml:"thumb"`
}
