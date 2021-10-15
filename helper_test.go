package main

import "testing"

func TestIsValidIp(t *testing.T) {
	table := []string{
		"",
		"dfsdfdsfds",
		"dfssd dfds dfds dsfds",
		"192.168.1.A",
		"192.168.1.256",
		"192.168.1.1.1",
	}

	for _, ip := range table {
		result := isValidIp(ip)
		if result {
			t.Errorf("incorrect for isValidIp(%s), got: %t, want: %t.", ip, result, false)
		}
	}

	table = []string{
		"192.168.1.1",
		"172.168.1.1",
		"127.0.0.1",
		"12.168.12.12",
		"8.8.8.8",
		"108.8.8.8",
	}

	for _, ip := range table {
		result := isValidIp(ip)
		if result != true {
			t.Errorf("incorrect for isValidIp(%s), got: %t, want: %t.", ip, result, false)
		}
	}
}
