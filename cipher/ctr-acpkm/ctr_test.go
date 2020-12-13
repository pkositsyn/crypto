package ctracpkm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCipher(t *testing.T) {
	key := make([]byte, 32)
	iv := make([]byte, 8)

	_, err := New(KuznyechikFactory, key, iv, 1, WithGammaSize(1))
	assert.NoError(t, err)

	invalidIV := make([]byte, 1)
	_, err = New(KuznyechikFactory, key, invalidIV, 1)
	assert.EqualError(t, err, ErrInvalidIVSize.Error())

	_, err = New(KuznyechikFactory, key, iv, 0)
	assert.EqualError(t, err, ErrInvalidKeyPressure.Error())

	_, err = New(KuznyechikFactory, key, iv, 1, WithGammaSize(3))
	assert.EqualError(t, err, ErrInvalidGammaSize.Error())

	_, err = New(KuznyechikFactory, key, iv, 1, WithGammaSize(1), WithGammaSize(2))
	assert.EqualError(t, err, ErrInvalidOptionsNumber.Error())
}

func TestCtrAcpkm(t *testing.T) {
	testCases := []struct {
		name              string
		messageString     string
		expectedString    string
		gamma             int
		keyPressureBlocks int
	}{
		{
			name:              "GOST test",
			messageString:     "1122334455667700ffeeddccbbaa998800112233445566778899aabbcceeff0a112233445566778899aabbcceeff0a002233445566778899aabbcceeff0a001133445566778899aabbcceeff0a001122445566778899aabbcceeff0a001122335566778899aabbcceeff0a0011223344",
			expectedString:    "f195d8bec10ed1dbd57b5fa240bda1b885eee733f6a13e5df33ce4b33c45dee44bceeb8f646f4c55001706275e85e800587c4df568d094393e4834afd0805046cf30f57686aeece11cfc6c316b8a896edffd07ec813636460c4f3b743423163e6409a9c282fac8d469d221e7fbd6de5d",
			gamma:             16,
			keyPressureBlocks: 2,
		},

		{
			name:              "gamma less than block size",
			messageString:     "1122334455667700ffeeddccbbaa998800112233445566778899aabbcceeff0a112233445566778899aabbcceeff0a002233445566778899aabbcceeff0a001133445566778899aabbcceeff0a001122445566778899aabbcceeff0a001122335566778899aabbcceeff0a0011223344",
			expectedString:    "f195d8bec10ed1db7a1118cc095ec1a2b4d9f9fcf7067f2c613b145f5895cd25bc9b395907e05722edf75634ec006f957749c272dcdc153780fad0fc8b0a07752c2ceb578f6d6af981c63805930fe4f57d8f20523a077aba1b4c1b8acce9625fa5c648769bed518aa7a064ebe9c5ba62",
			gamma:             8,
			keyPressureBlocks: 2,
		},

		{
			name:              "gamma is 1",
			messageString:     "1122334455667700ffeeddccbbaa998800112233445566778899aabbcceeff0a112233445566778899aabbcceeff0a002233445566778899aabbcceeff0a001133445566778899aabbcceeff0a001122445566778899aabbcceeff0a001122335566778899aabbcceeff0a0011223344",
			expectedString:    "f1a787ad3a88f9a0bc735293f98c12c3eb31621b9b2e6461c7ef73a2e6a6b1797ca6bdcfe3520b2e64461c1343781c3f2125b4e87ae4a5c1ba93d0f4c6af11385614487d55f8eb8a412c1f1493866801f6de6c693fe295d9581a2a1f2721b2e71b8f1ed6ed8c370428077877cd093640",
			gamma:             1,
			keyPressureBlocks: 2,
		},

		{
			name:              "key pressure is 1",
			messageString:     "1122334455667700ffeeddccbbaa998800112233445566778899aabbcceeff0a112233445566778899aabbcceeff0a002233445566778899aabbcceeff0a001133445566778899aabbcceeff0a001122445566778899aabbcceeff0a001122335566778899aabbcceeff0a0011223344",
			expectedString:    "f195d8bec10ed1dbd57b5fa240bda1b880f3dd9d7de317470284a3b762878336de32043f8e9909bcd8a1f764ddfee5b213871e346549707991fb7c0fb6417cc998d50b90865dde0dafd3e6dff90cadb3e319e31b5954643a54dd1d38e566f411cd69117d92afc9da4ecb63c1d5098d31",
			gamma:             16,
			keyPressureBlocks: 1,
		},
	}

	key, _ := hex.DecodeString("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef")
	iv, _ := hex.DecodeString("1234567890abcef0")

	for _, testCase := range testCases {
		message, _ := hex.DecodeString(testCase.messageString)
		expectedString := testCase.expectedString
		gamma, keyPressureBlocks := testCase.gamma, testCase.keyPressureBlocks
		t.Run(testCase.name, func(t *testing.T) {
			ctr, err := New(KuznyechikFactory, key, iv, keyPressureBlocks, WithGammaSize(gamma))
			require.NoError(t, err)

			dst := make([]byte, len(message))

			ctr.XORKeyStream(dst, message)
			assert.Equal(t, hex.EncodeToString(dst), expectedString)

			ctr, err = New(KuznyechikFactory, key, iv, keyPressureBlocks, WithGammaSize(gamma))
			require.NoError(t, err)

			dst2 := make([]byte, len(message))
			ctr.XORKeyStream(dst2, message)
			assert.Equal(t, dst, dst2)

			ctr, err = New(KuznyechikFactory, key, iv, keyPressureBlocks, WithGammaSize(gamma))
			require.NoError(t, err)

			dst2 = make([]byte, len(message))
			for i := range message {
				ctr.XORKeyStream(dst2[i:i+1], message[i:i+1])
			}
			assert.Equal(t, dst, dst2)
		})
	}
}

