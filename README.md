gofreegeoip
===========

[![Build Status](https://travis-ci.org/danward79/gofreegeoip.svg?branch=master)](https://travis-ci.org/danward79/gofreegeoip)

Provides access to [Freegeoip.net](https://freegeoip.net/?q=8.8.8.8) using GO

Freegeoip provides an API which can give results via, XML, JSON and CSV. This app uses the JSON interface.

freegeoip.net/{format}/{IP_or_hostname}

You are limited to *10,000* calls per hour. If this is not enough you have to provide your own server! ;-) In which case see the [Github project here.](https://github.com/fiorix/freegeoip)

###Install Assuming you have go installed and setup on your system.

```shell
go get github.com/danward79/gofreegeoip
```

######To run from source

```shell
cd $GOPATH/src/github.com/danward79/gofreegeoip
go run gofreegeoip.go -s freegeoip.net -a 8.8.8.8
```

######To compile as a genuine command line tool.

```shell
cd $GOPATH/src/github.com/danward79/gofreegeoip
go build gofreegeoip.go
```

###Command Line interface

```shell
cd $GOPATH/src/github.com/danward79/gofreegeoip
./gofreegeoip -s freegeoip.net -a 8.8.8.8
```

######Command Line Switches

-	-s server, Enter the IP address (IP:PORT) for the query server, leave blank for freegeoip.net

-	-a ip, Enter the IP address for the location you wish to quiery, leave blank for this External WAN IP

This will return two items a Location array and a status, which complies with HTTP status codes, defined in RFC 2616.

```
Location array
    {
      IP
  	CountryCode
  	CountryName
  	RegionCode  
  	RegionName  
  	City
  	ZipCode
  	TimeZone
  	Latitude
  	Longitude
  	MetroCode
    }
```

```
Status codes
        StatusOK                   = 200
        StatusForbidden            = 403 (Returned when you exceed your API allowance)
        StatusNotFound             = 404
```

The library will also return a status of 1, 2 or 3 if another error has occurred. This is accompanied by a log message in stdout

###Library call

*Example of call from code.*

```Go
package main

import (
	"fmt"
	"github.com/danward79/gofreegeoip"
)

func main() {

	loc, status := gofreegeoip.Query("freegeoip.net", "8.8.8.8")

	fmt.Printf("%+v", loc)
	fmt.Println(status)
}
```

TODO - Query Timeout handling. The ADSL internet provided where I live is terrible. So somekind of back-off is essential.
