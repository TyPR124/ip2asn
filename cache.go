package ip2asn

import (
	"net"

	"github.com/TyPR124/iptree"
)

const (
	numberOfSlash24s = uint32(256 * 256 * 256)
)

type cacheBitArray [numberOfSlash24s / 64]uint64

func setBitArray(x uint32, v bool) {
	if x > numberOfSlash24s {
		panic("Argument out of range")
	}
	if v { //turn on
		cacheBits[x/64] |= 0x1 << (x % 64)
	} else { //turn off
		cacheBits[x/64] &= ^(0x1 << (x % 64))
	}
}

func getBitArray(x uint32) bool {
	if x > numberOfSlash24s {
		panic("Argument out of range")
	}
	return cacheBits[x/64]&(0x1<<(x%64)) > 0
}

var (
	cacheEnabled bool
	cacheRoot    iptree.Root
	cacheBits    *cacheBitArray
	asnCache     map[ASNumber]ASInfo
)

func init() {
	cacheEnabled = false
}

func slash24(ip net.IP) uint32 {
	return uint32(ip[0])<<16 |
		uint32(ip[1])<<8 |
		uint32(ip[2])
}

func cachedIP(ip net.IP) *PrefixInfo {
	//Check if this IP's /24 was cached?
	s24 := slash24(ip)
	if !getBitArray(s24) {
		return nil
	}

	v, err := cacheRoot.Find(net.IPNet{
		IP:   ip,
		Mask: net.IPMask{255, 255, 255, 255},
	}, true)

	if err != nil || v == nil {
		return nil
	}
	pfx := v.(PrefixInfo)
	return &pfx
}

func addCachedPrefix(ip net.IP, pfx PrefixInfo) {
	s24 := slash24(ip)
	if err := cacheRoot.Insert(pfx.Prefix, pfx); err == nil {
		setBitArray(s24, true)
	}
}

func cachedASN(asn ASNumber) *ASInfo {
	info, ok := asnCache[asn]
	if !ok {
		return nil
	}
	return &info
}
