package streebog

// g is a g_N(h, M) function from GOST.
func (s *Streebog) g(h bit512, msg bit512) bit512 {
	key := xor512bit(h, s.n)
	result := e(lps(key), msg)
	result.xor512bitInPlace(xor512bit(h, msg))
	return result
}

// e is an E(K, M) function from GOST.
func e(key bit512, msg bit512) bit512 {
	for i := 0; i < 12; i++ {
		msg.xor512bitInPlace(key)
		key.xor512bitInPlace(constants[i])
		msg = lps(msg)
		key = lps(key)
	}
	msg.xor512bitInPlace(key)

	return msg
}
