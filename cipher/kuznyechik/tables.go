package kuznyechik

var (
	invertedLTable  [blockSize][256]bit128
	invertedSLTable [blockSize][256]bit128
	lsTable         [blockSize][256]bit128
)

// initTables creates lookup tables
// We are able to make them due to the fact l(x XOR y) = l(x) XOR l(y)
func initTables() {
	tmpLTable := makeLTable()

	fillInvertedLTable()
	fillLSTable(tmpLTable)
	fillInvertedSLTable()

	makeRoundKeysConstants(tmpLTable)
}

func makeLTable() (lTable [blockSize][256]bit128) {
	for i := 0; i < blockSize; i++ {
		lTable[i] = generateTableLine(i, l)
	}
	return
}

func fillInvertedLTable() {
	for i := 0; i < blockSize; i++ {
		invertedLTable[i] = generateTableLine(i, invertedL)
	}
}

func generateTableLine(i int, f func(dst, src []byte)) (result [256]bit128) {
	vec := make([]byte, 16)
	tmp := make([]byte, 16)
	bitValues := make([]bit128, 8)
	for j := 0; j < 8; j++ {
		tmp[i] = byte(1 << j)
		f(vec, tmp)
		bitValues[j] = byte16FromByteSlice(vec)
	}
	tmp[i] = 0

	for j := 0; j < 256; j++ {
		jCopy := j
		var resultByte16 bit128
		for i := range bitValues {
			if jCopy%2 == 1 {
				resultByte16[0] ^= bitValues[i][0]
				resultByte16[1] ^= bitValues[i][1]
			}
			jCopy /= 2
		}
		result[j] = resultByte16
	}
	return
}

func fillLSTable(lTable [blockSize][256]bit128) {
	for i := 0; i < blockSize; i++ {
		for j := 0; j < 256; j++ {
			lsTable[i][j] = lTable[i][nonlinear(byte(j))]
		}
	}
}

func fillInvertedSLTable() {
	for i := 0; i < blockSize; i++ {
		for j := 0; j < 256; j++ {
			invertedSLTable[i][j] = invertedLTable[i][inverseNonlinear(byte(j))]
		}
	}
}

func makeRoundKeysConstants(lTable [blockSize][256]bit128) {
	for i := range constants {
		constants[i] = lTable[blockSize-1][byte(i+1)]
	}
}
