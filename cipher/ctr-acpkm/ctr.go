package ctracpkm

import (
	"crypto/cipher"
)

const (
	maxCachedKeys = 10
)

// CTR is a cipher block mode with a limited key pressure.
type CTR struct {
	cipherFactory BlockFactory
	cipher        cipher.Block
	ctr           []byte
	firstCtr      []byte

	keyPressureBlocksNum int
	gammaSize            int

	keysCache  [maxCachedKeys][]byte
	keysCached int
	lastKey    []byte
	state      state
}

var _ cipher.Stream = (*CTR)(nil)

// BlockFactory is a factory for building cipher.Block
type BlockFactory func(key []byte) (cipher.Block, error)

// New creates a new instance of CTR.
// This is necessary to pass a BlockFactory. Error in creating cipher is checked only once.
// BlockFactory is called each time keys are rotated, because there is no way to reset cipher.Block state.
// keyPressureBlocksNum is the number of blocks processed before the keys are switched.
// options are additional options for cipher. Currently, there is only one option - gamma. For details see WithGamma.
func New(factory BlockFactory, firstKey, iv []byte, keyPressureBlocksNum int, options ...CTROptions) (cipher.Stream, error) {
	if len(options) > 1 {
		return nil, ErrInvalidOptionsNumber
	}

	if keyPressureBlocksNum < 1 {
		return nil, ErrInvalidKeyPressure
	}
	block, err := factory(firstKey)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize/2 {
		return nil, ErrInvalidIVSize
	}

	ctr := make([]byte, blockSize)
	copy(ctr, iv)

	firstCtr := make([]byte, blockSize)
	copy(firstCtr, iv)

	ctrBlockMode := &CTR{
		cipherFactory:        factory,
		cipher:               block,
		ctr:                  ctr,
		firstCtr:             firstCtr,
		keyPressureBlocksNum: keyPressureBlocksNum,
		gammaSize:            blockSize,
		lastKey:              firstKey,
		state:                state{streamBuffer: make([]byte, 0, 2*blockSize)}, // leave extra space for gamma
	}
	ctrBlockMode.keysCache[0] = firstKey
	ctrBlockMode.keysCached = 1

	if len(options) == 1 {
		gammaSize := int(options[0])
		if gammaSize == 0 || blockSize%gammaSize != 0 {
			return nil, ErrInvalidGammaSize
		}

		ctrBlockMode.gammaSize = gammaSize
	}
	return ctrBlockMode, nil
}

// XORKeyStream implements cipher.Stream interface. dst and src may overlap.
func (c *CTR) XORKeyStream(dst, src []byte) {
	if len(dst) < len(src) {
		panic("ctr-acpkm: output smaller than input")
	}

	for len(src) > 0 {
		if c.state.bufferUsedBytes >= len(c.state.streamBuffer)-c.gammaSize {
			c.refillStreamBuffer()
		}

		n := len(src)
		if n > c.gammaSize {
			n = c.gammaSize
		}
		if sbLen := len(c.state.streamBuffer); n > sbLen {
			n = sbLen
		}
		for i := 0; i < n; i++ {
			dst[i] = src[i] ^ c.state.streamBuffer[c.state.bufferUsedBytes]
			c.state.bufferUsedBytes++
		}
		src = src[n:]
		dst = dst[n:]
	}
}

func (c *CTR) refillStreamBuffer() {
	blockSize := c.cipher.BlockSize()
	leftBytes := len(c.state.streamBuffer) - c.state.bufferUsedBytes
	copy(c.state.streamBuffer, c.state.streamBuffer[c.state.bufferUsedBytes:])
	c.state.streamBuffer = c.state.streamBuffer[:cap(c.state.streamBuffer)]

	for leftBytes <= blockSize-c.gammaSize {
		c.cipher.Encrypt(c.state.streamBuffer[leftBytes:], c.ctr)
		c.state.keyCodedGammas++
		if c.state.keyCodedGammas*c.gammaSize%(c.keyPressureBlocksNum*blockSize) == 0 {
			c.startNewKey()
		}
		leftBytes += c.gammaSize

		for i := len(c.ctr) - 1; i >= 0; i-- {
			c.ctr[i]++
			if c.ctr[i] != 0 {
				break
			}
		}
	}
	c.state.streamBuffer = c.state.streamBuffer[:leftBytes]
	c.state.bufferUsedBytes = 0
}

func (c *CTR) setCachedKey() {
	c.lastKey = c.keysCache[c.state.keysUsed]
	c.cipher, _ = c.cipherFactory(c.lastKey)
}

func (c *CTR) generateNewKey() {
	c.lastKey = acpkm(c.cipher, c.lastKey)
	c.cipher, _ = c.cipherFactory(c.lastKey)

	if c.keysCached < maxCachedKeys {
		c.keysCache[c.keysCached] = c.lastKey
		c.keysCached++
	}
}

func (c *CTR) startNewKey() {
	c.state.keysUsed++
	if c.state.keysUsed < c.keysCached {
		c.setCachedKey()
	} else {
		c.generateNewKey()
	}
}
