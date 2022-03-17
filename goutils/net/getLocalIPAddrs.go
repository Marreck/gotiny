package net

import (
	"net"
)

func GetLocalIPAddrs() (ips []net.IP, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsUnspecified() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP)
		}
	}
	return
}
