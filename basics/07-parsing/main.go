package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Data ... refers to the root node
// which holds a list of Persons
type Data struct {
	XMLName    xml.Name `xml:"data" json:"-"` // '-' will always be omitted
	PersonList []Person `xml:"person" json:"people"`
}

// Person ...
// Holds the Individual persons data
// that will be marshalled into it
// including a reference to an Address struct
type Person struct {
	Firstname string   `xml:"firstname" json:"firstName"`
	Lastname  string   `xml:"lastname" json:"lastName"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}

// Address ...
// the address holds the address data
// that will be marshalled into it
type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func main() {
	var data Data
	rawXML, err := os.Open("sample.xml")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(rawXML)
	xml.Unmarshal(bytes, &data)
	jsonData, err := json.Marshal(data)
	fmt.Println(string(jsonData))
}
