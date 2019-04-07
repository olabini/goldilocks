// +build go1.12

package field

// limbs will return a slice of the appropriate
// data type for the field size used.
func (e *Element) limbs() ([]uint64, error) {
	return e.Limb.Uint64()
}
