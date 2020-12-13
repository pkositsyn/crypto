package streebog

import (
	"hash"
	"sync"
)

const (
	_ = iota * 256
	hashSize256
	hashSize512

	blockSize = 64
)

// initTableOnce is used to initialize lookup table only once.
// Every package entrypoint (New... functions) tries to initialize tables.
var initTableOnce sync.Once

// Streebog is a russian standard hash function described in GOST.
type Streebog struct {
	hashSize int
	h        bit512
	n        bit512
	sigma    bit512

	buf     []byte
	tmpBuf  [blockSize]byte
	bufBits bit512
}

// Compile-time check that Streebog implements hash.Hash.
var _ hash.Hash = (*Streebog)(nil)

// New256 creates 256bit output hash.
func New256() hash.Hash {
	initTableOnce.Do(initLPSTable)
	s := &Streebog{hashSize: hashSize256}
	s.Reset()
	return s
}

// New512 creates 512bit output hash.
func New512() hash.Hash {
	initTableOnce.Do(initLPSTable)
	s := &Streebog{hashSize: hashSize512}
	s.Reset()
	return s
}

// Reset resets the Hash to its initial state.
func (s *Streebog) Reset() {
	s.buf = []byte{}
	s.resetState()
}

// resetState resets only state but not buffer. Calling this function makes Sum idempotent.
func (s *Streebog) resetState() {
	s.n = bit512{}
	s.sigma = bit512{}
	if s.hashSize == hashSize256 {
		s.h = bit512{0x0101010101010101, 0x0101010101010101, 0x0101010101010101, 0x0101010101010101,
			0x0101010101010101, 0x0101010101010101, 0x0101010101010101, 0x0101010101010101}
	} else if s.hashSize == hashSize512 {
		s.h = bit512{}
	}
}

// Size returns the number of bytes Sum will return.
func (s *Streebog) Size() int {
	return len(s.buf)
}

// BlockSize returns the hash's underlying block size.
// There is no performance difference how much data to provide to Write.
func (s *Streebog) BlockSize() int {
	return blockSize
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state (i.e. sequential calls produce the same results).
func (s *Streebog) Sum(buf []byte) []byte {
	defer s.resetState()

	splitIndex := len(s.buf) - blockSize
	for ; splitIndex >= 0; splitIndex -= blockSize {
		s.encode(s.buf[splitIndex : splitIndex+blockSize])
	}
	splitIndex += blockSize // 0 <= splitIndex < blockSize
	copy(s.tmpBuf[blockSize-splitIndex:], s.buf)
	s.tmpBuf[blockSize-splitIndex-1] = 1
	for i := 0; i < blockSize-splitIndex-1; i++ {
		s.tmpBuf[i] = 0
	}

	s.bufBits = bit512FromBytes64(s.tmpBuf[:])

	s.h = s.g(s.h, s.bufBits)

	s.n = bit512{}
	s.h = s.g(s.h, add512bitUint64(s.n, bitSize(len(s.buf))))
	s.sigma.add512bitInPlace(s.bufBits)
	s.h = s.g(s.h, s.sigma)

	hashSum := s.h.Bytes64()
	return append(buf, hashSum[:byteSize(s.hashSize)]...)
}

// bitSize returns bit size from byte size.
func bitSize(byteSize int) uint64 {
	return uint64(8 * byteSize)
}

// byteSize returns byte size from bit size.
func byteSize(bitSize int) int {
	return bitSize / 8
}

// Write adds more data to the running hash.
// It never returns an error.
func (s *Streebog) Write(buf []byte) (int, error) {
	s.buf = append(s.buf, buf...)
	return len(buf), nil
}

// encode does encoding of a full block.
func (s *Streebog) encode(buf []byte) {
	s.bufBits = bit512FromBytes64(buf)

	s.h = s.g(s.h, s.bufBits)
	s.n.addUint64InPlace(512)
	s.sigma.add512bitInPlace(s.bufBits)
}
