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

// EmptyElement returns a newly created empty element
func EmptyElement() *Element {
	return CreateFieldElement()
}

// SubtractNonResidue subtracts a from b, placing the result in c.
// c can overlap with a and b.
// This function will not reduce after the operation
// Function: gf_sub_nr
func SubtractNonResidue(c, a, b *Element) {
	c.Limb.MakeMutable()
	defer c.Limb.MakeImmutable()

	subtractRaw(c, a, b)
	bias(c, 2)
}

// AddNonResidue will add a and b and put the result in c.
// c can overlap with a and b
// This function will not reduce at the end.
// Function: gf_add_nr
func AddNonResidue(c, a, b *Element) {
	c.Limb.MakeMutable()
	defer c.Limb.MakeImmutable()

	addRaw(c, a, b)
}

// Subtract will subtract a from b, putting the result in d.
// It is safe for d to overlap with a and b.
// This function reduces at the end
// Function: gf_sub
func Subtract(d, a, b *Element) {
	d.Limb.MakeMutable()
	defer d.Limb.MakeImmutable()

	subtract(d, a, b)
}

// Function: gf_sub
func subtract(d, a, b *Element) {
	subtractRaw(d, a, b)
	bias(d, 2)
	weakReduce(d)
}
