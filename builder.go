package main

import (
	"bytes"
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/namikaze-dev/link"
)

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNs   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Location string `xml:"loc"`
}

type HttpGet func(url string) ([]byte, error)

func FetchLinks(URL string, depth int, get HttpGet) ([]URL, error) {
	baseURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	visited := map[string]bool{}
	return crawl(baseURL, URL, depth, visited, get)
}

func crawl(base *url.URL, url string, depth int, visited map[string]bool, get HttpGet) ([]URL, error) {
	if depth < 0 {
		return nil, nil
	}

	fetched := []URL{{Location: url}}
	visited[url] = true

	html, err := get(url)
	if err != nil {
		return fetched, err
	}

	links, err := link.Parse(bytes.NewReader(html))
	if err != nil {
		return fetched, err
	}

	for _, link := range links {
		if norm, ok := sameDomain(base, link.Href); ok {
			if visited[norm] {
				continue
			}

			extra, err := crawl(base, norm, depth-1, visited, get)
			if err != nil {
				return fetched, err
			}
			fetched = append(fetched, extra...)
		}
	}
	return fetched, nil
}

func sameDomain(base *url.URL, curr string) (string, bool) {
	if curr == "" {
		return "", false
	}
	curr = normaliseFromBase(base, curr)
	return curr, strings.HasPrefix(curr, base.Scheme+"://"+base.Host)
}

func normaliseFromBase(base *url.URL, curr string) string {
	if curr[0] == '/' {
		return base.Scheme+"://"+base.Host+curr
	} else {
		return curr
	}
}

type MapOptions struct {
	URLs  []URL
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
