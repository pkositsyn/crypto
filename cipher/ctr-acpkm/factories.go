package ctracpkm

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/pkositysn/crypto/cipher/kuznyechik"
)

var (
	_ BlockFactory = KuznyechikFactory
	_ BlockFactory = AesFactory
)

// KuznyechikFactory is a BlockFactory for keys of length 32.
func KuznyechikFactory(key []byte) (cipher.Block, error) {
	return kuznyechik.NewCipher(key)
}

// AesFactory is a BlockFactory for keys of length 16 or 32. Length 24 will not work.
func AesFactory(key []byte) (cipher.Block, error) {
	return aes.NewCipher(key)
}
