// This tool uses Freegeoip to determine:
// - External IP address
// - Country, Region, City, Zip Code
// - Lattitude and Longitude
//
// See http://www.freegeoip.net for details

// Queries have the following format, ip address is optional. If it is not included, the IP of the query location is used.
// https://freegeoip.net/json/ip address

// JSON Data has the following format
// {"ip":"183.122.986.76","country_code":"AU","country_name":"Australia","region_code":"VIC","region_name":"Victoria","city":"Dandenong West","zip_code":"3175","time_zone":"Australia/Melbourne","latitude":-38.012,"longitude":145.216,"metro_code":0}

package gofreegeoip

import (
	"encoding/json"
	"strings"

	"io/ioutil"
	"log"
	"net/http"
)

// Location provides a type definition for the JSON structure returned from freegeoip.
type Location struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

// doQuery  checks the specified server for the location of the specified ip address.
func doQuery(s string, i string) (Location, int) {
	var loc Location

	res, err := http.Get(assembleURL(s, i))
	if err != nil {
		log.Printf("Error http.Get %s", assembleURL(s, i))
		log.Print(err)
		return loc, 1
	}
	defer res.Body.Close()

	st := res.StatusCode

	if st == http.StatusOK {

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Print("Error ioutil.ReadAll")
			log.Print(err)
			return loc, 2
		}

		err = json.Unmarshal(data, &loc)
		if err != nil {
			log.Print("Error json.Unmarshal")
			log.Print(err)
			return loc, 3
		}

	}

	return loc, st
}

//assembleURL creates a correctly formatted URL
func assembleURL(s string, i string) string {

	if !strings.Contains(s, "http") {
		s = "https://" + s
	}

	return s + "/json/" + i
}

// Query provides an external interface to gofreegeoip
func Query(s string, i string) (Location, int) {
	loc, st := doQuery(s, i)
	return loc, st
}
