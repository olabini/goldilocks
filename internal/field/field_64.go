// +build go1.12

package field

import (
	"fmt"

	"github.com/olabini/goldilocks/internal/field/field64"
)

// limbs will return a slice of the appropriate
// data type for the field size used.
func (e *Element) limbs() ([]uint64, error) {
	return e.Limb.Uint64()
}

// limbsUnsafe assumes the limbs are set up correctly
// and will panic if that's not correct
func (e *Element) limbsUnsafe() []uint64 {
	res, er := e.Limb.Uint64()
	if er != nil {
		panic(er.Error())
	}
	return res
}

// String represents a string representation of the element
func (e *Element) String() string {
	data, _ := e.limbs()
	return fmt.Sprintf("{0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x}",
		data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7])
}

// Square will square the element in a, and put the result in c.
// The two elements can overlap
func Square(c, a *Element) {
	c.Limb.MakeMutable()
	defer c.Limb.MakeImmutable()

	climbs, _ := c.limbs()
	alimbs, _ := a.limbs()

	field64.Square(climbs, alimbs)
}

// CreateElementFrom takes NLimbs uint64 entries and creates
// an element from them. It panics if given the wrong amount of
// data
func CreateElementFrom(data []uint64) *Element {
	if len(data) != NLimbs {
		panic(fmt.Sprintf("CreateElementFrom called with %d limbs", len(data)))
	}

	elm := CreateFieldElement()
	elm.Limb.MakeMutable()
	defer elm.Limb.MakeImmutable()

	lmbs, _ := elm.limbs()
	copy(lmbs, data)

	return elm
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

// Function: gf_sub_RAW
func subtractRaw(out, a, b *Element) {
	field64.SubtractRaw(out.limbsUnsafe(), a.limbsUnsafe(), b.limbsUnsafe())
}

// Function: gf_add_RAW
func addRaw(out, a, b *Element) {
	field64.AddRaw(out.limbsUnsafe(), a.limbsUnsafe(), b.limbsUnsafe())
}

// Function: gf_bias
func bias(*Element, int) {
	// empty - this implementation dosen't do biasing
}

// Function: gf_weak_reduce
func weakReduce(a *Element) {
	field64.WeakReduce(a.limbsUnsafe())
}

// Mul will multiply the as and bs elements, putting the result in
// cs. It's safe for cs to overlap with bs and as.
// Function: gf_mul
func Mul(cs, as, bs *Element) {
	cs.Limb.MakeMutable()
	defer cs.Limb.MakeImmutable()

	field64.MulField(cs.limbsUnsafe(), as.limbsUnsafe(), bs.limbsUnsafe())
}

// MulSigned will multiply by a signed integer. It is not
// constant time with regard to that integer
// Function: gf_mulw
func MulSigned(c, a *Element, w int32) {
	c.Limb.MakeMutable()
	defer c.Limb.MakeImmutable()

	if w > 0 {
		field64.MulFieldUnsigned(c.limbsUnsafe(), a.limbsUnsafe(), uint32(w))
	} else {
		field64.MulFieldUnsigned(c.limbsUnsafe(), a.limbsUnsafe(), uint32(-w))
		subtract(c, Zero, c)
	}
}
