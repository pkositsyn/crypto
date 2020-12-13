package streebog

import (
	"encoding/hex"
)

// hexDecoderUnsafe is a helper function for tests.
func hexDecoderUnsafe(value string) []byte {
	result, _ := hex.DecodeString(value)
	return result
}
