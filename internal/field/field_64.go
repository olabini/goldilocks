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
	res, er := e.limbs()
	if er != nil {
		panic(er.Error())
	}
	return res
}

// String represents a string representation of the element
func (e *Element) String() string {
	data := e.limbsUnsafe()
	return fmt.Sprintf("{0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x, 0x%016x}",
		data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7])
}

// Square will square the element in a, and put the result in c.
// The two elements can overlap
func Square(c, a *Element) {
	c.Limb.MakeMutable()
	defer c.Limb.MakeImmutable()

	field64.Square(c.limbsUnsafe(), a.limbsUnsafe())
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
