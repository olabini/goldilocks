package field64

import (
	"math/bits"

	"github.com/awnumar/memguard"
)

type uint128 struct {
	hi, lo uint64
}

func widemul(a, b uint64) uint128 {
	hi, lo := bits.Mul64(a, b)
	return uint128{hi, lo}
}

// widesum will sum the two addends.
// the result can NOT be larger than what fits
// in an uint128
func widesum(a, b uint128) uint128 {
	lo, c := bits.Add64(a.lo, b.lo, 0)
	hi, c := bits.Add64(a.hi, b.hi, c)
	if c != 0 {
		panic("goldilocks/field/field64/sum: unexpected overflow")
	}
	return uint128{hi, lo}
}

// widesub will subtract b from a
// the result can not be negative - a has to be
// greater than or equal to b
func widesub(a, b uint128) uint128 {
	lo, c := bits.Sub64(a.lo, b.lo, 0)
	hi, c := bits.Sub64(a.hi, b.hi, c)

	if c != 0 {
		panic("goldilocks/field/field64/sub: unexpected underflow")
	}

	return uint128{hi, lo}
}

// wideshiftleft shifts a by k steps to the left
// k can maximum be 64.
func wideshiftleft(a uint128, k uint8) uint128 {
	if k > 64 {
		panic("goldilocks/field/field64/shiftl: too large shift")
	}

	lo := a.lo << k
	hi := a.hi << k

	k2 := uint8(64 - k)
	lo2 := a.lo >> k2

	return uint128{hi | lo2, lo}
}

// wideshiftright shifts a by k steps to the right
// k can maximum be 64.
func wideshiftright(a uint128, k uint8) uint128 {
	if k > 64 {
		panic("goldilocks/field/field64/shiftr: too large shift")
	}

	lo := a.lo >> k
	hi := a.hi >> k

	k2 := uint8(64 - k)
	hi2 := a.hi << k2

	return uint128{hi, lo | hi2}
}

// MulField multiplies a and b into c
// The slices are expected to all be the same size
// which matches the field size.
// Function: gf_mul
func MulField(c, a, b []uint64) {
	mask := uint64((1 << 56) - 1)

	mm, _ := memguard.NewMutable(WordSize * 3 * 4)
	defer mm.Destroy()
	mmb, _ := mm.Uint64()

	aa := mmb[0:4]
	bb := mmb[4:8]
	bbb := mmb[8:12]

	for i := 0; i < 4; i++ {
		aa[i] = a[i] + a[i+4]
		bb[i] = b[i] + b[i+4]
		bbb[i] = bb[i] + b[i+4]
	}

	accum2 := widemul(a[0], b[0])
	accum1 := widemul(aa[0], bb[0])
	accum0 := widemul(a[4], b[4])

	accum2 = widesum(accum2, widemul(a[1], b[7]))
	accum1 = widesum(accum1, widemul(aa[1], bbb[3]))
	accum0 = widesum(accum0, widemul(a[5], bb[3]))

	accum2 = widesum(accum2, widemul(a[2], b[6]))
	accum1 = widesum(accum1, widemul(aa[2], bbb[2]))
	accum0 = widesum(accum0, widemul(a[6], bb[2]))

	accum2 = widesum(accum2, widemul(a[3], b[5]))
	accum1 = widesum(accum1, widemul(aa[3], bbb[1]))
	accum0 = widesum(accum0, widemul(a[7], bb[1]))

	accum1 = widesub(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	c[0] = accum0.lo & mask
	c[4] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	accum2 = widemul(a[0], b[1])
	accum1 = widesum(accum1, widemul(aa[0], bb[1]))
	accum0 = widesum(accum0, widemul(a[4], b[5]))

	accum2 = widesum(accum2, widemul(a[1], b[0]))
	accum1 = widesum(accum1, widemul(aa[1], bb[0]))
	accum0 = widesum(accum0, widemul(a[5], b[4]))

	accum2 = widesum(accum2, widemul(a[2], b[7]))
	accum1 = widesum(accum1, widemul(aa[2], bbb[3]))
	accum0 = widesum(accum0, widemul(a[6], bb[3]))

	accum2 = widesum(accum2, widemul(a[3], b[6]))
	accum1 = widesum(accum1, widemul(aa[3], bbb[2]))
	accum0 = widesum(accum0, widemul(a[7], bb[2]))

	accum1 = widesub(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	c[1] = accum0.lo & mask
	c[5] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	accum2 = widemul(a[0], b[2])
	accum1 = widesum(accum1, widemul(aa[0], bb[2]))
	accum0 = widesum(accum0, widemul(a[4], b[6]))

	accum2 = widesum(accum2, widemul(a[1], b[1]))
	accum1 = widesum(accum1, widemul(aa[1], bb[1]))
	accum0 = widesum(accum0, widemul(a[5], b[5]))

	accum2 = widesum(accum2, widemul(a[2], b[0]))
	accum1 = widesum(accum1, widemul(aa[2], bb[0]))
	accum0 = widesum(accum0, widemul(a[6], b[4]))

	accum2 = widesum(accum2, widemul(a[3], b[7]))
	accum1 = widesum(accum1, widemul(aa[3], bbb[3]))
	accum0 = widesum(accum0, widemul(a[7], bb[3]))

	accum1 = widesub(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	c[2] = accum0.lo & mask
	c[6] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	accum2 = widemul(a[0], b[3])
	accum1 = widesum(accum1, widemul(aa[0], bb[3]))
	accum0 = widesum(accum0, widemul(a[4], b[7]))

	accum2 = widesum(accum2, widemul(a[1], b[2]))
	accum1 = widesum(accum1, widemul(aa[1], bb[2]))
	accum0 = widesum(accum0, widemul(a[5], b[6]))

	accum2 = widesum(accum2, widemul(a[2], b[1]))
	accum1 = widesum(accum1, widemul(aa[2], bb[1]))
	accum0 = widesum(accum0, widemul(a[6], b[5]))

	accum2 = widesum(accum2, widemul(a[3], b[0]))
	accum1 = widesum(accum1, widemul(aa[3], bb[0]))
	accum0 = widesum(accum0, widemul(a[7], b[4]))

	accum1 = widesub(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	c[3] = accum0.lo & mask
	c[7] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	accum0 = widesum(accum0, accum1)
	accum0 = widesum(accum0, uint128{hi: 0x00, lo: c[4]})
	accum1 = widesum(accum1, uint128{hi: 0x00, lo: c[0]})

	c[4] = accum0.lo & mask
	c[0] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	c[5] += accum0.lo
	c[1] += accum1.lo
}
