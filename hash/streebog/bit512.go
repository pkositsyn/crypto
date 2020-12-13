package streebog

import (
	"encoding/binary"
	"encoding/hex"
)

// bit512 is an effective representation of 512 bits.
// Go leaks 128+ bit native structures so this is the most efficient way for doing xor512bit.
type bit512 [8]uint64

// bit512FromBytes64 should be called only on bytes slice with len 64.
func bit512FromBytes64(bytes []byte) bit512 {
	var b bit512
	b[0] = binary.BigEndian.Uint64(bytes[0:8])
	b[1] = binary.BigEndian.Uint64(bytes[8:16])
	b[2] = binary.BigEndian.Uint64(bytes[16:24])
	b[3] = binary.BigEndian.Uint64(bytes[24:32])
	b[4] = binary.BigEndian.Uint64(bytes[32:40])
	b[5] = binary.BigEndian.Uint64(bytes[40:48])
	b[6] = binary.BigEndian.Uint64(bytes[48:56])
	b[7] = binary.BigEndian.Uint64(bytes[56:64])
	return b
}

// bit512FromString can be called only on hex strings representing at least 512 bits.
// Otherwise function panics.
func bit512FromString(s string) bit512 {
	var b bit512
	bytes := hexDecoderUnsafe(s)
	b[0] = binary.BigEndian.Uint64(bytes[0:8])
	b[1] = binary.BigEndian.Uint64(bytes[8:16])
	b[2] = binary.BigEndian.Uint64(bytes[16:24])
	b[3] = binary.BigEndian.Uint64(bytes[24:32])
	b[4] = binary.BigEndian.Uint64(bytes[32:40])
	b[5] = binary.BigEndian.Uint64(bytes[40:48])
	b[6] = binary.BigEndian.Uint64(bytes[48:56])
	b[7] = binary.BigEndian.Uint64(bytes[56:64])
	return b
}

// String gives hex string representation of 512 bits.
func (b bit512) String() string {
	var result [64]byte
	for i := 0; i < 8; i++ {
		binary.BigEndian.PutUint64(result[8*i:8*i+8], b[i])
	}
	return hex.EncodeToString(result[:])
}

// Bytes64 converts internal 512 bits to byte slice.
func (b bit512) Bytes64() [64]byte {
	var result [64]byte
	binary.BigEndian.PutUint64(result[0:8], b[0])
	binary.BigEndian.PutUint64(result[8:16], b[1])
	binary.BigEndian.PutUint64(result[16:24], b[2])
	binary.BigEndian.PutUint64(result[24:32], b[3])
	binary.BigEndian.PutUint64(result[32:40], b[4])
	binary.BigEndian.PutUint64(result[40:48], b[5])
	binary.BigEndian.PutUint64(result[48:56], b[6])
	binary.BigEndian.PutUint64(result[56:64], b[7])
	return result
}

// add512bit adds 2 bit512 fields modulo 2^512.
func add512bit(this bit512, that bit512) bit512 {
	var result bit512
	var overflow uint64
	for i := 7; i >= 0; i-- {
		result[i] = this[i] + that[i] + overflow
		// Overflow check
		if result[i] < this[i] {
			overflow = 1
		} else {
			overflow = 0
		}
	}
	return result
}

// add512bitInPlace adds 2 bit512 fields inplace modulo 2^512.
func (b *bit512) add512bitInPlace(that bit512) {
	var overflow uint64
	for i := 7; i >= 0; i-- {
		b[i] += that[i] + overflow
		// Overflow check
		if b[i] < that[i] {
			overflow = 1
		} else {
			overflow = 0
		}
	}
}

// add512bitUint64 adds uint64 to bit512 modulo 2^512.
func add512bitUint64(this bit512, that uint64) bit512 {
	this[7] += that
	for i := 7; i >= 1; i-- {
		// Overflow check
		if this[i] < that {
			that = 1
			this[i-1] += that
		} else {
			that = 0
		}
	}
	return this
}

// addUint64InPlace adds uint64 to bit512 inplace modulo 2^512.
func (b *bit512) addUint64InPlace(that uint64) {
	b[7] += that
	for i := 7; i >= 1; i-- {
		// Overflow check
		if b[i] < that {
			that = 1
			b[i-1] += that
		} else {
			that = 0
		}
	}
}

// xor512bit does xor512bit operation on 2 bit512 fields. Implemented in the way it is inlined.
func xor512bit(this bit512, that bit512) bit512 {
	this[0] ^= that[0]
	this[1] ^= that[1]
	this[2] ^= that[2]
	this[3] ^= that[3]
	this[4] ^= that[4]
	this[5] ^= that[5]
	this[6] ^= that[6]
	this[7] ^= that[7]
	return this
}

// xor512bitInPlace does xor512bit inplace operation on 2 bit512 fields. Implemented in the way it is inlined.
func (b *bit512) xor512bitInPlace(that bit512) {
	b[0] ^= that[0]
	b[1] ^= that[1]
	b[2] ^= that[2]
	b[3] ^= that[3]
	b[4] ^= that[4]
	b[5] ^= that[5]
	b[6] ^= that[6]
	b[7] ^= that[7]
}
