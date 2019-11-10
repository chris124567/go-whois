package whoisclient

import (
	"net"
	"fmt"
	"log"
	"io/ioutil"
)

func Whois_Query(domain_name string) string { // TODO: make asynchronous, improve performance (something to do with handling of strings?)
	tld := tld_split(domain_name)
	log.Printf("Parsed TLD (%s) from domain %s", tld, domain_name)

	server_addr, err := net.ResolveIPAddr("ip", Info_Table[tld].whois_server) // use DNS to resolve hostname of WHOIS server.  TODO: use cache
	if err != nil {
		log.Printf("Failed to resolve IP of WHOIS Server for TLD %s", tld)
		return "" // error
	}

	formatted_addr := net.JoinHostPort(server_addr.IP.String(), "43") // format IP and port in golangs format
	conn, err := net.Dial("tcp", formatted_addr) // actually connect to server
	if err != nil {
		log.Printf("Failed to connect to WHOIS server via TCP")
		return "" // error
	}

	log.Printf("Successfully connected to WHOIS server (%s)", formatted_addr)

	message := fmt.Sprintf(Info_Table[tld].query_format + "\r\n", domain_name) // add carriage return to end of message
	conn.Write([]byte(message)) // send message to server
	log.Printf("Sending message %s to WHOIS server", message)

	received_data, err := ioutil.ReadAll(conn) // receive data from server until connection closed (EOF)
	if err != nil {
		log.Printf("Failed to received response from server")
		return "" // error
	}

	return string(received_data)
}