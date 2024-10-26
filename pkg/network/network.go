package network

import (
	"fmt"
	"net"
)

/// [Method] to find the local IP address of the machine
/// [Ignores] unsuitable addresses (loopback, IPv6)
///
/// Returns the local IP address of the machine

func GetLocalNetworkIP() (string, error) {
	///
	///
	///`addrs` contains all network interfaces on your computer (WiFi, Ethernet, Loopback, etc.)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	///
	///
	///Each `addr` represents one network interface
	for _, addr := range addrs {
		///
		///
		///`ipnet` is the IP address of the network interface
		///
		///`ok` is a boolean that is true if the IP address is not a loopback address
		///
		///This is a type assertion to convert `addr` to `*net.IPNet`
		///
		///`ipnet.IP.To4()` checks if the IP address is an IPv4 address
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no local IP address found")
}
