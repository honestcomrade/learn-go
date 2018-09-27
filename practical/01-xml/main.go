package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// SitemapIndex ...
// create a type that we will dump the result of xml.Unmarshal into
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	// hit the URL
	var s SitemapIndex
	var n News
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

	// instantiate a SitemapIndex to hold the slice of bytes from request
	// Unmarshal the xml into the SitemapIndex
	xml.Unmarshal(bytes, &s)

	// for each Location in the SitemapIndex
	for _, Location := range s.Locations {
		// hit that location
		res, err := http.Get(Location)
		if err != nil {
			panic("Err with get")
		}
		// take the response body and read it into a slice of bytes
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic("Err with reading bytes")
		}

		//  unmarshal the bytes into the News struct
		xml.Unmarshal(bytes, &n)
	}
	res.Body.Close()
}
