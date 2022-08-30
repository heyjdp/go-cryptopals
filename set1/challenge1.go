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

func ToBase64(in []byte) string {
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)
	defer encoder.Close()
	_, err := encoder.Write(in)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
