package kuznyechik

func rStepFast(dst, src []byte, shift int) {
	dst[blockSize-1-shift] = lHelperFast(dst, src, shift)
}

func invertedRStepFast(dst, src []byte, shift int) {
	dst[shift] = invertedLHelperFast(dst, src, shift)
}

var (
	coefficients = [...]byte{148, 32, 133, 16, 194, 192, 1, 251, 1, 192, 194, 16, 133, 32, 148, 1}
)

func lHelperFast(vec1, vec2 []byte, shift int) (result byte) {
	// shift elements are processed here
	for index, value := range vec1[blockSize-shift:] {
		// last iteration with index equal shift-1
		result ^= mul(coefficients[index], value)
	}

	// blockSize-shift elements are processed here
	for index, value := range vec2[:blockSize-shift] {
		// first iteration with index equal shift
		result ^= mul(coefficients[index+shift], value)
	}

	// shift + (blockSize-shift) elements are processed total
	return
}

func invertedLHelperFast(vec1, vec2 []byte, shift int) (result byte) {
	// blockSize-shift-1 elements are processed here
	for index, value := range vec2[shift+1:] {
		// last iteration with index equal blockSize-shift-2
		result ^= mul(coefficients[index], value)
	}

	// shift elements are processed here
	for index, value := range vec1[:shift] {
		// last iteration with index equal blockSize-2
		result ^= mul(coefficients[blockSize-1-shift+index], value)
	}
	// 1 element is processed here
	result ^= mul(coefficients[blockSize-1], vec2[shift])

	// total (blockSize-shift-1) + shift + 1 = blockSize elements processed
	return
}

// mul is multiplication in GF(2)[x]/p(x) field, addition is simply xor
// See https://en.wikipedia.org/wiki/Finite_field_arithmetic#C_programming_example
func mul(a, b byte) (p byte) {
	for a > 0 && b > 0 {
		if b&1 == 1 {
			p ^= a
		}

		highBit := (a & 0x80) != 0
		a <<= 1
		if highBit {
			a ^= 0xC3 // p(x) = x^8 + x^7 + x^6 + x + 1
		}
		b >>= 1
	}
	return p
}
