package ip2asn

import (
	"net"

	"github.com/TyPR124/iptree"
)

/* Important info about cache - please read

Cache is only allowed for IPv4 queries at a max granularity of /24
Cache supports both PrefixInfo and ASInfo objects.
Cache supports Origin queries, but NOT Peer queries (peers change too much)

Enabling cache will immediately allocate 2MB of memory.
 This space is used as a bit array to determine if a given /24 has been queried.
 (there are 2^24 (2MB) possible /24's in the IPv4 address space)

Saving and loading cache will require two files, an index file and a data file.

*/

func EnableCache() {
	//Initialize bit array
	if !cacheEnabled {
		cacheBits = new(cacheBitArray) //2MB bit array
		asnCache = map[ASNumber]ASInfo{}
		cacheRoot = iptree.NewDefaultRoot(net.IPv4len, nil)
		cacheEnabled = true
	}
}

func DisableCache() {
	cacheBits = nil
	asnCache = nil
	cacheRoot = nil
	cacheEnabled = false
}

func SaveCache() error {
	return ErrNotImplimented
}

func LoadCache() error {
	return ErrNotImplimented
}
