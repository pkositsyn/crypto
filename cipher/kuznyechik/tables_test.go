package kuznyechik

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkInitTables-4 24    48460701 ns/op    2097585 B/op    65536 allocs/op  (init lsTable and invertedLTable)
// BenchmarkInitTables-4 15    73584033 ns/op    4195017 B/op    131072 allocs/op (added invertedSL table so achieved approx 45% boost on decrypt)
// BenchmarkInitTables-4 15    69971653 ns/op    702 B/op        0 allocs/op      (implemented fast invertedL)
// BenchmarkInitTables-4 15    71492760 ns/op    1605 B/op       32 allocs/op     (added initialization for roundKeysConstants)
// BenchmarkInitTables-4 22    46273980 ns/op    372 B/op        0 allocs/op      (reuse tables for creating each other)
// BenchmarkInitTables-4 675    1721889 ns/op    1024 B/op       64 allocs/op     (calculate less l and invertedL due to xor)
func BenchmarkInitTables(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initTables()
	}
}

func TestMakeRoundKeysConstants(t *testing.T) {
	testCases := []struct {
		expected []string
	}{
		{
			[]string{
				"6ea276726c487ab85d27bd10dd849401",
				"dc87ece4d890f4b3ba4eb92079cbeb02",
				"b2259a96b4d88e0be7690430a44f7f03",
				"7bcd1b0b73e32ba5b79cb140f2551504",
				"156f6d791fab511deabb0c502fd18105",
				"a74af7efab73df160dd208608b9efe06",
				"c9e8819dc73ba5ae50f5b570561a6a07",
				"f6593616e6055689adfba18027aa2a08",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			var outputStrings []string
			for _, byteOutput := range constants[:8] {
				outputStrings = append(outputStrings, hex.EncodeToString(byteOutput.toByteSlice()))
			}
			assert.Equal(t, testCase.expected, outputStrings, "expected and result differ")
		})
	}
}
