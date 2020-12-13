package hmac

import "hash"

type hmac struct {
	opad, ipad []byte
	inner, outer hash.Hash
}

type HashFactory func() hash.Hash

func New(hashFactory HashFactory, key []byte) hash.Hash {
	h := &hmac{
		inner: hashFactory(),
		outer: hashFactory(),
	}
	blockSize := h.inner.BlockSize()
	h.ipad = make([]byte, blockSize)
	h.opad = make([]byte, blockSize)
	if len(key) > blockSize {
		h.outer.Sum(key)
		key = h.outer.Sum(nil)
	}

	copy(h.ipad, key)
	copy(h.opad, key)

	for i := range h.ipad {
		h.ipad[i] ^= 0x36
	}
	for i := range h.opad {
		h.ipad[i] ^= 0x5c
	}
	h.inner.Write(h.ipad)
	return h
}

func (h hmac) Write(p []byte) (n int, err error) {
	return h.inner.Write(p)
}

func (h hmac) Sum(b []byte) []byte {
	initialLen := len(b)
	b = h.inner.Sum(b)
	h.outer.Reset()
	h.outer.Write(h.opad)
	h.outer.Write(b)
	h.outer.Write(b[initialLen:])
	return h.outer.Sum(b[:initialLen])
}

func (h hmac) Reset() {
	h.inner.Reset()
	h.inner.Write(h.ipad)
}

func (h hmac) Size() int {
	return h.outer.Size()
}

func (h hmac) BlockSize() int {
	return h.inner.BlockSize()
}
