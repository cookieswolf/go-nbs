package cid

import (
	"fmt"
	"github.com/multiformats/go-multihash"
	"strings"
)

var ErrPossiblyInsecureHashFunction = fmt.Errorf("potentially insecure hash functions not allowed")
var ErrBelowMinimumHashLength = fmt.Errorf("hashes must be at %d least bytes long", minimumHashLength)

const minimumHashLength = 20

var goodset = map[uint64]bool{
	multihash.SHA2_256:     true,
	multihash.SHA2_512:     true,
	multihash.SHA3_224:     true,
	multihash.SHA3_256:     true,
	multihash.SHA3_384:     true,
	multihash.SHA3_512:     true,
	multihash.SHAKE_256:    true,
	multihash.DBL_SHA2_256: true,
	multihash.KECCAK_224:   true,
	multihash.KECCAK_256:   true,
	multihash.KECCAK_384:   true,
	multihash.KECCAK_512:   true,
	multihash.ID:           true,

	multihash.SHA1: true, // not really secure but still useful
}

func IsGoodHash(code uint64) bool {
	good, found := goodset[code]
	if good {
		return true
	}

	if !found {
		if code >= multihash.BLAKE2B_MIN+19 && code <= multihash.BLAKE2B_MAX {
			return true
		}
		if code >= multihash.BLAKE2S_MIN+19 && code <= multihash.BLAKE2S_MAX {
			return true
		}
	}

	return false
}

func ValidateCid(c *Cid) error {

	if !IsGoodHash(c.HashType) {
		return ErrPossiblyInsecureHashFunction
	}

	if c.HashType != multihash.ID && c.HashLength() < minimumHashLength {
		return ErrBelowMinimumHashLength
	}

	return nil
}

func IsValidPath(path string) (*Cid, []string, error) {
	if path[0] == '/' {
		path = path[1:]
	}

	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return nil, nil, ErrBelowMinimumHashLength
	}

	var firstUri = parts[0]
	if parts[0] == "nbs" {
		if len(parts) < 2 {
			return nil, nil, ErrBelowMinimumHashLength
		}
		firstUri = parts[1]
	}

	c, err := Decode(firstUri)
	if err != nil {
		return nil, nil, err
	}

	return c, parts, nil
}
