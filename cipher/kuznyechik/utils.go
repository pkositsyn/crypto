package kuznyechik

import (
	"encoding/hex"
	"log"
)

func hexDecoderUnsafe(value string) []byte {
	result, err := hex.DecodeString(value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
