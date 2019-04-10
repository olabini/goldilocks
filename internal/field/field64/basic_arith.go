package field64

// Function: gf_sub_RAW
func SubtractRaw(out, a, b []uint64) {
	co1 := uint64((1<<56)-1) * 2
	co2 := co1 - 2
	out[0] = a[0] - b[0] + co1
	out[1] = a[1] - b[1] + co1
	out[2] = a[2] - b[2] + co1
	out[3] = a[3] - b[3] + co1
	out[4] = a[4] - b[4] + co2
	out[5] = a[5] - b[5] + co1
	out[6] = a[6] - b[6] + co1
	out[7] = a[7] - b[7] + co1
	WeakReduce(out)
}

// Function: gf_add_RAW
func AddRaw(out, a, b []uint64) {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	out[4] = a[4] + b[4]
	out[5] = a[5] + b[5]
	out[6] = a[6] + b[6]
	out[7] = a[7] + b[7]
	WeakReduce(out)
}
