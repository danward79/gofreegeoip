package main

import (
	"flag"
	"fmt"

	"github.com/danward79/gofreegeoip"
)

func main() {

	//Command line variables
	server := flag.String("s", "freegeoip.net", "Enter the IP address for the quiery server, leave blank for freegeoip.net.")
	ip := flag.String("a", "", "Enter the IP address for the location you wish to quiery, leave blank for this WAN IP.")
	flag.Parse()

	loc := gofreegeoip.Quiery(*server, *ip)

	fmt.Printf("%+v", loc)
}
