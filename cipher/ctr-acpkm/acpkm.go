package ctracpkm

import "crypto/cipher"

var acpkmConstant = [...]byte{
	0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87,
	0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
	0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97,
	0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f,
}

// acpkm is a procedure of making a new key for previous.
// Key length must be divisible by block length and not greater than 32.
func acpkm(cipher cipher.Block, lastKey []byte) []byte {
	newKey := make([]byte, len(lastKey))
	blockSize := cipher.BlockSize()
	for i := 0; i < len(newKey); i += blockSize {
		cipher.Encrypt(newKey[i:i+blockSize], acpkmConstant[i:i+blockSize])
	}
	return newKey
}
