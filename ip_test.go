package net

import (
	"testing"

	"github.com/toolkits/net"
)

func TestIntranetIPs(t *testing.T) {
	testIP := map[string]bool{
		"10.10.123.24":    true,
		"10.0.123.345":    false,
		"172.30.3.45":     true,
		"172.100.32.2":    false,
		"192.168.56.32":   true,
		"192.169.45.9":    false,
		"127.0.0.1":       false,
		"0.0.0.0":         false,
		"255.255.255.255": false,
	}

	for ip, res := range testIP {
		ok := net.IsIntranet(ip)
		if ok != res {
			t.Fatalf("%s not passed, the res is %v", ip, ok)
		}
	}
}
