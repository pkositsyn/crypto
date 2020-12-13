package streebog

import "testing"

// BenchmarkInitLPSTable-4  102705  11036 ns/op  185.57 MB/s  0 B/op  0 allocs/op
func BenchmarkInitLPSTable(b *testing.B) {
	b.ReportAllocs()
	// table size
	b.SetBytes(blockSize / 8 * 256)
	for i := 0; i < b.N; i++ {
		initLPSTable()
	}
}
