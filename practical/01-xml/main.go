package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	res, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		panic("Err with get")
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Err with reading bytes")
	}
	res.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}
