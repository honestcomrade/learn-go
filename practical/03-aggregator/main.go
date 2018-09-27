package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// SitemapIndex ...
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News ...
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMap ...
type NewsMap struct {
	Keyword  string
	Location string
}

// NewsAggPage ...
type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap)
	res, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		panic("Err with get")
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Err with reading bytes")
	}
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations {
		res, err := http.Get(Location)
		if err != nil {
			panic("Err with get")
		}
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic("Err with reading bytes")
		}
		xml.Unmarshal(bytes, &n)
		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	p := NewsAggPage{Title: "Example News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("agg-tpl.gohtml")
	fmt.Println(t.Execute(w, p))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index page</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":4000", nil)
}
