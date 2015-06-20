package main

import (
	"flag"
	"fmt"

	"github.com/danward79/gofreegeoip/lib"
)

func main() {

	//Command line variable settings
	server := flag.String("s", "freegeoip.net", "Enter the IP address (IP:PORT) for the quiery server, leave blank for freegeoip.net.")
	ip := flag.String("a", "", "Enter the IP address for the location you wish to quiery, leave blank for this External WAN IP.")
	flag.Parse()

	loc, status := gofreegeoip.Query(*server, *ip)

	fmt.Printf("%+v", loc)
	fmt.Println(status)
}
