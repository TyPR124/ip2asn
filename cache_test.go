package ip2asn

import (
	"testing"
)

func TestIPCache(t *testing.T) {
	EnableCache()
	defer DisableCache()

	pfx, err := QueryIP([]byte{4, 2, 2, 2}, Origin)
	if err != nil {
		t.Errorf("Error on query: %v", err)
		return
	}

	//Make a copy of pfx
	//pfxcopy := *pfx

	//Check that there is something cached
	cached := cachedIP([]byte{4, 2, 2, 2})
	if cached == nil {
		t.Errorf("Did not receive cached PrefixInfo")
		return
	}

	if pfx.Prefix.String() != cached.Prefix.String() ||
		pfx.AllocDate != cached.AllocDate ||
		pfx.ASN[0] != cached.ASN[0] ||
		pfx.CountryCode != cached.CountryCode ||
		pfx.Relation != cached.Relation ||
		pfx.Repository != cached.Repository {
		t.Errorf("mismatched cache")
	}
}

func TestASCache(t *testing.T) {
	EnableCache()
	defer DisableCache()

	as, err := QueryASN(55)
	if err != nil {
		t.Errorf("Error on query asn: %v", err)
		return
	}

	//Check cache
	cached := cachedASN(55)
	if cached == nil {
		t.Errorf("Did not receive cached ASInfo")
		return
	}

	if as.AllocDate != cached.AllocDate ||
		as.ASN[0] != cached.ASN[0] ||
		as.CountryCode != cached.CountryCode ||
		as.Description != cached.Description ||
		as.Repository != cached.Repository {
		t.Errorf("mismatched ASN cache")
	}
}
