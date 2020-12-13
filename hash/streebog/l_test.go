package streebog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestL(t *testing.T) { // nolint:funlen
	testCases := []struct {
		toEncodeString string
		expectedString string
	}{
		{
			toEncodeString: "fcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfc",
			expectedString: "b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574",
		},

		{
			toEncodeString: "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			expectedString: "23c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f15",
		},

		{
			toEncodeString: "46433ed624df433e452f5e7d92452f5ed98937e4acd989375f14f117995f14f1c0b64bc266c0b64bbe2d092067be2d09ec4e7ab0e0ec4e7a2cfdea48eb2cfdea",
			expectedString: "e60059d4d8e0758024c73f6f3183653f56579189602ae4c21e7953ebc0e212a0ce78a8df475c2fd4fc43fc4b71c01e35be465fb20dad2cf690cdf65028121bb9",
		},

		{
			toEncodeString: "ddbf4eb3d17755b2f6f29bd9b658f4114449d6ea14f8d7e8e6419e733bef177ee104207d9c78dd7f5f450f709227a719575335a1888acb20336f96d735a1123d",
			expectedString: "d0b00807642fd78f13f2c3ebc774e80de0e902d23aef2ee9a73d010807dae9c188be14f0b2da27973569cd2ba051301036f728bd1d7eec33f4d18af70c46cf1e",
		},

		{
			toEncodeString: "659993f1f0e99993c0a6d24bf4c0a6d261d89053fe61d8903219ff8a6d3219ff79f5a9a8c979f5a951a22acc3a51a22af39ab29d78f39ab25a015c21185a015c",
			expectedString: "e549368917a0a2611d5e08c9c2fd5b3c563f18c0f68c410d84ae9d5fbdfb934055650121b7aa6d7b3e7d09d46ac4358adaa6ae44fa3b0402c4166d2c3eb2ef02",
		},

		{
			toEncodeString: "ec30230ef3f5ef63d90441f6a3c992c85e58dc76048628f6285811d91bf28a3626320aac6593c32c455fd36314bb4dd8a85a03508f7cf0f139fa119b93fc8ff0",
			expectedString: "18ee8f3176b2ebea3bd6cb8233694cea349769df88be26bf451cfab6a904a549da22de93a66a66b19c7e6b5eea633511e611d68c8401bfcd0c7d0cc39d4a5eb9",
		},

		{
			toEncodeString: "f251de2cde47b74791966f735435963d3114e911044d9304ac85e785e14085e418985cf9428b7f8be6e684068fe66ee613c80ca8a83aa8eb03e843a8bfecbf00",
			expectedString: "909aa733e1f52321a2fe35bfb8f67e92fbc70ef544709d5739d8faaca4acf126e83e273745c25b7b8f4a83a7436f6353753cbbbe492262cd3a868eace0104af1",
		},

		{
			toEncodeString: "8d49118311e4d9e44fe2012b1faee26a9304dd7714cd311482ada7ad959fad0087c8475d0c0e2c0e47470abce8473847a73b4157572f57a56cd15b2d0bd20b86",
			expectedString: "a3a72a2e0fb5e6f812681222fec037b0db972086a395a387a6084508cae13093aa71d352dcbce288e9a39718a727f6fd4c5da5d0bc10fac3707ccd127fe45475",
		},
	}

	for i, testCase := range testCases {
		toEncodeString, expectedString := testCase.toEncodeString, testCase.expectedString
		t.Run(fmt.Sprintf("l %d", i), func(t *testing.T) {
			toEncode := bit512FromString(toEncodeString)
			expected := bit512FromString(expectedString)
			var received bit512
			for i := range toEncode {
				for j := 0; j < 64; j++ {
					if toEncode[i]&(1<<j) != 0 {
						received[i] ^= linear(blockSize - 1 - j)
					}
				}
			}
			assert.Equal(t, expected.String(), received.String())
		})
	}
}
