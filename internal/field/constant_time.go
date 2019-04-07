package field

// Will implement these when we need them
//lookup
//insert
//select

// ConditionalSwap will swap all data in a and b in
// constant time. The swap argument is either 1 or 0.
// This function maps to constant_time_cond_swap
func ConditionalSwap(a, b *Element, swap Word) {
	mask := leftShiftExtend(swap)

	a.Limb.MakeMutable()
	defer a.Limb.MakeImmutable()
	b.Limb.MakeMutable()
	defer b.Limb.MakeImmutable()

	ab, _ := a.limbs()
	bb, _ := b.limbs()

	for i := range ab {
		s := (ab[i] ^ bb[i]) & mask
		ab[i] ^= s
		bb[i] ^= s
	}
}
