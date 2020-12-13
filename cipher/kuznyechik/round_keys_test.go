package kuznyechik

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRoundKeys(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			"8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef",
			[]string{
				"8899aabbccddeeff0011223344556677",
				"fedcba98765432100123456789abcdef",
				"db31485315694343228d6aef8cc78c44",
				"3d4553d8e9cfec6815ebadc40a9ffd04",
				"57646468c44a5e28d3e59246f429f1ac",
				"bd079435165c6432b532e82834da581b",
				"51e640757e8745de705727265a0098b1",
				"5a7925017b9fdd3ed72a91a22286f984",
				"bb44e25378c73123a5f32f73cdb6e517",
				"72e9dd7416bcf45b755dbaa88e4a4043",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			input := hexDecoderUnsafe(testCase.input)
			output := generateRoundKeys(input)

			var outputStrings []string
			for _, byteOutput := range output {
				outputStrings = append(outputStrings, hex.EncodeToString(byteOutput.toByteSlice()))
			}
			assert.Equal(t, testCase.expected, outputStrings, "expected and result differ")
		})
	}
}
