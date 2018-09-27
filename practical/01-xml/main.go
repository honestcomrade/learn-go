package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex ...
// create a type that we will dump the result of xml.Unmarshal into
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location ...
// create a type that will hold each loc tag that we parse
type Location struct {
	Loc string `xml:"loc"`
}

// create a method on the Location type so that we can print
// sorta like "overloading" in c++
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	// hit the URL
	res, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		panic("Err with get")
	}
	// take the response body and read it into a slice of bytes
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Err with reading bytes")
	}
	// close the connection made by the request
	res.Body.Close()

	// instantiate a SitemapIndex to hold the slice of bytes from request
	var s SitemapIndex
	// Unmarshal the xml into the SitemapIndex
	xml.Unmarshal(bytes, &s)

	// print every Location in the SitemapIndex, which will each be the parsed xml <loc> tags
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}
