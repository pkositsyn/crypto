package ctracpkm

import "errors"

var (
	// ErrInvalidOptionsNumber is returned when number of options is greater than 1
	ErrInvalidOptionsNumber = errors.New("number of options must be either 0 or 1")

	// ErrInvalidGammaSize is returned when gamma size does not divide block size
	ErrInvalidGammaSize = errors.New("gamma size must be divisible by cipher block size")

	// ErrInvalidIVSize is returned when IV is not blockSize / 2
	ErrInvalidIVSize = errors.New("invalid iv size - must be half block size")

	// ErrInvalidKeyPressure is returned when key pressure is non-positive
	ErrInvalidKeyPressure = errors.New("ivalid key pressure blocks num - must be positive")
)
