package field64

// WordBits contains the number of bits in the field
const WordBits = 64

// WordSize contains the number of bytes in the field
const WordSize = WordBits / 8

// Word represents a full word in this representation.
type Word uint64

// type DWord [2]uint64
// type SWord int64
// type DSWord [2]int64

func constEq(x, y uint32) uint32 {
	return uint32((uint64(x^y) - 1) >> 63)
}

// WordIsZero returns 1 if the given argument is 0
// otherwise it returns 0. This function is
// constant time
// Function: word_is_zero
func WordIsZero(a Word) Word {
	topz := constEq(uint32(a>>32), 0)
	bottomz := constEq(uint32(a&0xFFFFFFFF), 0)

	return Word(topz & bottomz)
}

// WideMul bla bla bla
// Function: widemul
// func WideMul(a, b Word) (hi, low Word) {
// 	var hix, lox uint64
// 	hix, lox = bits.Mul64(uint64(a), uint64(b))
// 	return Word(hix), Word(lox)
// }

// LeftShiftExtend creates a mask from the word. If the argument is 1, the mask will
// be all ones, if the argument is 0, it will be all 0.
// This function is constant time, and doesn't reveal the bit inside of the argument
func LeftShiftExtend(a Word) Word {
	result := a & 0x1

	result = (result << 1) | result
	result = (result << 2) | result
	result = (result << 4) | result
	result = (result << 8) | result
	result = (result << 16) | result
	result = (result << 32) | result

	return result
}
