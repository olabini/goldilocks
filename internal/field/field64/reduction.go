package field64

func WeakReduce(a []uint64) {
	mask := uint64((1 << 56) - 1)
	tmp := a[7] >> 56
	a[4] += tmp
	for i := 7; i > 0; i-- {
		a[i] = (a[i] & mask) + (a[i-1] >> 56)
	}
	a[0] = (a[0] & mask) + tmp
}
