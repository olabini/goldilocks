// +build go1.12

package field

import "github.com/olabini/goldilocks/internal/field/field64"

const WordBits = field64.WordBits
const WordSize = field64.WordSize

type Word field64.Word
type DWord field64.DWord
type SWord field64.SWord
type DSWord field64.DSWord

func WordIsZero(a Word) Word {
	return Word(field64.WordIsZero(field64.Word(a)))
}

func WideMul(a, b Word) (hi, low Word) {
	var hix, lox field64.Word
	hix, lox = field64.WideMul(field64.Word(a), field64.Word(b))
	return Word(hix), Word(lox)
}

func leftShiftExtend(a Word) uint64 {
	return uint64(field64.LeftShiftExtend(field64.Word(a)))
}
