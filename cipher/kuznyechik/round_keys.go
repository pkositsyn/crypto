package kuznyechik

// constants are initialized alongside with tables
var constants [keyLength]bit128

// generateRoundKeys extracts key sequence from given key
func generateRoundKeys(key []byte) [numRounds]bit128 {
	keyHigh := byte16FromByteSlice(key[0:16])
	keyLow := byte16FromByteSlice(key[16:32])
	return stretchKeys(keyHigh, keyLow)
}

// generateDecryptRoundKeys is a part of decrypt modification. See comment for doDecrypt
func generateDecryptRoundKeys(encryptKeys [numRounds]bit128) (decryptKeys [numRounds]bit128) {
	_ = decryptKeys[numRounds-1]
	decryptKeys[0] = encryptKeys[0]
	for index := 1; index < numRounds; index++ {
		decryptKeys[index][0], decryptKeys[index][1] = inverseL(encryptKeys[index][0], encryptKeys[index][1])
	}
	return
}

// feistel - see https://en.wikipedia.org/wiki/Feistel_cipher for details
func feistel(key, key1, key2 bit128) (bit128, bit128) {
	first := bit128{}
	first[0] = key[0] ^ key1[0]
	first[1] = key[1] ^ key1[1]

	first[0], first[1] = ls(first[0], first[1])

	first[0] ^= key2[0]
	first[1] ^= key2[1]
	return first, key1
}

// stretchKeys generates key sequence for encrypt/decrypt rounds from main key
func stretchKeys(keyHigh, keyLow bit128) (keys [numRounds]bit128) {
	keys[0], keys[1] = keyHigh, keyLow
	constantsIndex := 0
	for i := 2; i < numRounds; i += 2 {
		keys1 := keys[i-2]
		keys2 := keys[i-1]
		for j := 0; j < 8; j++ {
			keys1, keys2 = feistel(constants[constantsIndex], keys1, keys2)
			constantsIndex++
		}
		keys[i], keys[i+1] = keys1, keys2
	}
	return
}