func TestAesBasic(t *testing.T) {
	testCases := []struct {
		name      string
		keyString string
		ivString  string
	}{
		{
			name:      "32 key size",
			keyString: "8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef",
			ivString:  "1234567890abcef00000000000000000",
		},

		{
			name:      "16 key size",
			keyString: "8899aabbccddeeff0011223344556677",
			ivString:  "1234567890abcef00000000000000000",
		},
	}

	for _, testCase := range testCases {
		key, _ := hex.DecodeString(testCase.keyString)
		iv, _ := hex.DecodeString(testCase.ivString)
		t.Run(testCase.name, func(t *testing.T) {
			const (
				keyPressure = 4
				blockSize   = 16
			)

			// Aes is 16, 24 or 32 byte key cipher. Block size is 16
			// 24 byte key is not supported because for acpkm block size must be divisible by key length
			ctrAcpkm, err := New(AesFactory, key, iv[:len(iv)/2], keyPressure)
			require.NoError(t, err)

			randomData := make([]byte, blockSize*keyPressure*2)
			_, err = rand.Read(randomData)
			require.NoError(t, err)

			a, err := aes.NewCipher(key)
			require.NoError(t, err)

			ctr := cipher.NewCTR(a, iv)

			output := make([]byte, blockSize)
			output2 := make([]byte, blockSize)

			for i := 0; i < len(randomData); i += blockSize {
				if i > 0 && i/blockSize%keyPressure == 0 {
					iv[len(iv)-1] = byte(i / blockSize)
					key = acpkm(a, key)
					a, err = aes.NewCipher(key)
					require.NoError(t, err)
					ctr = cipher.NewCTR(a, iv)
				}
				ctr.XORKeyStream(output, randomData[i:i+blockSize])
				ctrAcpkm.XORKeyStream(output2, randomData[i:i+blockSize])
				require.Equal(t, output, output2)
			}
		})
	}
}

// BenchmarkKuznyechikCtrAcpkm-4  376836  3137 ns/op  81.60 MB/s  0 B/op  0 allocs/op
func BenchmarkKuznyechikCtrAcpkm(b *testing.B) {
	key, _ := hex.DecodeString("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef")
	iv, _ := hex.DecodeString("1234567890abcef0")
	ctr, err := New(KuznyechikFactory, key, iv, 1024*256)
	require.NoError(b, err)

	randomDataLen := 1024 * 1024 * 24
	randomData := make([]byte, randomDataLen)
	_, err = rand.Read(randomData)
	require.NoError(b, err)

	blockSize := 16

	benchDataLen := 16 * blockSize
	output := make([]byte, benchDataLen)

	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(int64(benchDataLen))
	j := 0
	for i := 0; i < b.N; i++ {
		blockToEncode := randomData[j : j+benchDataLen]
		ctr.XORKeyStream(output, blockToEncode)
		j = (j + benchDataLen) % randomDataLen
	}
}

// BenchmarkAesCtrAcpkm-4  1344072  880 ns/op  290.83 MB/s  0 B/op  0 allocs/op
func BenchmarkAesCtrAcpkm(b *testing.B) {
	key, _ := hex.DecodeString("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef")
	iv, _ := hex.DecodeString("1234567890abcef0")

	ctr, err := New(AesFactory, key, iv, 1024*256)
	require.NoError(b, err)

	randomDataLen := 1024 * 1024 * 24
	randomData := make([]byte, randomDataLen)
	_, err = rand.Read(randomData)
	require.NoError(b, err)

	blockSize := 16

	benchDataLen := 16 * blockSize
	output := make([]byte, benchDataLen)

	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(int64(benchDataLen))
	j := 0
	for i := 0; i < b.N; i++ {
		blockToEncode := randomData[j : j+benchDataLen]
		ctr.XORKeyStream(output, blockToEncode)
		j = (j + benchDataLen) % randomDataLen
	}
}
