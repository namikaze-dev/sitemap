package main_test

import (
	"reflect"
	"testing"

	"github.com/namikaze-dev/sitemap"
)

func TestMapToXML(t *testing.T) {
	urls := []main.URL{
		{Location: "http://www.example.com/"},
		{Location: "http://www.example.com/dogs"},
	}

	got, err := main.MapToXML(main.MapOptions{URLs: urls})
	if err != nil {
		t.Fatal(err)
	}

	want := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>`

	if string(got) != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFetchLinks(t *testing.T) {
	url := "http://www.example.com"
	got, err := main.FetchLinks(url, httpGetMock)
	if err != nil {
		t.Fatalf("unexpected error from fetch links: %v", err)
	}

	want := []main.URL{
		{Location: "http://www.example.com"},
		{Location: "http://www.example.com/dogs"},
		{Location: "http://www.example.com/cats"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func httpGetMock(url string) ([]byte, error) {
	switch url {
	case "http://www.example.com":
		return []byte(`<div><a href="/dogs">check dogs</a></div>`), nil
	case "http://www.example.com/dogs":
		return []byte(`<p><a href="http://www.example.com/cats">Click</a> to checkout cats</p>`), nil
	default:
		return []byte(`<html>time to go back <a href="http://www.example.com">home</a><a href=""></a></html>`), nil
	}
}
