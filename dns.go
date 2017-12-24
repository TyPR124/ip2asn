package ip2asn

import (
	"net"
	"strconv"
	"strings"
)

//dns.go supports querying cymru.com via DNS

const (
	origin4Server = "origin.asn.cymru.com"
	origin6Server = "origin6.asn.cymru.com"
	peerServer    = "peer.asn.cymru.com"
	asnServer     = "asn.cymru.com"
)

var hexDigit = [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

func reverseIP4(ip net.IP) string {
	ret := ""
	for i := len(ip) - 1; i >= 0; i-- {
		ret += "." + strconv.Itoa(int(ip[i]))
	}
	return ret[1:] //Remove the first "."
}

func reverseIP6(ip net.IP) string {
	ret := ""
	for i := len(ip) - 1; i >= 0; i-- {
		//We have 8 bits, need to split into two 4 bit chunks
		hi := (ip[i] & 0xF0) >> 4
		lo := ip[i] & 0x0F

		ret += "." + hexDigit[lo] + "." + hexDigit[hi]
	}
	return ret[1:]
}

func ipQueryHost(ip net.IP, rel Relation) (string, error) {
	var reverse string
	if len(ip) == net.IPv4len {
		reverse = reverseIP4(ip)
	} else if len(ip) == net.IPv6len {
		reverse = reverseIP6(ip)
	} else {
		return "", ErrBadIPLength
	}
	if rel == Origin {
		if len(ip) == net.IPv4len {
			return reverse + "." + origin4Server, nil
		} else if len(ip) == net.IPv6len {
			return reverse + "." + origin6Server, nil
		}
	} else if rel == Peer {
		return reverse + "." + peerServer, nil
	} //else
	return "", ErrInvalidRelation
}

func asnQueryHost(asn ASNumber) string {
	return asn.String() + "." + asnServer
}

func prepDNSAnswer(in string) ([]ASNumber, []string, error) {
	//Turn this format
	// AS# AS# AS# AS# ... | sometext | moretext | blahtext | someotherstuff | etc...
	//into a list of ASNumber and strings
	//May  have any number of ASNumber (at least 1) and at least one text element
	//Error if format is bad

	var asn []ASNumber

	split := strings.Split(in, "|")
	n := len(split)
	if n < 2 {
		return nil, nil, ErrInvalidAnswerFormat
	}

	asntext := strings.Split(strings.Trim(split[0], " "), " ")
	if len(asntext) < 1 {
		return nil, nil, ErrInvalidAnswerFormat
	}
	asn = make([]ASNumber, len(asntext))
	for i, x := range asntext {
		asint, err := strconv.Atoi(x)
		if err != nil {
			return nil, nil, ErrInvalidAnswerFormat
		}
		asn[i] = ASNumber(asint)
	}

	for i := 1; i < n; i++ {
		split[i] = strings.Trim(split[i], " ")
	}

	return asn, split[1:], nil
}
