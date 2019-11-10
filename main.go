package main 

import (
	"fmt"
	"flag"
	"os"
	whoisclient "./pkg/whoisclient"
)

func whois(domain_name string) {
	proper_domain := whoisclient.Standardize_Domain_Name(domain_name)
	fmt.Printf("%s", whoisclient.Whois_Query(proper_domain))
}
func main() {
	domain_ptr := flag.String("domain", "", "Domain to be queried.")
	flag.Parse()
	if(*domain_ptr == "") {
		fmt.Printf("WHOIS Client 1.0\nPlease supply an argument.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	whois(*domain_ptr)
}
