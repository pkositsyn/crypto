package streebog

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOverflow(t *testing.T) {
	var first bit512
	var second bit512

	first[7] = math.MaxUint64
	second[7] = 1

	sum := add512bit(first, second)
	assert.Equal(t, bit512{0, 0, 0, 0, 0, 0, 1, 0}, sum)
}

func TestAddInPlaceOverflow(t *testing.T) {
	var first bit512
	var second bit512

	first[7] = math.MaxUint64
	second[7] = 1

	first.add512bitInPlace(second)
	assert.Equal(t, bit512{0, 0, 0, 0, 0, 0, 1, 0}, first)
}

func TestAddModulus(t *testing.T) {
	var first bit512
	var second bit512

	first[0] = math.MaxUint64
	second[0] = 1

	sum := add512bit(first, second)
	assert.Equal(t, bit512{}, sum)
}

func TestAddInPlaceModulus(t *testing.T) {
	var first bit512
	var second bit512

	first[0] = math.MaxUint64
	second[0] = 1

	first.add512bitInPlace(second)
	assert.Equal(t, bit512{}, first)
}

func TestAddUintOverflow(t *testing.T) {
	var first bit512

	first[7] = math.MaxUint64

	sum := add512bitUint64(first, uint64(3))
	assert.Equal(t, bit512{0, 0, 0, 0, 0, 0, 1, 2}, sum)
}

func TestAddUintInPlaceOverflow(t *testing.T) {
	var first bit512

	first[7] = math.MaxUint64

	first.addUint64InPlace(uint64(3))
	assert.Equal(t, bit512{0, 0, 0, 0, 0, 0, 1, 2}, first)
}
