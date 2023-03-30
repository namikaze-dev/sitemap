package main

import "encoding/xml"

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNs   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Location string `xml:"loc"`
}

type MapOptions struct {
	URLs []URL
	XMLNs string
}

func MapToXML(opt MapOptions) ([]byte, error) {
	if opt.XMLNs == "" {
		opt.XMLNs = "http://www.sitemaps.org/schemas/sitemap/0.9"
	}
	urlset := UrlSet{URLs: opt.URLs, XMLNs: opt.XMLNs}

	data, err := xml.MarshalIndent(urlset, "", "  ")
	header := []byte(xml.Header)
	return append(header, data...), err
}
