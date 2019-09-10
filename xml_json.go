package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Data struct {
	XMLName    xml.Name `xml:"data" json:"-"`
	PersonList []Person `xml:"person" json:"people"`
}
type Person struct {
	XMLName   xml.Name `xml:"person" json:"-"`
	Firstname string   `xml:"firstname" json:"firstname"`
	Lastname  string   `xml:"lastname" json:"lastname"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}
type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func main() {
	fil, err := os.Open("sample.xml")
	if err != nil {
		fmt.Println("ERROR OPENING THE FILE")
	}
	fildat, _ := ioutil.ReadAll(fil)
	var data Data
	xml.Unmarshal([]byte(fildat), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}
