package ip2asn

import "testing"

func TestASNQuery(t *testing.T) {
	tests := []struct {
		asn    int
		result ASInfo
	}{
		{
			701,
			ASInfo{
				ASN:         []ASNumber{ASNumber(701)},
				CountryCode: "US",
				Repository:  "arin",
				AllocDate:   "1990-08-03",
				Description: "UUNET - MCI Communications Services, Inc. d/b/a Verizon Business, US",
			},
		},
	}

	for _, x := range tests {
		res, err := QueryASN(ASNumber(x.asn))
		if err != nil {
			t.Errorf("Received error %v", err)
		}
		if res.CountryCode != x.result.CountryCode ||
			res.Repository != x.result.Repository ||
			res.AllocDate != x.result.AllocDate ||
			res.Description != x.result.Description {
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
