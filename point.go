package goldilocks

import (
	"github.com/olabini/goldilocks/internal/field"
)

// We will do two interfaces - one that is functional, and one that is OO, based on the functional one

// First functionality to put in:
// - goldilocks_448_point_eq
// - goldilocks_448_point_double
// - goldilocks_448_point_sub
// - goldilocks_448_point_negate
// - goldilocks_448_point_scalarmul
// - goldilocks_448_point_double_scalarmul
// - goldilocks_448_point_dual_scalarmul
// - goldilocks_448_point_valid

type point struct {
	x, y, z, t *field.Element /* Twisted extended homogeneous coordinates */
}

func newPoint() *point {
	return &point{
		x: field.EmptyElement(),
		y: field.EmptyElement(),
		z: field.EmptyElement(),
		t: field.EmptyElement(),
	}
}

// Function: goldilocks_448_point_add
func pointAdd(p, q, r *point) {
	a := field.EmptyElement()
	defer a.Destroy()
	b := field.EmptyElement()
	defer b.Destroy()
	c := field.EmptyElement()
	defer c.Destroy()
	d := field.EmptyElement()
	defer d.Destroy()

	field.SubtractNonResidue(b, q.y, q.x) // 3+e
	field.SubtractNonResidue(c, r.y, r.x) // 3+e
	field.AddNonResidue(d, r.y, r.x)      // 2+e
	field.Mul(a, c, b)
	field.AddNonResidue(b, q.y, q.x) // 2+e
	field.Mul(p.y, d, b)
	field.Mul(b, r.t, q.t)
	field.MulSigned(p.x, b, 2*EffD)
	field.AddNonResidue(b, a, p.y)      // 2+e
	field.SubtractNonResidue(c, p.y, a) // 3+e
	field.Mul(a, q.z, r.z)
	field.AddNonResidue(a, a, a)        // 2+e
	field.AddNonResidue(p.y, a, p.x)    // 3+e or 2+e
	field.SubtractNonResidue(a, a, p.x) // 4+e or 3+e
	field.Mul(p.z, a, p.y)
	field.Mul(p.x, p.y, c)
	field.Mul(p.y, a, b)
	field.Mul(p.t, b, c)
}
