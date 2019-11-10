package whoisclient

import (
	"strings"
	"log"
	"golang.org/x/net/publicsuffix"
)

func tld_split(domain_name string) string { // get last split of '.', which should be the TLD
	ss := strings.Split(domain_name, ".")
	return ss[len(ss)-1]
}

func Standardize_Domain_Name(unformatted_domain_name string) string { // 	// standardize domain name: remove subdomains, lower case everything
	proper_domain, err := publicsuffix.EffectiveTLDPlusOne(strings.ToLower(unformatted_domain_name))
	if err != nil {
		log.Printf("Failed to parse domain name %s", unformatted_domain_name)
		panic(err)
	}
	return proper_domain
}