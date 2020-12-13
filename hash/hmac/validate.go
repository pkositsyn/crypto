package hmac

import (
	"crypto/subtle"

	"github.com/pkositysn/crypto/hash/streebog"
)

func ValidMAC(message, messageTag, key []byte) bool {
	h := New(streebog.New512, key)
	h.Write(message)
	expectedTag := h.Sum(nil)
	return Equal(expectedTag, messageTag)
}

func Equal(tag1, tag2 []byte) bool {
	return subtle.ConstantTimeCompare(tag1, tag2) == 1
}
