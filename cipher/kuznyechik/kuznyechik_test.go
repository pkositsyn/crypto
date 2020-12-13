package kuznyechik

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCipher(t *testing.T) {
	_, err := NewCipher(hexDecoderUnsafe("0000"))
	assert.EqualError(t, err, "kuznyechik: invalid key size 2")

	_, err = NewCipher(hexDecoderUnsafe("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"))
	assert.NoError(t, err)
}

// BenchmarkNewCipher-4 6687    173928 ns/op    29.44 MB/s    10400 B/op    641 allocs/op (initial)
// BenchmarkNewCipher-4 7180    161264 ns/op    31.75 MB/s    672 B/op      33 allocs/op  (do feistel with uints)
// BenchmarkNewCipher-4 6238    173390 ns/op    29.53 MB/s    832 B/op      33 allocs/op  (upgrade decrypt - also generate decryption keys)
// BenchmarkNewCipher-4 943480  1170 ns/op      4377.52 MB/s  320 B/op      1 allocs/op   (initialize constants with tables)
func BenchmarkNewCipher(b *testing.B) {
	// runs only once - can init here
	initTablesOnce.Do(initTables)

	randomDataLen := 1024 * 1024 * 24

	randomData := make([]byte, randomDataLen)
	if _, err := rand.Read(randomData); err != nil {
		panic(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(int64(numRounds * keyLength * blockSize))
	j := 0
	for i := 0; i < b.N; i++ {
		key := randomData[j : j+keyLength]
		_, _ = NewCipher(key)
		j = (j + keyLength) % randomDataLen
	}
}

func TestEncrypt(t *testing.T) {
	encryptor, err := NewCipher(hexDecoderUnsafe("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"))
	require.NoError(t, err)

	output := make([]byte, blockSize)
	encryptor.Encrypt(output, hexDecoderUnsafe("1122334455667700ffeeddccbbaa9988"))

	outputString := hex.EncodeToString(output)
	assert.Equal(t, "7f679d90bebc24305a468d42b9d4edcd", outputString, "expected and result differ")
}

// BenchmarkEncrypt-4 3284576    364 ns/op    43.94 MB/s    0 B/op    0 allocs/op (Implemented lsTable)
// BenchmarkEncrypt-4 3845746    321 ns/op    49.86 MB/s    0 B/op    0 allocs/op (add bounds check for tmp2)
// BenchmarkEncrypt-4 8575030    139 ns/op    115.02 MB/s   0 B/op    0 allocs/op (use uint64 instead of byte slices)
// BenchmarkEncrypt-4 7613686    143 ns/op    111.82 MB/s   0 B/op    0 allocs/op (add 2 more tables for decrypting - got slightly worse)
func BenchmarkEncrypt(b *testing.B) {
	encryptor, err := NewCipher(hexDecoderUnsafe("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"))
	require.NoError(b, err)

	randomDataLen := 1024 * 1024 * 24

	randomData := make([]byte, randomDataLen)
	if _, err := rand.Read(randomData); err != nil {
		panic(err)
	}
	output := make([]byte, blockSize)

	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(blockSize)
	j := 0
	for i := 0; i < b.N; i++ {
		blockToEncode := randomData[j : j+blockSize]
		encryptor.Encrypt(output, blockToEncode)
		j = (j + blockSize) % randomDataLen
	}
}

func TestDecrypt(t *testing.T) {
	encryptor, err := NewCipher(hexDecoderUnsafe("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"))
	require.NoError(t, err)

	output := make([]byte, blockSize)
	encryptor.Decrypt(output, hexDecoderUnsafe("7f679d90bebc24305a468d42b9d4edcd"))

	outputString := hex.EncodeToString(output)
	assert.Equal(t, "1122334455667700ffeeddccbbaa9988", outputString, "expected and result differ")
}

// BenchmarkDecrypt-4 5396226    225 ns/op    71.04 MB/s    0 B/op    0 allocs/op (initial with inverted l table + manually inverted s)
// BenchmarkDecrypt-4 7739772    154 ns/op    103.74 MB/s   0 B/op    0 allocs/op (invertedSL table + decrypt keys - major improvement)
func BenchmarkDecrypt(b *testing.B) {
	decryptor, err := NewCipher(hexDecoderUnsafe("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"))
	require.NoError(b, err)

	randomDataLen := 1024 * 1024 * 24

	randomData := make([]byte, randomDataLen)
	if _, err := rand.Read(randomData); err != nil {
		panic(err)
	}
	output := make([]byte, blockSize)

	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(blockSize)
	j := 0
	for i := 0; i < b.N; i++ {
		blockToDecode := randomData[j : j+blockSize]
		decryptor.Decrypt(output, blockToDecode)
		j = (j + blockSize) % randomDataLen
	}
}

func TestKeySetGet(t *testing.T) {
	encryptor, err := NewCipher(hexDecoderUnsafe("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"))
	require.NoError(t, err)

	err = encryptor.SetKey(nil)
	assert.EqualError(t, err, "kuznyechik: invalid key size 0")

	randomDataLen := 1024 * 1024

	randomData := make([]byte, randomDataLen)
	_, err = rand.Read(randomData)
	require.NoError(t, err)

	output := make([]byte, blockSize)
	for i := 0; i < randomDataLen; i += blockSize {
		encryptor.Encrypt(output, randomData[i:i+blockSize])
	}

	newKey := hexDecoderUnsafe("8899aabbccddeeff7766554433221100fedcba98765432100123456789abcdef")
	encryptorNewKey, err := NewCipher(newKey)
	require.NoError(t, err)

	err = encryptor.SetKey(newKey)
	require.NoError(t, err)

	assert.Equal(t, newKey, encryptor.GetKey())

	output2 := make([]byte, blockSize)

	for i := 0; i < randomDataLen; i += blockSize {
		encryptor.Encrypt(output, randomData[i:i+blockSize])
		encryptorNewKey.Encrypt(output2, randomData[i:i+blockSize])

		require.Equal(t, output, output2)
	}
}
