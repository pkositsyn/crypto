package streebog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS(t *testing.T) { // nolint:funlen
	testCases := []struct {
		toEncodeString string
		expectedString string
	}{
		{
			toEncodeString: "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expectedString: "fcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfc",
		},

		{
			toEncodeString: bit512{0x0101010101010101, 0x0101010101010101, 0x0101010101010101, 0x0101010101010101,
				0x0101010101010101, 0x0101010101010101, 0x0101010101010101, 0x0101010101010101}.String(),
			expectedString: "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		},

		{
			toEncodeString: "d82f14ab5f5ba0eed3240eb0455bbff8032d02a05b9eafe7d2e511b05e977fe4033f1cbe55997f39cb331dad525bb7f3cd2406b042aa7f39cb351ca5525bbac4",
			expectedString: "8d4f93828747a76c49e204adc8473bd11101dda7470a415b832b77ad5dbc572d111f14950ce8570be4aecd9f0e472fd2d9e231ad2c38570be46a14000e47a586",
		},

		{
			toEncodeString: "92cdb59aaeb185fcc80ec1c1701e230a0caf98039e3e8f03528b56cdc5fe9be968b90ed1221c36148187c448141b8c0026b39a767c0f1236fe458b1942dd1a12",
			expectedString: "ecd95e282645a83930045858325f5afa2341dc110ad303110ef676d9ac63509bf3a3041b65148f93f5c986f293bb7cfcef92288ac34df08f63c8f6362cd8f1f0",
		},

		{
			toEncodeString: "028ba7f4d01e7f9d5848d3af0eb1d96b9ce98a6de0917562c2cd44a3bb516188f8ff1cbf5cb3cc7511c1d6266ab47661b6f5881802a0e8576e0399773c72e073",
			expectedString: "ddf644e6e15f5733bff249410445536f4e9bd69e200f3596b3d9ea737d70a1d7d1b6143b9c9288357758f8ef78278aa155f4d717dda7cb12b211e87e7f19203d",
		},

		{
			toEncodeString: "486906c521f45a8f43621cde3bf44599936b10ce2531558642a303de2038858593790ed02b3685585b750fc32cf44d925d6214de3c0585585b730ecb2cf440a5",
			expectedString: "f29131ac18e613035196148598e6c8e8de6fe9e75c840c432c731185f906a8a8de5404e1428fa8bf47354d408be63aecb79693857f6ea8bf473d04e48be6eb00",
		},

		{
			toEncodeString: "92cdb59aaeb185fcc80ec1c1701e230a0caf98039e3e8f03528b56cdc5fe9be968b90ed1221c36148187c448141b8c0026b39a767c0f1236fe458b1942dd1a12",
			expectedString: "ecd95e282645a83930045858325f5afa2341dc110ad303110ef676d9ac63509bf3a3041b65148f93f5c986f293bb7cfcef92288ac34df08f63c8f6362cd8f1f0",
		},

		{
			toEncodeString: "22f7df708943682316f1dd72814b662d14f3db7483496e251afdd976854f6c2712f5d778874d6a2110f7df708943682316f1dd72814b662d14f3db7483496e25",
			expectedString: "65c061327951f35a99a6d819f5a29a0193d290ffa92ab25cf14b538aa8cc9d21f0f4fe6dc93a7818e9c061327951f35a99a6d819f5a29a0193d290ffa92ab25c",
		},

		{
			toEncodeString: "028ba7f4d01e7f9d5848d3af0eb1d96b9ce98a6de0917562c2cd44a3bb516188f8ff1cbf5cb3cc7511c1d6266ab47661b6f5881802a0e8576e0399773c72e073",
			expectedString: "ddf644e6e15f5733bff249410445536f4e9bd69e200f3596b3d9ea737d70a1d7d1b6143b9c9288357758f8ef78278aa155f4d717dda7cb12b211e87e7f19203d",
		},

		{
			toEncodeString: "b2b1cd1ef7ec924286b7cf1cffe49c4c84b5c91afde694448abbcb18fbe0964682b3c516f9e2904080b1cd1ef7ec924286b7cf1cffe49c4c84b5c91afde69444",
			expectedString: "4645d95fc0beec2c432f8914b62d4efd3e5e37f14b097aead67de417c220b0482492ac996667e0ebdf45d95fc0beec2c432f8914b62d4efd3e5e37f14b097aea",
		},
	}

	for i, testCase := range testCases {
		toEncodeString, expectedString := testCase.toEncodeString, testCase.expectedString
		t.Run(fmt.Sprintf("s %d", i), func(t *testing.T) {
			toEncode := hexDecoderUnsafe(toEncodeString)
			expected := hexDecoderUnsafe(expectedString)
			received := make([]byte, len(toEncode))
			for i := range received {
				received[i] = nonlinear(int(toEncode[i]))
			}
			assert.Equal(t, expected, received)
		})
	}
}
