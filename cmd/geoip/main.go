package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/umahmood/geoip"
)

// command line flags.
var ip string

func init() {
	flag.StringVar(&ip, "ip", "", "IP to geo locate can be v4 or v6 address.")
	flag.Parse()

	if ip != "" {
		addr := net.ParseIP(ip)
		if addr == nil {
			log.Fatalln("not a valid IP address.")
		}
	}
}

func main() {
	loc, err := geoip.Location(ip)
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range loc {
		fmt.Println(k, ":", v)
	}
}
