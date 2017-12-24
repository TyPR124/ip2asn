package ip2asn

import "testing"
import "net"

func TestIPQueryOrigin(t *testing.T) {
	tests := []struct {
		ip     net.IP
		result PrefixInfo
	}{
		{
			net.IP{4, 2, 2, 2},
			PrefixInfo{
				Prefix: net.IPNet{
					IP:   net.IP{4, 0, 0, 0},
					Mask: net.CIDRMask(9, 32)},
				ASN:         []ASNumber{ASNumber(3356)},
				Relation:    Origin,
				CountryCode: "US",
				Repository:  "arin",
				AllocDate:   "1992-12-01",
			},
		},
	}

	for _, x := range tests {
		res, err := QueryIP(x.ip, Origin)
		if err != nil {
			t.Errorf("Received error %v", err)
		}
		if res.Prefix.String() != x.result.Prefix.String() || //String comparison here is not great but works
			res.Relation != x.result.Relation ||
			res.CountryCode != x.result.CountryCode ||
			res.Repository != x.result.Repository ||
			res.AllocDate != x.result.AllocDate {
			t.Errorf("ASN result did not matched expected")
		}
		for i, a := range x.result.ASN {
			if res.ASN[i] != a {
				t.Errorf("ASN numbers did not match")
				break
			}
		}
	}
}

func TestIPQueryPeer(t *testing.T) {
	tests := []struct {
		ip     net.IP
		result PrefixInfo
	}{
		{
			net.IP{4, 2, 2, 2},
			PrefixInfo{
				Prefix: net.IPNet{
					IP:   net.IP{4, 0, 0, 0},
					Mask: net.CIDRMask(9, 32)},
				ASN:         []ASNumber{ASNumber(174), ASNumber(2914), ASNumber(3257), ASNumber(21385)},
				Relation:    Peer,
				CountryCode: "US",
				Repository:  "arin",
				AllocDate:   "1992-12-01",
			},
		},
	}

	for _, x := range tests {
		res, err := QueryIP(x.ip, Peer)
		if err != nil {
			t.Errorf("Received error %v", err)
		}
		if res.Prefix.String() != x.result.Prefix.String() || //String comparison here is not great but works
			res.Relation != x.result.Relation ||
			res.CountryCode != x.result.CountryCode ||
			res.Repository != x.result.Repository ||
			res.AllocDate != x.result.AllocDate {
			t.Errorf("Prefix result did not matched expected")
		}
		for i, a := range x.result.ASN {
			if res.ASN[i] != a {
				t.Errorf("ASN numbers did not match, got %v, expected %v", res.ASN, x.result.ASN)
				break
			}
		}
	}
}
