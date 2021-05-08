package bp

import "math/big"

const (
	// HashLength is the expected length of the hash
	HashLength = 32
)
// Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [32]byte

//BytesToHash sets b to hash.
//If b is larger than len(h), b will be cropped from the left.
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func HashToBytes(h Hash) []byte {
	var b []byte
	copy(b, h[:])
	return b
}

// SetBytes sets the hash to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

// Big converts a hash to a big integer.
func (h Hash) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }
