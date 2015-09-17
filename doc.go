/*
Package geoip queries location information for any IP address.

Example usage:

    package main

    import (
        "fmt"
        "log"

        "github.com/umahmood/geoip"
    )

    func main() {
        loc, err := geoip.Location("217.140.98.70")

        if err != nil {
            log.Fatalln(err)
        }

        for k, v := range loc {
            fmt.Println(k, ":", v)
        }
    }
*/
package geoip
