package kuznyechik

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			"ffeeddccbbaa99881122334455667700",
			"b66cd8887d38e8d77765aeea0c9a7efc",
		},
		{
			"b66cd8887d38e8d77765aeea0c9a7efc",
			"559d8dd7bd06cbfe7e7b262523280d39",
		},
		{
			"559d8dd7bd06cbfe7e7b262523280d39",
			"0c3322fed531e4630d80ef5c5a81c50b",
		},
		{
			"0c3322fed531e4630d80ef5c5a81c50b",
			"23ae65633f842d29c5df529c13f5acda",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			input := hexDecoderUnsafe(testCase.input)
			output := make([]byte, blockSize)
			s(output, input)

			outputString := hex.EncodeToString(output)
			assert.Equal(t, testCase.expected, outputString, "expected and result differ")
		})

		t.Run(fmt.Sprintf("test_inverted %d", i), func(t *testing.T) {
			input := hexDecoderUnsafe(testCase.expected)
			output := make([]byte, blockSize)
			invertedS(output, input)

			i1, i2 := binary.BigEndian.Uint64(input[0:8]), binary.BigEndian.Uint64(input[8:16])
			o1, o2 := inverseS(i1, i2)
			output2 := bit128{o1, o2}.toByteSlice()

			outputString := hex.EncodeToString(output)
			assert.Equal(t, testCase.input, outputString, "expected and result differ")
			outputString = hex.EncodeToString(output2)
			assert.Equal(t, testCase.input, outputString, "expected and result differ")
		})
	}
}
