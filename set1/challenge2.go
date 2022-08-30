package set1

import "encoding/hex"

func Challenge2(in, xor string) string {
	out := XorEncode(FromHex(in), FromHex(xor))
	return ToHex(out)
}

func XorEncode(in []byte, xor []byte) []byte {
	n := len(in)
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
