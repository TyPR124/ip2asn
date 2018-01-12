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

//EnableCache Enables caching PrefixInfo and ASInfo objects
//Caching is thread safe, however calling EnableCache and DisableCache is NOT thread safe
func EnableCache() {
	if !cacheEnabled {
		cacheBits = new(cacheBitArray) //2MB bit array
		asnCache = map[ASNumber]ASInfo{}
		cacheRoot = iptree.NewDefaultRoot(net.IPv4len, nil)
		cacheEnabled = true
	}
}

//DisableCache Disables caching PrefixInfo and ASInfo objects
//Caching is thread safe, however calling EnableCache and DisableCache is NOT thread safe
func DisableCache() {
	cacheBits = nil
	asnCache = nil
	cacheRoot = nil
	cacheEnabled = false
}

//SaveCache is not implimented yet
//SaveCache will eventually allow saving cache to disk
func SaveCache() error {
	return ErrNotImplimented
}

//LoadCache is not implimented yet
//LoadCache will eventually allow loading cache from disk
func LoadCache() error {
	return ErrNotImplimented
}
