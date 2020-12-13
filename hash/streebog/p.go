package streebog

// permutePosition is a reversed (as a slice, not inverted) P transform from GOST. Means permutation of bytes positions.
func permutePosition(i int) int {
	return permutation[i]
}

// permutation is a reversed permutation from GOST.
var permutation = [...]int{
	63, 55, 47, 39, 31, 23, 15, 7, 62, 54, 46, 38, 30, 22, 14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 60, 52, 44, 36, 28, 20, 12, 4, 59, 51, 43, 35, 27, 19, 11, 3, 58, 50,
	42, 34, 26, 18, 10, 2, 57, 49, 41, 33, 25, 17, 9, 1, 56, 48, 40, 32, 24, 16, 8, 0,
}
