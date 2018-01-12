package ip2asn

import (
	"net"
	"strconv"
)

//An ASNumber is an Autonomous System Number (ASN).
type ASNumber uint64

func (asn ASNumber) String() string {
	return "AS" + strconv.FormatUint(uint64(asn), 10)
}

//ASInfo contains information about a particular ASNumber
type ASInfo struct {
	ASN         []ASNumber
	CountryCode string
	Repository  string
	AllocDate   string
	Description string
}

func parseASInfo(record string) (*ASInfo, error) {
	// AS#s | CC | Repo | Alloc-Date | Descr
	asn, text, err := prepDNSAnswer(record)
	if err != nil {
		return nil, err
	}
	if len(text) != 4 {
		return nil, ErrInvalidASNRecord
	}
	cc := text[0]    //TODO: Validate country code
	repo := text[1]  //Maybe validate repo?
	alloc := text[2] //TODO: Validate Alloc Date
	descr := text[3] //Can descr be validated? probably not

	return &ASInfo{asn, cc, repo, alloc, descr}, nil
}

//QueryASN takes an ASNumber and returns an *ASInfo
func QueryASN(asn ASNumber) (*ASInfo, error) {
	if cacheEnabled {
		if info := cachedASN(asn); info != nil {
			return info, nil
		}
	}
	q := asnQueryHost(asn)
	a, err := net.LookupTXT(q)
	if err != nil {
		return nil, err
	}
	if len(a) != 1 {
		return nil, ErrTooManyAnswers
	}
	ret, err := parseASInfo(a[0])
	if err != nil {
		return nil, err
	}

	if cacheEnabled {
		addCachedASN(*ret)
	}

	return ret, nil
}
