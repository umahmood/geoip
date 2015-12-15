package geoip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var baseURI = "https://freegeoip.net/json/"

// Location queries location information of an IP address. 'ip' can be a IPv4
// or IPv6 address. It is the users job to make sure 'ip' is a valid IP address.
func Location(ip string) (map[string]string, error) {
	uri := fmt.Sprintf("%s%s", baseURI, ip)
	body, err := performRequest(uri)
	if err != nil {
		return nil, err
	}
	data, err := extractJSON(body)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for k, v := range data {
		if k == "latitude" || k == "longitude" || k == "metro_code" {
			m[k] = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		} else {
			m[k] = v.(string)
		}
	}
	return m, nil
}

func performRequest(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		e := fmt.Sprintf("http %d from %s", res.StatusCode, baseURI)
		return nil, errors.New(e)
	}
	return body, nil
}

func extractJSON(jsonBlob []byte) (map[string]interface{}, error) {
	var j map[string]interface{}
	err := json.Unmarshal(jsonBlob, &j)
	if err != nil {
		return nil, err
	}
	return j, nil
}
