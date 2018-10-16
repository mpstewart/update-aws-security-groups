package utils

import (
	"fmt"
	"net"
)

// Holds everything we care about for a port
type Port struct {
	Port        int64  `json:"port"`
	Protocol    string `json:"protocol"`
	Description string `json:"description"`
}

// Obtain the IP address for the DDNS hostname
func GetHomeIP() (s string) {
	config := GetConfig()
	addrs, err := net.LookupIP(config.Hostname)

	if err != nil {
		Logger.Fatalln(err)
	}

	s = fmt.Sprintf("%s/32", addrs[0])

	return
}
