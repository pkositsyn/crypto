package kuznyechik

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {
	testCases := []struct {
		a        byte
		b        byte
		expected byte
	}{
		{
			0x94,
			0x94,
			0xA4,
		},
		{
			0x01,
			0x01,
			0x01,
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			output := mul(testCase.a, testCase.b)
			assert.Equal(t, testCase.expected, output, "expected and result differ")
		})
	}
}
