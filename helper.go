package main

import "net"

func isValidIp(ip string) bool {
	return net.ParseIP(ip) != nil
}
