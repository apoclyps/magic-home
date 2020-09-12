package magichome

import (
	"bytes"
	"fmt"
	"net"
)

var (
	minIP = net.ParseIP("10.0.0.0")
	maxIP = net.ParseIP("192.168.255.255")
)

// IsPrivateIpv4 validate's an IP address exists within the private range
// on a local network.
func IsPrivateIpv4(ip string) bool {
	ipAddr := net.ParseIP(ip)

	if ipAddr.To4() == nil {
		fmt.Printf("%v is not an IPv4 address\n", ipAddr)
		return false
	}

	if bytes.Compare(ipAddr, minIP) >= 0 && bytes.Compare(ipAddr, maxIP) <= 0 {
		fmt.Printf("%v is between %v and %v\n", ipAddr, minIP, maxIP)
		return true
	}

	fmt.Printf("%v is NOT between %v and %v\n", ipAddr, minIP, maxIP)

	return false
}
