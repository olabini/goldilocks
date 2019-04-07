package field64

import (
	"math/bits"
)

const WordBits = 64
const WordSize = WordBits / 8

type Word uint64
type DWord [2]uint64
type SWord int64
type DSWord [2]int64

func constEq(x, y uint32) uint32 {
	return uint32((uint64(x^y) - 1) >> 63)
}

// WordIsZero returns 1 if the given argument is 0
// otherwise it returns 0. This function is
// constant time
// This function maps to word_is_zero
func WordIsZero(a Word) Word {
	topz := constEq(uint32(a>>32), 0)
	bottomz := constEq(uint32(a&0xFFFFFFFF), 0)

	return Word(topz & bottomz)
}

// WideMul bla bla bla
// This function maps to widemul
func WideMul(a, b Word) (hi, low Word) {
	var hix, lox uint64
	hix, lox = bits.Mul64(uint64(a), uint64(b))
	return Word(hix), Word(lox)
}

// LeftShiftExtend bla bla bla
func LeftShiftExtend(a Word) Word {
	ext := a & 0x1
	result := Word(0)

	for i := 0; i < 64; i++ {
		result = (result << 1) | ext
	}

	return result
}
