package field

import "github.com/awnumar/memguard"

// Element contains the limbs for fields
// It is the equivalent of gf_448_s, gf_448_p and gf
// in the libgoldilocks source
type Element struct {
	Limb *memguard.LockedBuffer
}

// CreateFieldElement will return a newly created, empty field element
func CreateFieldElement() *Element {
	var e Element
	e.Limb, _ = memguard.NewImmutable(NLimbs * WordSize)
	return &e
}

// Destroy will release the memory used by the field element. This should
// be done as soon as an element is done being used.
func (e *Element) Destroy() {
	e.Limb.Destroy()
}

// SquareN will square x, n times
// Function: gf_sqrn
func SquareN(y, x *Element, n int) {
	tmp := CreateFieldElement()
	defer tmp.Destroy()

	if n&1 != 0 {
		Square(y, x)
		n--
	} else {
		Square(tmp, x)
		Square(y, tmp)
		n -= 2
	}
	for ; n > 0; n -= 2 {
		Square(tmp, y)
		Square(y, tmp)
	}
}
