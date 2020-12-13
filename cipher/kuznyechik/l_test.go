package kuznyechik

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestL(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			"64a59400000000000000000000000000",
			"d456584dd0e3e84cc3166e4b7fa2890d",
		},
		{
			"d456584dd0e3e84cc3166e4b7fa2890d",
			"79d26221b87b584cd42fbc4ffea5de9a",
		},
		{
			"79d26221b87b584cd42fbc4ffea5de9a",
			"0e93691a0cfc60408b7b68f66b513c13",
		},
		{
			"0e93691a0cfc60408b7b68f66b513c13",
			"e6a8094fee0aa204fd97bcb0b44b8580",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			input := hexDecoderUnsafe(testCase.input)
			output := make([]byte, blockSize)
			l(output, input)

			outputString := hex.EncodeToString(output)
			assert.Equal(t, testCase.expected, outputString, "expected and result differ")
		})

		t.Run(fmt.Sprintf("test_inverted %d", i), func(t *testing.T) {
			input := hexDecoderUnsafe(testCase.expected)
			output := make([]byte, blockSize)
			invertedL(output, input)

			i1, i2 := binary.BigEndian.Uint64(input[0:8]), binary.BigEndian.Uint64(input[8:16])
			o1, o2 := inverseL(i1, i2)
			output2 := bit128{o1, o2}.toByteSlice()

			outputString := hex.EncodeToString(output)
			assert.Equal(t, testCase.input, outputString, "expected and result differ")
			outputString = hex.EncodeToString(output2)
			assert.Equal(t, testCase.input, outputString, "must be equal")
		})
	}
}
