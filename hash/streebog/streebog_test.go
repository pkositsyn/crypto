package streebog

import (
	"crypto/rand"
	"encoding/hex"
	"hash"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkSum512-4  14275  84178 ns/op  76.03 MB/s  64 B/op  1 allocs/op
func BenchmarkSum512(b *testing.B) {
	h := New512()
	const n = 100
	bytes := generateRandomBytes(blockSize * n)
	_, _ = h.Write(bytes)

	b.ReportAllocs()
	b.SetBytes(blockSize * n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = h.Sum(nil)
	}
}

// BenchmarkSum256-4  13179  83928 ns/op  76.26 MB/s  32 B/op  1 allocs/op
func BenchmarkSum256(b *testing.B) {
	h := New256()
	const n = 100
	bytes := generateRandomBytes(blockSize * n)
	_, _ = h.Write(bytes)

	b.ReportAllocs()
	b.SetBytes(blockSize * n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = h.Sum(nil)
	}
}

func TestNew(t *testing.T) {
	// Test no panic
	_ = New256()
	_ = New512()
}

func TestSize(t *testing.T) {
	h := New512()
	assert.Equal(t, 0, h.Size())

	_, _ = h.Write(generateRandomBytes(20))
	assert.Equal(t, 20, h.Size())

	_, _ = h.Write(generateRandomBytes(20))
	assert.Equal(t, 40, h.Size())

	h.Reset()
	assert.Equal(t, 0, h.Size())
}

func TestBlockSize(t *testing.T) {
	h := New512()
	assert.Equal(t, 64, h.BlockSize())
}

func TestWriteReturn(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		hashSize int
	}{
		{
			name:     "zero bytes hash 512",
			n:        0,
			hashSize: hashSize512,
		},

		{
			name:     "some bytes hash 512",
			n:        10,
			hashSize: hashSize512,
		},

		{
			name:     "many bytes hash 512",
			n:        1000,
			hashSize: hashSize512,
		},

		{
			name:     "zero bytes hash 256",
			n:        0,
			hashSize: hashSize256,
		},

		{
			name:     "some bytes hash 256",
			n:        10,
			hashSize: hashSize256,
		},

		{
			name:     "many bytes hash 256",
			n:        1000,
			hashSize: hashSize256,
		},
	}

	for _, testCase := range testCases {
		n, hashSize := testCase.n, testCase.hashSize
		t.Run(testCase.name, func(t *testing.T) {
			h := &Streebog{hashSize: hashSize}
			h.Reset()

			written, err := h.Write(generateRandomBytes(n))
			assert.NoError(t, err)
			assert.Equal(t, n, written)
		})
	}
}

func TestSum(t *testing.T) { // nolint:funlen
	testCases := []struct {
		name           string
		msgString      string
		expectedString string
		hashSize       int
	}{
		{
			name:           "msg shorter than 512, hash 512",
			msgString:      "323130393837363534333231303938373635343332313039383736353433323130393837363534333231303938373635343332313039383736353433323130",
			expectedString: "486f64c1917879417fef082b3381a4e211c324f074654c38823a7b76f830ad00fa1fbae42b1285c0352f227524bc9ab16254288dd6863dccd5b9f54a1ad0541b",
			hashSize:       hashSize512,
		},

		{
			name:           "msg longer than 512, hash 512",
			msgString:      "fbe2e5f0eee3c820fbeafaebef20fffbf0e1e0f0f520e0ed20e8ece0ebe5f0f2f120fff0eeec20f120faf2fee5e2202ce8f6f3ede220e8e6eee1e8f0f2d1202ce8f0f2e5e220e5d1",
			expectedString: "28fbc9bada033b1460642bdcddb90c3fb3e56c497ccd0f62b8a2ad4935e85f037613966de4ee00531ae60f3b5a47f8dae06915d5f2f194996fcabf2622e6881e",
			hashSize:       hashSize512,
		},

		{
			name:           "msg shorter than 512, hash 256",
			msgString:      "323130393837363534333231303938373635343332313039383736353433323130393837363534333231303938373635343332313039383736353433323130",
			expectedString: "00557be5e584fd52a449b16b0251d05d27f94ab76cbaa6da890b59d8ef1e159d",
			hashSize:       hashSize256,
		},

		{
			name:           "msg longer than 512, hash 256",
			msgString:      "fbe2e5f0eee3c820fbeafaebef20fffbf0e1e0f0f520e0ed20e8ece0ebe5f0f2f120fff0eeec20f120faf2fee5e2202ce8f6f3ede220e8e6eee1e8f0f2d1202ce8f0f2e5e220e5d1",
			expectedString: "508f7e553c06501d749a66fc28c6cac0b005746d97537fa85d9e40904efed29d",
			hashSize:       hashSize256,
		},
	}

	for _, testCase := range testCases {
		msgString, expectedString, hashSize := testCase.msgString, testCase.expectedString, testCase.hashSize
		t.Run(testCase.name, func(t *testing.T) {
			var h hash.Hash
			switch hashSize {
			case hashSize256:
				h = New256()
			case hashSize512:
				h = New512()
			default:
				t.Fatal("invalid hash size")
			}

			msg := hexDecoderUnsafe(msgString)
			expected := hexDecoderUnsafe(expectedString)

			n, err := h.Write(msg)
			assert.NoError(t, err)
			assert.Equal(t, len(msg), n)

			assert.Equal(t, hex.EncodeToString(expected), hex.EncodeToString(h.Sum(nil)))

			// Sum shouldn't change state
			assert.Equal(t, hex.EncodeToString(expected), hex.EncodeToString(h.Sum(nil)))
		})
	}
}

