package ip2asn

import (
	"net"
)

//Relation describes the relationship between the ASNumber and the prefix
type Relation byte

const (
	//Origin is a Relation, indicating the ASNumber is the origin of the prefix
	Origin Relation = iota
	//Peer is a Relation, indicating the ASNumber is a peer of the prefix
	Peer
)

//PrefixInfo contains information about a prefix
type PrefixInfo struct {
	Prefix      net.IPNet
	ASN         []ASNumber
	Relation    Relation
	CountryCode string
	Repository  string
	AllocDate   string
}

//QueryIP takes an IPAddr and returns a *PrefixInfo
func QueryIP(ip net.IP, rel Relation) (*PrefixInfo, error) {
	q, err := ipQueryHost(ip, rel)
	if err != nil {
		return nil, err
	}
	a, err := net.LookupTXT(q)
	if err != nil {
		return nil, err
	}
	if len(a) != 1 {
		return nil, ErrTooManyAnswers
	}

	ret, err := parsePrefixInfo(a[0], rel)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func parsePrefixInfo(record string, rel Relation) (*PrefixInfo, error) {
	// AS# AS# AS# | Prefix | CC | Repo | Alloc-Date
	asn, text, err := prepDNSAnswer(record)
	if err != nil {
		return nil, err
	}
	if len(text) != 4 {
		return nil, ErrInvalidPrefixRecord
	}

	_, pfx, err := net.ParseCIDR(text[0])
	if err != nil {
		return nil, ErrInvalidPrefixString
	}
	cc := text[1]    //TODO: Validate country code
	repo := text[2]  //Maybe validate repo?
	alloc := text[3] //TODO: Validate Alloc Date

	return &PrefixInfo{*pfx, asn, rel, cc, repo, alloc}, nil
}

//QueryBulkIP takes a slice of IP and returns a slice of *PrefixInfo
func QueryBulkIP([]net.IP) ([]*PrefixInfo, error) {
	//stuff
	return nil, ErrNotImplimented
}
