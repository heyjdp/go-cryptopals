package set1

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

func Challenge1(in string) string {
	out := ToBase64(FromHex(in))
	return out
}

func FromHex(in string) []byte {
	bytes := []byte(in)
	out := make([]byte, hex.DecodedLen(len(bytes)))
	n, err := hex.Decode(out, bytes)
	if err != nil {
		panic(err)
	}
	return out[0:n]
}

func ToBase64(inp []byte) string {
	buf := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	defer encoder.Close()
	_, err := encoder.Write(inp)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
