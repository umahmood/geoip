# GeoIP

GeoIP is a command line tool, which allows you to query location information from 
any IP address. GeoIP also exposes a library so you can also use it in your 
application. 

**Note:** This tool uses the http://www.telize.com/ REST API to provide query 
location information.

# Installation

> $ go get github.com/umahmood/geoip

> $ cd $GOPATH/src/github.com/umahmood/geoip

> $ go test 

> $ go install ./...

# Usage

> $ geoip -ip 66.102.15.255

    country : United States
    area_code : 0
    city : Mountain View
    continent_code : NA
    asn : AS15169
    timezone : America/Los_Angeles
    country_code3 : USA
    country_code : US
    longitude : -122.0574
    latitude : 37.4192
    dma_code : 0
    isp : Google Inc.
    offset : -7
    region_code : CA
    ip : 66.102.15.255
    region : California
    postal_code : 94043

You can also query IPv6 addresses:

> $ geoip -ip 2a02:2770::21a:4aff:feb3:2ee

Providing no 'ip' flag, geoip will query your IP address:

> $ geoip

# Documenation

> http://godoc.org/github.com/umahmood/geoip

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

