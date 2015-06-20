package gofreegeoip

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"
)

//TestURLForm
func TestURLForm(t *testing.T) {
	expected := "https://freegeoip.org/json/"
	got := assembleURL("freegeoip.org", "")
	if got != expected {
		t.Fatal("Expected: ", expected, ", Got: ", got)
	}
}

//TestURLandIPForm
func TestURLandIPForm(t *testing.T) {
	expected := "https://freegeoip.org/json/8.8.8.8"
	got := assembleURL("freegeoip.org", "8.8.8.8")
	if got != expected {
		t.Fatal("Expected: ", expected, ", Got: ", got)
	}
}

// TestSimpleCommand URL only no IP
func TestSimpleCommand(t *testing.T) {
	fakeJSON := `{"ip":"145.114.17.16","country_code":"AU","country_name":"Australia","region_code":"VIC","region_name":"Victoria","city":"Melbourne","zip_code":"8001","time_zone":"Australia/Melbourne","latitude":-37.814,"longitude":144.963,"metro_code":0}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fakeJSON)
	}))
	defer ts.Close()

	loc, status := Query(ts.URL, "")
	if status != 200 {
		t.Fatal("Expected StatusCode 200. Got: ", status)
	}

	var fakeLoc Location
	err := json.Unmarshal([]byte(fakeJSON), &fakeLoc)
	if err != nil {
		t.Fatal(err)
	}

	if loc != fakeLoc {
		t.Fatal("Result not as expected. Got: ", loc)
	}
}

//TestFullCommand URL & IP
func TestFullCommand(t *testing.T) {

	fakeJSON := `{"ip":"8.8.8.8","country_code":"US","country_name":"United States","region_code":"CA","region_name":"California","city":"Mountain View","zip_code":"94040","time_zone":"America/Los_Angeles","latitude":37.386,"longitude":-122.084,"metro_code":807}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fakeJSON)
	}))
	defer ts.Close()

	loc, status := Query(ts.URL, "8.8.8.8")
	if status != 200 {
		t.Fatal("Expected StatusCode 200. Got: ", status)
	}

	var fakeLoc Location
	err := json.Unmarshal([]byte(fakeJSON), &fakeLoc)
	if err != nil {
		t.Fatal(err)
	}

	if loc != fakeLoc {
		t.Fatal("Result not as expected. Got: ", loc)
	}
}

//TestRequestLimit When quota is exceeded StatusForbidden 403 should be returned by the server.
func TestRequestLimit(t *testing.T) {

	fakeResp := "Quota exceeded"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
		fmt.Fprint(w, fakeResp)
	}))
	defer ts.Close()

	_, status := Query(ts.URL, "")
	if status != 403 {
		t.Fatal("Expected StatusCode 403. Got: ", status)
	}

}
