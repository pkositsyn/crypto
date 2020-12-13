package kuznyechik

import "encoding/binary"

// bit128 represents effective 16 byte storage
type bit128 [2]uint64

// toByteSlice converts byte16 to byte slice
func (b bit128) toByteSlice() []byte {
	dst := make([]byte, 16)
	_ = dst[15]
	binary.BigEndian.PutUint64(dst[0:8], b[0])
	binary.BigEndian.PutUint64(dst[8:16], b[1])
	return dst
}

// byte16FromByteSlice converts byte slice to byte16
func byte16FromByteSlice(src []byte) (blk bit128) {
	_ = src[15]
	return bit128{binary.BigEndian.Uint64(src[0:8]), binary.BigEndian.Uint64(src[8:16])}
}
