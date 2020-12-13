package streebog

// lpsTable is a lookup table used for fast L(P(S(x)) computing (this is not really LPS, but very close to it).
// The idea is that lps(x X y) = lps(x) X lps(y), where X is a cartesian product (i.e. x & y = 0).
// Thus for x, y such that x & y = 0 it is true that lps(x XOR y) = lps(x) XOR lps(y).
// The table consists of lps values for sparse blocks with only 1 non-zero byte.
// Indexes meaning:
//     1st dim - [blockSize / 8] - index of a non-zero byte (turns out we need positions only divisible by 8)
//     2nd dim - [256]           - value of the byte
// Result value is also a sparse block with at most 8 non-zero bytes.
// One can notice that position of the result depends only on the first dimension. Thus it should be computed separately.
var lpsTable [blockSize / 8][256]uint64

// initLPSTable initializes lookup table
func initLPSTable() {
	for i := 0; i < blockSize/8; i++ {
		p := permutePosition(8 * i)
		// For linear transform we need only 8 byte block. Other bytes make no sense.
		mod := p % 8
		// Only one byte from 8 is not zero. All its bit positions should be shifted.
		bitLinearPos := mod * 8

		// 8 values for each non-zero bit in byte.
		var bitValues bit512
		for j := 0; j < 8; j++ {
			// linear takes a non-zero bit position counting from left to right.
			bitValues[j] = linear(blockSize - 1 - (bitLinearPos + j))
		}

		for j := 0; j < 256; j++ {
			var resultValue uint64
			// Firstly S in applied
			nonLinear := nonlinear(j)
			for k := 0; k < 8; k++ {
				if nonLinear%2 == 1 {
					// Here LP is applied for every non-zero bit.
					// Do xor here as well to same computations in rounds. That is why it is not really LPS.
					resultValue ^= bitValues[k]
				}
				nonLinear /= 2
			}
			lpsTable[i][j] = resultValue
		}
	}
}
