package streebog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) { // nolint:funlen
	testCases := []struct {
		toEncodeString string
		expectedString string
	}{
		{
			toEncodeString: "fcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfc",
			expectedString: "fcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfcfc",
		},

		{
			toEncodeString: "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			expectedString: "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		},

		{
			toEncodeString: "4645d95fc0beec2c432f8914b62d4efd3e5e37f14b097aead67de417c220b0482492ac996667e0ebdf45d95fc0beec2c432f8914b62d4efd3e5e37f14b097aea",
			expectedString: "46433ed624df433e452f5e7d92452f5ed98937e4acd989375f14f117995f14f1c0b64bc266c0b64bbe2d092067be2d09ec4e7ab0e0ec4e7a2cfdea48eb2cfdea",
		},

		{
			toEncodeString: "ddf644e6e15f5733bff249410445536f4e9bd69e200f3596b3d9ea737d70a1d7d1b6143b9c9288357758f8ef78278aa155f4d717dda7cb12b211e87e7f19203d",
			expectedString: "ddbf4eb3d17755b2f6f29bd9b658f4114449d6ea14f8d7e8e6419e733bef177ee104207d9c78dd7f5f450f709227a719575335a1888acb20336f96d735a1123d",
		},

		{
			toEncodeString: "65c061327951f35a99a6d819f5a29a0193d290ffa92ab25cf14b538aa8cc9d21f0f4fe6dc93a7818e9c061327951f35a99a6d819f5a29a0193d290ffa92ab25c",
			expectedString: "659993f1f0e99993c0a6d24bf4c0a6d261d89053fe61d8903219ff8a6d3219ff79f5a9a8c979f5a951a22acc3a51a22af39ab29d78f39ab25a015c21185a015c",
		},

		{
			toEncodeString: "ecd95e282645a83930045858325f5afa2341dc110ad303110ef676d9ac63509bf3a3041b65148f93f5c986f293bb7cfcef92288ac34df08f63c8f6362cd8f1f0",
			expectedString: "ec30230ef3f5ef63d90441f6a3c992c85e58dc76048628f6285811d91bf28a3626320aac6593c32c455fd36314bb4dd8a85a03508f7cf0f139fa119b93fc8ff0",
		},

		{
			toEncodeString: "f29131ac18e613035196148598e6c8e8de6fe9e75c840c432c731185f906a8a8de5404e1428fa8bf47354d408be63aecb79693857f6ea8bf473d04e48be6eb00",
			expectedString: "f251de2cde47b74791966f735435963d3114e911044d9304ac85e785e14085e418985cf9428b7f8be6e684068fe66ee613c80ca8a83aa8eb03e843a8bfecbf00",
		},

		{
			toEncodeString: "ddf644e6e15f5733bff249410445536f4e9bd69e200f3596b3d9ea737d70a1d7d1b6143b9c9288357758f8ef78278aa155f4d717dda7cb12b211e87e7f19203d",
			expectedString: "ddbf4eb3d17755b2f6f29bd9b658f4114449d6ea14f8d7e8e6419e733bef177ee104207d9c78dd7f5f450f709227a719575335a1888acb20336f96d735a1123d",
		},

		{
			toEncodeString: "ecd95e282645a83930045858325f5afa2341dc110ad303110ef676d9ac63509bf3a3041b65148f93f5c986f293bb7cfcef92288ac34df08f63c8f6362cd8f1f0",
			expectedString: "ec30230ef3f5ef63d90441f6a3c992c85e58dc76048628f6285811d91bf28a3626320aac6593c32c455fd36314bb4dd8a85a03508f7cf0f139fa119b93fc8ff0",
		},
	}

	for i, testCase := range testCases {
		toEncodeString, expectedString := testCase.toEncodeString, testCase.expectedString
		t.Run(fmt.Sprintf("p %d", i), func(t *testing.T) {
			toEncode := hexDecoderUnsafe(toEncodeString)
			expected := hexDecoderUnsafe(expectedString)
			received := make([]byte, len(toEncode))
			for i := range toEncode {
				received[permutePosition(i)] = toEncode[blockSize-1-i]
			}
			assert.Equal(t, expected, received)
		})
	}
}
