package streebog

// lps is a fast L(P(S(X))) transform. It must be strongly optimized - does around 70% of computational work.
func lps(value bit512) bit512 {
	return bit512{
		lpsTable[0][uint8(value[0]>>56)] ^ lpsTable[1][uint8(value[1]>>56)] ^ lpsTable[2][uint8(value[2]>>56)] ^ lpsTable[3][uint8(value[3]>>56)] ^ lpsTable[4][uint8(value[4]>>56)] ^ lpsTable[5][uint8(value[5]>>56)] ^ lpsTable[6][uint8(value[6]>>56)] ^ lpsTable[7][uint8(value[7]>>56)],
		lpsTable[0][uint8(value[0]>>48)] ^ lpsTable[1][uint8(value[1]>>48)] ^ lpsTable[2][uint8(value[2]>>48)] ^ lpsTable[3][uint8(value[3]>>48)] ^ lpsTable[4][uint8(value[4]>>48)] ^ lpsTable[5][uint8(value[5]>>48)] ^ lpsTable[6][uint8(value[6]>>48)] ^ lpsTable[7][uint8(value[7]>>48)],
		lpsTable[0][uint8(value[0]>>40)] ^ lpsTable[1][uint8(value[1]>>40)] ^ lpsTable[2][uint8(value[2]>>40)] ^ lpsTable[3][uint8(value[3]>>40)] ^ lpsTable[4][uint8(value[4]>>40)] ^ lpsTable[5][uint8(value[5]>>40)] ^ lpsTable[6][uint8(value[6]>>40)] ^ lpsTable[7][uint8(value[7]>>40)],
		lpsTable[0][uint8(value[0]>>32)] ^ lpsTable[1][uint8(value[1]>>32)] ^ lpsTable[2][uint8(value[2]>>32)] ^ lpsTable[3][uint8(value[3]>>32)] ^ lpsTable[4][uint8(value[4]>>32)] ^ lpsTable[5][uint8(value[5]>>32)] ^ lpsTable[6][uint8(value[6]>>32)] ^ lpsTable[7][uint8(value[7]>>32)],
		lpsTable[0][uint8(value[0]>>24)] ^ lpsTable[1][uint8(value[1]>>24)] ^ lpsTable[2][uint8(value[2]>>24)] ^ lpsTable[3][uint8(value[3]>>24)] ^ lpsTable[4][uint8(value[4]>>24)] ^ lpsTable[5][uint8(value[5]>>24)] ^ lpsTable[6][uint8(value[6]>>24)] ^ lpsTable[7][uint8(value[7]>>24)],
		lpsTable[0][uint8(value[0]>>16)] ^ lpsTable[1][uint8(value[1]>>16)] ^ lpsTable[2][uint8(value[2]>>16)] ^ lpsTable[3][uint8(value[3]>>16)] ^ lpsTable[4][uint8(value[4]>>16)] ^ lpsTable[5][uint8(value[5]>>16)] ^ lpsTable[6][uint8(value[6]>>16)] ^ lpsTable[7][uint8(value[7]>>16)],
		lpsTable[0][uint8(value[0]>>8)] ^ lpsTable[1][uint8(value[1]>>8)] ^ lpsTable[2][uint8(value[2]>>8)] ^ lpsTable[3][uint8(value[3]>>8)] ^ lpsTable[4][uint8(value[4]>>8)] ^ lpsTable[5][uint8(value[5]>>8)] ^ lpsTable[6][uint8(value[6]>>8)] ^ lpsTable[7][uint8(value[7]>>8)],
		lpsTable[0][uint8(value[0])] ^ lpsTable[1][uint8(value[1])] ^ lpsTable[2][uint8(value[2])] ^ lpsTable[3][uint8(value[3])] ^ lpsTable[4][uint8(value[4])] ^ lpsTable[5][uint8(value[5])] ^ lpsTable[6][uint8(value[6])] ^ lpsTable[7][uint8(value[7])],
	}
}
