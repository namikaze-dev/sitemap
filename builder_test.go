package main_test

import (
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
