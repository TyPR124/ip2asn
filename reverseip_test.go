package ip2asn

import (
	"net"
	"testing"
)

func TestReverseIP4(t *testing.T) {
	tests := []struct {
		ip      string
		reverse string
	}{
		{"192.168.1.1", "1.1.168.192"},
		{"10.5.2.66", "66.2.5.10"},
		{"4.2.2.2", "2.2.2.4"},
		{"0.0.0.0", "0.0.0.0"},
		{"1.2.3.4", "4.3.2.1"},
	}

	for _, x := range tests {
		if reverseIP4(net.ParseIP(x.ip).To4()) != x.reverse {
			t.Errorf("IP %v reversed %v, expected %v", x.ip, reverseIP4(net.ParseIP(x.ip).To4()), x.reverse)
		}
	}
}
