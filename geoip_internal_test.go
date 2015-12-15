package geoip

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestServer() *httptest.Server {
	s := make(chan bool)
	var ts *httptest.Server
	go func() {
		ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			p := r.URL.Path
			switch p {
			case "/2a02:2770::21a:4aff:feb3:2ee":
				response := `{"ip":"2a02:2770::21a:4aff:feb3:2ee",
							"country_code":"NL",
							"country_name":"Netherlands",
							"region_code":"",
							"region_name":"",
							"city":"",
							"zip_code":"",
							"time_zone":"Europe/Amsterdam",
							"latitude":52.25,
							"longitude":5.75,
							"metro_code":0}`
				fmt.Fprintf(w, response)
			case "/66.102.15.255":
				response := `{"ip":"66.102.15.255",
							"country_code":"US",
							"country_name":"United States",
							"region_code":"CA",
							"region_name":"California",
							"city":"Mountain View",
							"zip_code":"94043",
							"time_zone":"America/Los_Angeles",
							"latitude":37.4192,
							"longitude":-122.0574,
							"metro_code":807}`
				fmt.Fprintf(w, response)
			case "/":
				response := `{"ip":"217.140.98.70",
							"country_code":"GB",
							"country_name":"United Kingdom",
							"region_code":"ENG",
							"region_name":"England",
							"city":"Saint Neots",
							"zip_code":"CB5",
							"time_zone":"Europe/London",
							"longitude":0.1167,
							"latitude":52.2,
							"metro_code":0}`
				fmt.Fprintf(w, response)
			case "/github.com":
				response := `{"ip":"192.30.252.130",
							"country_code":"US",
							"country_name":"United States",
							"region_code":"CA",
							"region_name":"California",
							"city":"San Francisco",
							"zip_code":"94107",
							"time_zone":"America/Los_Angeles",
							"latitude":37.7697,
							"longitude":-122.3933,
							"metro_code":807}`
				fmt.Fprintf(w, response)
			case "/abcxyz":
				response := `404 page not found`
				fmt.Fprintf(w, response)
			}
		}))
		baseURI = ts.URL + "/"
		s <- true
	}()
	_ = <-s
	return ts
}

// TestLocationWithIP6 test location method using a IPv6 address.
func TestLocationWithIP6(t *testing.T) {
	ts := startTestServer()
	defer ts.Close()

	input := "2a02:2770::21a:4aff:feb3:2ee"

	want := map[string]string{
		"ip":           "2a02:2770::21a:4aff:feb3:2ee",
		"country_code": "NL",
		"country_name": "Netherlands",
		"region_code":  "",
		"region_name":  "",
		"city":         "",
		"zip_code":     "",
		"time_zone":    "Europe/Amsterdam",
		"latitude":     "52.25",
		"longitude":    "5.75",
		"metro_code":   "0",
	}

	got, err := Location(input)

	if err != nil {
		t.Errorf("Location error: method returned error %s", err)
	}

	for k, v := range got {
		if want[k] != v {
			t.Errorf("Location data: got %s - %s want %s - %s", k, v, k, want[k])
		}
	}

}

// TestLocationWithIP6 test location method using a IPv4 address.
func TestLocationWithIP4(t *testing.T) {
	ts := startTestServer()
	defer ts.Close()

	input := "66.102.15.255"

	want := map[string]string{
		"ip":           "66.102.15.255",
		"country_code": "US",
		"country_name": "United States",
		"region_code":  "CA",
		"region_name":  "California",
		"city":         "Mountain View",
		"zip_code":     "94043",
		"time_zone":    "America/Los_Angeles",
		"latitude":     "37.4192",
		"longitude":    "-122.0574",
		"metro_code":   "807",
	}

	got, err := Location(input)

	if err != nil {
		t.Errorf("Location error: method returned error %s", err)
	}

	for k, v := range got {
		if want[k] != v {
			t.Errorf("Location data: got %s - %s want %s - %s", k, v, k, want[k])
		}
	}
}

// TestLocationWithIP6 test location method using a blank IP address.
func TestLocationWithNoIPProvided(t *testing.T) {
	ts := startTestServer()
	defer ts.Close()

	want := map[string]string{
		"ip":           "217.140.98.70",
		"country_code": "GB",
		"country_name": "United Kingdom",
		"region_code":  "ENG",
		"region_name":  "England",
		"city":         "Saint Neots",
		"zip_code":     "CB5",
		"time_zone":    "Europe/London",
		"longitude":    "0.1167",
		"latitude":     "52.2",
		"metro_code":   "0",
	}

	got, err := Location("")

	if err != nil {
		t.Errorf("Location error: method returned error %s", err)
	}

	for k, v := range got {
		if want[k] != v {
			t.Errorf("Location data: got %s - %s want %s - %s", k, v, k, want[k])
		}
	}
}

// TestLocationWithDomain test location with a domain i.e. github.com
func TestLocationWithDomain(t *testing.T) {
	ts := startTestServer()
	defer ts.Close()

	input := "github.com"

	want := map[string]string{
		"ip":           "192.30.252.130",
		"country_code": "US",
		"country_name": "United States",
		"region_code":  "CA",
		"region_name":  "California",
		"city":         "San Francisco",
		"zip_code":     "94107",
		"time_zone":    "America/Los_Angeles",
		"latitude":     "37.7697",
		"longitude":    "-122.3933",
		"metro_code":   "807",
	}

	got, err := Location(input)

	if err != nil {
		t.Errorf("Location error: method returned error %s", err)
	}

	for k, v := range got {
		if want[k] != v {
			t.Errorf("Location data: got %s - %s want %s - %s", k, v, k, want[k])
		}
	}

}

// TestLocation404 test location with bad host or ip
func TestLocation404(t *testing.T) {

}
