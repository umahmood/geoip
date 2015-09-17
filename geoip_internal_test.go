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
							"isp":"Tilaa V.O.F.",
							"area_code":"0",
							"country_code":"NL",
							"country":"Netherlands",
							"offset":"2",
							"timezone":"Europe/Amsterdam",
							"country_code3":"NLD",
							"continent_code":"EU",
							"longitude":5.75,
							"dma_code":"0",
							"asn":"AS196752",
							"latitude":52.5}`
				fmt.Fprintf(w, response)
			case "/66.102.15.255":
				response := `{"dma_code":"0",
							"ip":"66.102.15.255",
							"asn":"AS15169",
							"city":"Mountain View",
							"latitude":37.4192,
							"country_code":"US",
							"offset":"-7",
							"country":"United States",
							"region_code":"CA",
							"isp":"Google Inc.",
							"timezone":"America/Los_Angeles",
							"area_code":"0",
							"continent_code":"NA",
							"longitude":-122.0574,
							"region":"California",
							"postal_code":"94043",
							"country_code3":"USA"}`
				fmt.Fprintf(w, response)
			case "/":
				response := `{"country":"United Kingdom",
							"isp":"ARM Ltd",
							"longitude":0.1167,
							"city":"Cambridge",
							"region":"Cambridgeshire",
							"postal_code":"CB5",
							"offset":"1",
							"asn":"AS28939",
							"timezone":"Europe/London",
							"ip":"217.140.98.70",
							"continent_code":"EU",
							"country_code3":"GBR",
							"latitude":52.2,
							"dma_code":"0",
							"country_code":"GB",
							"region_code":"C3",
							"area_code":"0"}`
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
		"ip":             "2a02:2770::21a:4aff:feb3:2ee",
		"isp":            "Tilaa V.O.F.",
		"area_code":      "0",
		"country_code":   "NL",
		"country":        "Netherlands",
		"offset":         "2",
		"timezone":       "Europe/Amsterdam",
		"country_code3":  "NLD",
		"continent_code": "EU",
		"longitude":      "5.75",
		"dma_code":       "0",
		"asn":            "AS196752",
		"latitude":       "52.5",
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
		"dma_code":       "0",
		"ip":             "66.102.15.255",
		"asn":            "AS15169",
		"city":           "Mountain View",
		"latitude":       "37.4192",
		"country_code":   "US",
		"offset":         "-7",
		"country":        "United States",
		"region_code":    "CA",
		"isp":            "Google Inc.",
		"timezone":       "America/Los_Angeles",
		"area_code":      "0",
		"continent_code": "NA",
		"longitude":      "-122.0574",
		"region":         "California",
		"postal_code":    "94043",
		"country_code3":  "USA",
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
		"country":        "United Kingdom",
		"isp":            "ARM Ltd",
		"longitude":      "0.1167",
		"city":           "Cambridge",
		"region":         "Cambridgeshire",
		"postal_code":    "CB5",
		"offset":         "1",
		"asn":            "AS28939",
		"timezone":       "Europe/London",
		"ip":             "217.140.98.70",
		"continent_code": "EU",
		"country_code3":  "GBR",
		"latitude":       "52.2",
		"dma_code":       "0",
		"country_code":   "GB",
		"region_code":    "C3",
		"area_code":      "0",
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