func TestSumMultipleWrite(t *testing.T) { // nolint:funlen
	testCases := []struct {
		name           string
		msgString      string
		expectedString string
		hashSize       int
	}{
		{
			name:           "msg shorter than 512, hash 512",
			msgString:      "323130393837363534333231303938373635343332313039383736353433323130393837363534333231303938373635343332313039383736353433323130",
			expectedString: "486f64c1917879417fef082b3381a4e211c324f074654c38823a7b76f830ad00fa1fbae42b1285c0352f227524bc9ab16254288dd6863dccd5b9f54a1ad0541b",
			hashSize:       hashSize512,
		},

		{
			name:           "msg longer than 512, hash 512",
			msgString:      "fbe2e5f0eee3c820fbeafaebef20fffbf0e1e0f0f520e0ed20e8ece0ebe5f0f2f120fff0eeec20f120faf2fee5e2202ce8f6f3ede220e8e6eee1e8f0f2d1202ce8f0f2e5e220e5d1",
			expectedString: "28fbc9bada033b1460642bdcddb90c3fb3e56c497ccd0f62b8a2ad4935e85f037613966de4ee00531ae60f3b5a47f8dae06915d5f2f194996fcabf2622e6881e",
			hashSize:       hashSize512,
		},

		{
			name:           "msg shorter than 512, hash 256",
			msgString:      "323130393837363534333231303938373635343332313039383736353433323130393837363534333231303938373635343332313039383736353433323130",
			expectedString: "00557be5e584fd52a449b16b0251d05d27f94ab76cbaa6da890b59d8ef1e159d",
			hashSize:       hashSize256,
		},

		{
			name:           "msg longer than 512, hash 256",
			msgString:      "fbe2e5f0eee3c820fbeafaebef20fffbf0e1e0f0f520e0ed20e8ece0ebe5f0f2f120fff0eeec20f120faf2fee5e2202ce8f6f3ede220e8e6eee1e8f0f2d1202ce8f0f2e5e220e5d1",
			expectedString: "508f7e553c06501d749a66fc28c6cac0b005746d97537fa85d9e40904efed29d",
			hashSize:       hashSize256,
		},
	}

	for _, testCase := range testCases {
		msgString, expectedString, hashSize := testCase.msgString, testCase.expectedString, testCase.hashSize
		t.Run(testCase.name, func(t *testing.T) {
			var h hash.Hash
			switch hashSize {
			case hashSize256:
				h = New256()
			case hashSize512:
				h = New512()
			default:
				t.Fatal("invalid hash size")
			}

			msg := hexDecoderUnsafe(msgString)
			expected := hexDecoderUnsafe(expectedString)

			for i := range msg {
				n, err := h.Write(msg[i : i+1])
				assert.NoError(t, err)
				assert.Equal(t, 1, n)
			}

			assert.Equal(t, hex.EncodeToString(expected), hex.EncodeToString(h.Sum(nil)))

			// Sum shouldn't be changed. Actually, Streebog cannot have partial state,
			// so implementation just recalculates hash for the whole data
			assert.Equal(t, hex.EncodeToString(expected), hex.EncodeToString(h.Sum(nil)))
		})
	}
}

func TestE2E(t *testing.T) {
	data1 := generateRandomBytes(40)
	data2 := generateRandomBytes(40)
	dataLong := generateRandomBytes(10_000)

	h := New512()
	sum0 := h.Sum(nil)

	_, _ = h.Write(data1)

	rollingHash := []byte{1, 2, 3}
	sum1 := h.Sum(rollingHash)
	sum2 := h.Sum(rollingHash)
	assert.Equal(t, sum1, sum2)
	assert.Len(t, sum1, 3+blockSize)
	assert.Equal(t, rollingHash, sum1[:3])

	h.Reset()
	sumReset := h.Sum(nil)
	assert.Equal(t, sum0, sumReset)

	h.Reset()
	h.Reset()
	sumReset = h.Sum(nil)
	assert.Equal(t, sum0, sumReset)

	_, _ = h.Write(data1)
	_, _ = h.Write(data2)
	sum1 = h.Sum(nil)

	h.Reset()
	_, _ = h.Write(append(data1, data2...))
	sum2 = h.Sum(nil)
	assert.Equal(t, sum1, sum2)

	_, _ = h.Write(nil)
	sum2 = h.Sum(nil)
	assert.Equal(t, sum1, sum2)

	h.Reset()
	_, _ = h.Write(dataLong)
	sum1 = h.Sum(sum1)

	h.Reset()
	for i := 0; i < 10; i++ {
		bytesLen := len(dataLong) / 10
		_, _ = h.Write(dataLong[i*bytesLen : (i+1)*bytesLen])
	}
	sum2 = h.Sum(sum2)
	assert.Equal(t, sum1, sum2)
}

func generateRandomBytes(n int) []byte {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return bytes
}
