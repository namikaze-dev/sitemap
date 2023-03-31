package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	base := flag.String("url", "", "base URL to build sitemap from")
	depth := flag.Int("depth", 1, "max depth of pages to crawl")
	flag.Parse()

	urls, err := FetchLinks(*base, *depth, func(u string) ([]byte, error) {
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	})
	if err != nil {
		panic(err)
	}

	xml, err := MapToXML(MapOptions{URLs: urls})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xml))
}
