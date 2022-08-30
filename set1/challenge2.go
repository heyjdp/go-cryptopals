package set1

import (
	"encoding/hex"
)

func Challenge2(in, xor string) string {
	out := XorEncode(FromHex(in), FromHex(xor))
	return ToHex(out)
}

func XorEncode(in []byte, key []byte) []byte {
	n := len(in)
	xor := make([]byte, n)
	if len(key) == len(in) {
		copy(xor, key)
	} else {
		// Stretch the key
		j := 0
		for i := 0; i < len(in); i++ {
			j = i % len(key)
			xor[i] = key[j]
		}
	}
	out := make([]byte, n)
	for i := range in {
		out[i] = in[i] ^ xor[i]
	}
	return out[0:n]
}

func ToHex(inp []byte) string {
	out := make([]byte, hex.EncodedLen(len(inp)))
	n := hex.Encode(out, inp)
	return string(out[:n])
}
