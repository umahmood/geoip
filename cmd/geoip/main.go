package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/umahmood/geoip"
)

// command line flags.
var ip string

func init() {
	flag.StringVar(&ip, "ip", "", "IP to geo locate can be v4/v6 address or domain name.")
	flag.Parse()
}

func main() {
	loc, err := geoip.Location(ip)
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range loc {
		if v == "" {
			fmt.Println(k, ": n/a")
		} else {
			fmt.Println(k, ":", v)
		}
	}
}
