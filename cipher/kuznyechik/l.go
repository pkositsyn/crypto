package kuznyechik

// l is a linear transformation from LSX step
func l(dst, src []byte) {
	for i := range dst {
		rStepFast(dst, src, i)
	}
}

// invertedL is an inverse for linear transformation from LSX step
func invertedL(dst, src []byte) {
	for i := range dst {
		invertedRStepFast(dst, src, i)
	}
}
