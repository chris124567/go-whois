package main 

import (
	"log"
	whoisclient "./pkg/whoisclient"
)

func main() {
	unformatted_domain := "nic.americanfamily"
	proper_domain := whoisclient.Standardize_Domain_Name(unformatted_domain)
	log.Printf("%s", whoisclient.Whois_Query(proper_domain)) // query whois server about domain name and print the response
}
