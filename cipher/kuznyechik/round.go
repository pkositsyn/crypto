package kuznyechik

func ls(high, low uint64) (highNew, lowNew uint64) {
	d0, d1, d2, d3, d4, d5, d6, d7 := uint8(high>>56), uint8(high>>48), uint8(high>>40), uint8(high>>32), uint8(high>>24), uint8(high>>16), uint8(high>>8), uint8(high)
	e0, e1, e2, e3, e4, e5, e6, e7 := uint8(low>>56), uint8(low>>48), uint8(low>>40), uint8(low>>32), uint8(low>>24), uint8(low>>16), uint8(low>>8), uint8(low)
	_ = lsTable[15]
	highNew = lsTable[0][d0][0] ^ lsTable[1][d1][0] ^ lsTable[2][d2][0] ^ lsTable[3][d3][0] ^ lsTable[4][d4][0] ^ lsTable[5][d5][0] ^ lsTable[6][d6][0] ^ lsTable[7][d7][0]
	highNew ^= lsTable[8][e0][0] ^ lsTable[9][e1][0] ^ lsTable[10][e2][0] ^ lsTable[11][e3][0] ^ lsTable[12][e4][0] ^ lsTable[13][e5][0] ^ lsTable[14][e6][0] ^ lsTable[15][e7][0]
	lowNew = lsTable[0][d0][1] ^ lsTable[1][d1][1] ^ lsTable[2][d2][1] ^ lsTable[3][d3][1] ^ lsTable[4][d4][1] ^ lsTable[5][d5][1] ^ lsTable[6][d6][1] ^ lsTable[7][d7][1]
	lowNew ^= lsTable[8][e0][1] ^ lsTable[9][e1][1] ^ lsTable[10][e2][1] ^ lsTable[11][e3][1] ^ lsTable[12][e4][1] ^ lsTable[13][e5][1] ^ lsTable[14][e6][1] ^ lsTable[15][e7][1]
	return
}

func inverseL(high, low uint64) (highNew, lowNew uint64) {
	d0, d1, d2, d3, d4, d5, d6, d7 := uint8(high>>56), uint8(high>>48), uint8(high>>40), uint8(high>>32), uint8(high>>24), uint8(high>>16), uint8(high>>8), uint8(high)
	e0, e1, e2, e3, e4, e5, e6, e7 := uint8(low>>56), uint8(low>>48), uint8(low>>40), uint8(low>>32), uint8(low>>24), uint8(low>>16), uint8(low>>8), uint8(low)
	_ = invertedLTable[15]
	highNew = invertedLTable[0][d0][0] ^ invertedLTable[1][d1][0] ^ invertedLTable[2][d2][0] ^ invertedLTable[3][d3][0] ^ invertedLTable[4][d4][0] ^ invertedLTable[5][d5][0] ^ invertedLTable[6][d6][0] ^ invertedLTable[7][d7][0]
	highNew ^= invertedLTable[8][e0][0] ^ invertedLTable[9][e1][0] ^ invertedLTable[10][e2][0] ^ invertedLTable[11][e3][0] ^ invertedLTable[12][e4][0] ^ invertedLTable[13][e5][0] ^ invertedLTable[14][e6][0] ^ invertedLTable[15][e7][0]
	lowNew = invertedLTable[0][d0][1] ^ invertedLTable[1][d1][1] ^ invertedLTable[2][d2][1] ^ invertedLTable[3][d3][1] ^ invertedLTable[4][d4][1] ^ invertedLTable[5][d5][1] ^ invertedLTable[6][d6][1] ^ invertedLTable[7][d7][1]
	lowNew ^= invertedLTable[8][e0][1] ^ invertedLTable[9][e1][1] ^ invertedLTable[10][e2][1] ^ invertedLTable[11][e3][1] ^ invertedLTable[12][e4][1] ^ invertedLTable[13][e5][1] ^ invertedLTable[14][e6][1] ^ invertedLTable[15][e7][1]
	return
}

func inverseSL(high, low uint64) (highNew, lowNew uint64) {
	d0, d1, d2, d3, d4, d5, d6, d7 := uint8(high>>56), uint8(high>>48), uint8(high>>40), uint8(high>>32), uint8(high>>24), uint8(high>>16), uint8(high>>8), uint8(high)
	e0, e1, e2, e3, e4, e5, e6, e7 := uint8(low>>56), uint8(low>>48), uint8(low>>40), uint8(low>>32), uint8(low>>24), uint8(low>>16), uint8(low>>8), uint8(low)
	_ = invertedSLTable[15]
	highNew = invertedSLTable[0][d0][0] ^ invertedSLTable[1][d1][0] ^ invertedSLTable[2][d2][0] ^ invertedSLTable[3][d3][0] ^ invertedSLTable[4][d4][0] ^ invertedSLTable[5][d5][0] ^ invertedSLTable[6][d6][0] ^ invertedSLTable[7][d7][0]
	highNew ^= invertedSLTable[8][e0][0] ^ invertedSLTable[9][e1][0] ^ invertedSLTable[10][e2][0] ^ invertedSLTable[11][e3][0] ^ invertedSLTable[12][e4][0] ^ invertedSLTable[13][e5][0] ^ invertedSLTable[14][e6][0] ^ invertedSLTable[15][e7][0]
	lowNew = invertedSLTable[0][d0][1] ^ invertedSLTable[1][d1][1] ^ invertedSLTable[2][d2][1] ^ invertedSLTable[3][d3][1] ^ invertedSLTable[4][d4][1] ^ invertedSLTable[5][d5][1] ^ invertedSLTable[6][d6][1] ^ invertedSLTable[7][d7][1]
	lowNew ^= invertedSLTable[8][e0][1] ^ invertedSLTable[9][e1][1] ^ invertedSLTable[10][e2][1] ^ invertedSLTable[11][e3][1] ^ invertedSLTable[12][e4][1] ^ invertedSLTable[13][e5][1] ^ invertedSLTable[14][e6][1] ^ invertedSLTable[15][e7][1]
	return
}
