/*
Package geoip queries location information for any IP address or domain name.

Example usage:

    package main

    import (
        "fmt"
        "log"

        "github.com/umahmood/geoip"
    )

    func main() {

        t := "217.140.98.70" // or t := "bbc.co.uk"

        loc, err := geoip.Location(t)

        if err != nil {
            log.Fatalln(err)
        }

        for k, v := range loc {
            fmt.Println(k, ":", v)
        }
    }
*/
package geoip
