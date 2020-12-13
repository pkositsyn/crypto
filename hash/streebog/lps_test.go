package streebog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLPS(t *testing.T) {
	initTableOnce.Do(initLPSTable)

	testCases := []struct {
		toEncodeString string
		expectedString string
	}{
		{
			toEncodeString: "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expectedString: "b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574b383fc2eced4a574",
		},

		{
			toEncodeString: bit512{0x0101010101010101, 0x0101010101010101, 0x0101010101010101, 0x0101010101010101,
				0x0101010101010101, 0x0101010101010101, 0x0101010101010101, 0x0101010101010101}.String(),
			expectedString: "23c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f1523c5ee40b07b5f15",
		},

		{
			toEncodeString: "b2b1cd1ef7ec924286b7cf1cffe49c4c84b5c91afde694448abbcb18fbe0964682b3c516f9e2904080b1cd1ef7ec924286b7cf1cffe49c4c84b5c91afde69444",
			expectedString: "e60059d4d8e0758024c73f6f3183653f56579189602ae4c21e7953ebc0e212a0ce78a8df475c2fd4fc43fc4b71c01e35be465fb20dad2cf690cdf65028121bb9",
		},

		{
			toEncodeString: "028ba7f4d01e7f9d5848d3af0eb1d96b9ce98a6de0917562c2cd44a3bb516188f8ff1cbf5cb3cc7511c1d6266ab47661b6f5881802a0e8576e0399773c72e073",
			expectedString: "d0b00807642fd78f13f2c3ebc774e80de0e902d23aef2ee9a73d010807dae9c188be14f0b2da27973569cd2ba051301036f728bd1d7eec33f4d18af70c46cf1e",
		},

		{
			toEncodeString: "22f7df708943682316f1dd72814b662d14f3db7483496e251afdd976854f6c2712f5d778874d6a2110f7df708943682316f1dd72814b662d14f3db7483496e25",
			expectedString: "e549368917a0a2611d5e08c9c2fd5b3c563f18c0f68c410d84ae9d5fbdfb934055650121b7aa6d7b3e7d09d46ac4358adaa6ae44fa3b0402c4166d2c3eb2ef02",
		},

		{
			toEncodeString: "92cdb59aaeb185fcc80ec1c1701e230a0caf98039e3e8f03528b56cdc5fe9be968b90ed1221c36148187c448141b8c0026b39a767c0f1236fe458b1942dd1a12",
			expectedString: "18ee8f3176b2ebea3bd6cb8233694cea349769df88be26bf451cfab6a904a549da22de93a66a66b19c7e6b5eea633511e611d68c8401bfcd0c7d0cc39d4a5eb9",
		},

		{
			toEncodeString: "486906c521f45a8f43621cde3bf44599936b10ce2531558642a303de2038858593790ed02b3685585b750fc32cf44d925d6214de3c0585585b730ecb2cf440a5",
			expectedString: "909aa733e1f52321a2fe35bfb8f67e92fbc70ef544709d5739d8faaca4acf126e83e273745c25b7b8f4a83a7436f6353753cbbbe492262cd3a868eace0104af1",
		},

		{
			toEncodeString: "d82f14ab5f5ba0eed3240eb0455bbff8032d02a05b9eafe7d2e511b05e977fe4033f1cbe55997f39cb331dad525bb7f3cd2406b042aa7f39cb351ca5525bbac4",
			expectedString: "a3a72a2e0fb5e6f812681222fec037b0db972086a395a387a6084508cae13093aa71d352dcbce288e9a39718a727f6fd4c5da5d0bc10fac3707ccd127fe45475",
		},
	}

	for i, testCase := range testCases {
		toEncodeString, expectedString := testCase.toEncodeString, testCase.expectedString
		t.Run(fmt.Sprintf("lps %d", i), func(t *testing.T) {
			toEncode := bit512FromString(toEncodeString)
			expected := bit512FromString(expectedString)
			assert.Equal(t, expected.String(), lps(toEncode).String())
		})
	}
}
