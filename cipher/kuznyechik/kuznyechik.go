package kuznyechik

import (
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"sync"
)

const (
	blockSize = 16
	keyLength = 32
	numRounds = 10
)

var initTablesOnce sync.Once

type KeySizeError int

func (k KeySizeError) Error() string {
	return fmt.Sprintf("kuznyechik: invalid key size %d", k)
}

// Cipher implements russian block cipher algorithm standard called "Kuznyechik".
type Cipher struct {
	key              []byte
	roundKeys        [numRounds]bit128
	decryptRoundKeys [numRounds]bit128
}

// Compile-time check that the cipher implements cipher.Block interface.
var _ cipher.Block = &Cipher{}

// NewCipher creates a new kuznyechik cipher instance.
func NewCipher(key []byte) (*Cipher, error) {
	k := len(key)
	if k != keyLength {
		return nil, KeySizeError(k)
	}

	initTablesOnce.Do(initTables)
	roundKeys := generateRoundKeys(key)
	return &Cipher{
		key:              key,
		roundKeys:        roundKeys,
		decryptRoundKeys: generateDecryptRoundKeys(roundKeys),
	}, nil
}

// BlockSize is cipher block size.
func (c *Cipher) BlockSize() int {
	return blockSize
}

// GetKey returns the current cipher key.
func (c *Cipher) GetKey() []byte {
	return c.key
}

// SetKey sets new key for cipher.
func (c *Cipher) SetKey(key []byte) error {
	if k := len(key); k != keyLength {
		return KeySizeError(k)
	}
	c.key = key
	c.roundKeys = generateRoundKeys(c.key)
	c.decryptRoundKeys = generateDecryptRoundKeys(c.roundKeys)
	return nil
}

// Encrypt encrypts blockSize bytes from src to dst. These buffers can overlap.
func (c *Cipher) Encrypt(dst, src []byte) {
	if len(src) < blockSize {
		panic("kuznyechik: input not full block")
	}
	if len(dst) < blockSize {
		panic("kuznyechik: output not full block")
	}
	c.doEncrypt(dst, src)
}

// Decrypt decrypts blockSize bytes from src to dst. These buffers can overlap.
func (c *Cipher) Decrypt(dst, src []byte) {
	if len(src) < blockSize {
		panic("kuznyechik: input not full block")
	}
	if len(dst) < blockSize {
		panic("kuznyechik: output not full block")
	}
	c.doDecrypt(dst, src)
}

// Encryption is XOR[k_9] l s XOR[k_8] ... l s XOR[k_0] (src), where XOR[x](y) is (x XOR y).
// We can use ls lookup table to fasten these steps.
func (c *Cipher) doEncrypt(dst, src []byte) {
	high := binary.BigEndian.Uint64(src[0:8])
	low := binary.BigEndian.Uint64(src[8:16])

	_ = c.roundKeys[numRounds-1]
	for index := 0; index < numRounds-1; index++ {
		high ^= c.roundKeys[index][0]
		low ^= c.roundKeys[index][1]
		high, low = ls(high, low)
	}
	high ^= c.roundKeys[numRounds-1][0]
	low ^= c.roundKeys[numRounds-1][1]

	binary.BigEndian.PutUint64(dst[0:8], high)
	binary.BigEndian.PutUint64(dst[8:16], low)
}

// Decryption is XOR[k_9] InvS InvL XOR[k_8] ... InvS InvL XOR[k_0] (src).
// As long as InvL(x XOR y) = InvL(x) XOR InvL(y) we can do following -
// Precalculate d_i = InvL(K_i), thus we get equal expression for decrypting (using precalculated InvL InvS):
// XOR[d_9] InvS XOR[d_8] InvL InvS XOR[d_7] ... InvL InvS XOR[d_0] InvL (src).
func (c *Cipher) doDecrypt(dst, src []byte) {
	high := binary.BigEndian.Uint64(src[0:8])
	low := binary.BigEndian.Uint64(src[8:16])

	high, low = inverseL(high, low)
	_ = c.decryptRoundKeys[numRounds-1]
	for index := numRounds - 1; index > 1; index-- {
		high ^= c.decryptRoundKeys[index][0]
		low ^= c.decryptRoundKeys[index][1]
		high, low = inverseSL(high, low)
	}
	high ^= c.decryptRoundKeys[1][0]
	low ^= c.decryptRoundKeys[1][1]
	high, low = inverseS(high, low)
	high ^= c.decryptRoundKeys[0][0]
	low ^= c.decryptRoundKeys[0][1]

	binary.BigEndian.PutUint64(dst[0:8], high)
	binary.BigEndian.PutUint64(dst[8:16], low)
}
