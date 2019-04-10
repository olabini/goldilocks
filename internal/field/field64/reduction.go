package field64

// WeakReduce will reduce the field represented by a
// Function: gf_weak_reduce
func WeakReduce(a []uint64) {
	mask := uint64((1 << 56) - 1)
	tmp := a[7] >> 56
	a[4] += tmp

	a[7] = (a[7] & mask) + (a[6] >> 56)
	a[6] = (a[6] & mask) + (a[5] >> 56)
	a[5] = (a[5] & mask) + (a[4] >> 56)
	a[4] = (a[4] & mask) + (a[3] >> 56)
	a[3] = (a[3] & mask) + (a[2] >> 56)
	a[2] = (a[2] & mask) + (a[1] >> 56)
	a[1] = (a[1] & mask) + (a[0] >> 56)
	a[0] = (a[0] & mask) + tmp
}
