package ip2asn

import "testing"
import "net"

func TestIPQueryOriginHost(t *testing.T) {
	tests := []struct {
		ip   net.IP
		host string
	}{
		{net.ParseIP("192.168.1.1").To4(), "1.1.168.192.origin.asn.cymru.com"},
		{net.ParseIP("7.3.99.230").To4(), "230.99.3.7.origin.asn.cymru.com"},
		{net.ParseIP("abcd::ffee").To16(), "e.e.f.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.d.c.b.a.origin6.asn.cymru.com"},
		{net.ParseIP("2a00:1450:400a:804::2004").To16(), "4.0.0.2.0.0.0.0.0.0.0.0.0.0.0.0.4.0.8.0.a.0.0.4.0.5.4.1.0.0.a.2.origin6.asn.cymru.com"},
	}

	for _, x := range tests {
		r, err := ipQueryHost(x.ip, Origin)
		if err != nil {
			t.Errorf("Threw error %v", err)
		} else if r != x.host {
			t.Errorf("Incorrect query host for %v, got %v, expected %v", x.ip, r, x.host)
		}
	}

	//Do a bad length test
	r, err := ipQueryHost(net.IP{1, 2, 3, 4, 5}, Origin)
	if err != ErrBadIPLength || r != "" {
		t.Errorf("Expected bad length, but got result")
	}
}
