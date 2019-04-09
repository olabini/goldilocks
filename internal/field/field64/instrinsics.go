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
func widesub(a, b uint128) uint128 {
	// is_wrap_1 := b.hi > a.hi
	// is_wrap_2 := (a.hi == b.hi) & (b.lo > a.lo)
	// is_wrap := is_wrap_1 | is_Wrap_2
	// how to turn this into a mask?

	lo, c := bits.Sub64(a.lo, b.lo, 0)
	hi, c := bits.Sub64(a.hi, b.hi, c)

	// TODO: wraparound
	// - always calculate the possibility of wraparound
	// - use a constant time swap for hi and lo values if
	// - the wraparound actually happens

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

// MulFieldUnsigned multiplies a and b into c
// Function: gf_mulw_unsigned
func MulFieldUnsigned(c, a []uint64, b uint32) {
	mask := uint64((1 << 56) - 1)

	var accum0, accum4 uint128

	for i := 0; i < 4; i++ {
		accum0 = widesum(accum0, widemul(uint64(b), a[i]))
		accum4 = widesum(accum4, widemul(uint64(b), a[i+4]))

		c[i] = accum0.lo & mask
		accum0 = wideshiftright(accum0, 56)

		c[i+4] = accum4.lo & mask
		accum4 = wideshiftright(accum4, 56)
	}

	accum0 = widesum(accum0, widesum(accum4, uint128{hi: 0x00, lo: c[4]}))
	c[4] = accum0.lo & mask
	c[5] += accum0.lo >> 56

	accum4 = widesum(accum4, uint128{hi: 0x00, lo: c[0]})
	c[0] = accum4.lo & mask
	c[1] += accum4.lo >> 56
}

// Square will square the limb a and put it in c
// Function: gf_sqr
func Square(c, a []uint64) {
	mask := uint64((1 << 56) - 1)

	mm, _ := memguard.NewMutable(WordSize * 1 * 4)
	defer mm.Destroy()
	mmb, _ := mm.Uint64()
	aa := mmb[0:4]

	for i := 0; i < 4; i++ {
		aa[i] = a[i] + a[i+4]
	}

	accum2 := widemul(a[0], a[3])
	accum0 := widemul(aa[0], aa[3])
	accum1 := widemul(a[4], a[7])

	accum2 = widesum(accum2, widemul(a[1], a[2]))
	accum0 = widesum(accum0, widemul(aa[1], aa[2]))
	accum1 = widesum(accum1, widemul(a[5], a[6]))

	accum0 = widesub(accum0, accum2)
	accum1 = widesum(accum1, accum2)

	c[3] = (accum1.lo << 1) & mask
	c[7] = (accum0.lo << 1) & mask

	accum0 = wideshiftright(accum0, 55)
	accum1 = wideshiftright(accum1, 55)

	accum0 = widesum(accum0, widemul(2*aa[1], aa[3]))
	accum1 = widesum(accum1, widemul(2*a[5], a[7]))
	accum0 = widesum(accum0, widemul(aa[2], aa[2]))
	accum1 = widesum(accum1, accum0)

	accum0 = widesub(accum0, widemul(2*a[1], a[3]))
	accum1 = widesum(accum1, widemul(a[6], a[6]))

	accum2 = widemul(a[0], a[0])

	accum1 = widesub(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	accum0 = widesub(accum0, widemul(a[2], a[2]))
	accum1 = widesum(accum1, widemul(aa[0], aa[0]))
	accum0 = widesum(accum0, widemul(a[4], a[4]))

	c[0] = accum0.lo & mask
	c[4] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	accum2 = widesum(accum2, widemul(2*aa[2], aa[3]))
	accum0 = widesub(accum0, widemul(2*a[2], a[3]))
	accum1 = widesum(accum1, widemul(2*a[6], a[7]))

	accum1 = widesum(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	accum2 = widemul(2*a[0], a[1])
	accum1 = widesum(accum1, widemul(2*aa[0], aa[1]))
	accum0 = widesum(accum0, widemul(2*a[4], a[5]))

	accum1 = widesub(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	c[1] = accum0.lo & mask
	c[5] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	accum2 = widemul(aa[3], aa[3])
	accum0 = widesub(accum0, widemul(a[3], a[3]))
	accum1 = widesum(accum1, widemul(a[7], a[7]))

	accum1 = widesum(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	accum2 = widemul(2*a[0], a[2])
	accum1 = widesum(accum1, widemul(2*aa[0], aa[2]))
	accum0 = widesum(accum0, widemul(2*a[4], a[6]))

	accum2 = widesum(accum2, widemul(a[1], a[1]))
	accum1 = widesum(accum1, widemul(aa[1], aa[1]))
	accum0 = widesum(accum0, widemul(a[5], a[5]))

	accum1 = widesub(accum1, accum2)
	accum0 = widesum(accum0, accum2)

	c[2] = accum0.lo & mask
	c[6] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	accum0 = widesum(accum0, uint128{hi: 0x00, lo: c[3]})
	accum1 = widesum(accum1, uint128{hi: 0x00, lo: c[7]})

	c[3] = accum0.lo & mask
	c[7] = accum1.lo & mask

	accum0 = wideshiftright(accum0, 56)
	accum1 = wideshiftright(accum1, 56)

	c[4] += accum0.lo + accum1.lo
	c[0] += accum1.lo

}
