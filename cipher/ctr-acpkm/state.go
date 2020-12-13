package ctracpkm

type state struct {
	keyCodedGammas int
	keysUsed       int

	bufferUsedBytes int
	streamBuffer    []byte
}
