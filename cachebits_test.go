package ip2asn

import "testing"

func TestCacheBits(t *testing.T) {
	numToTest := uint32(10000)
	cacheBits = new(cacheBitArray)
	for i := uint32(0); i < 10000; i++ {
		if getBitArray(i) {
			t.Error("Got an on bit after initializing")
		}
	}

	for i := uint32(0); i < numToTest; i += 3 {
		setBitArray(i, true)
	}

	for i := uint32(0); i < numToTest; i++ {
		if on := getBitArray(i); i%3 == 0 && !on {
			t.Error("Got an off bit after setting it on")
		} else if i%3 > 0 && on {
			t.Error("Got an on bit when it was never set")
		}
	}

	for i := uint32(0x00ffffff); i > uint32(0x00ffffff-numToTest); i-- {
		if getBitArray(i) {
			t.Error("Got an on bit after initializing")
		}

		setBitArray(i, true)

		if !getBitArray(i) {
			t.Error("Got on off bit after setting it on")
		}

		setBitArray(i, false)
	}

	//Should get out of range
	defer func() {
		err := recover()
		if err == nil {
			t.Error("Expected panic, didn't get one")
		}
	}()
	setBitArray(uint32(0x01000000), true)
}
