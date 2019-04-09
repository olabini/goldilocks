package field64

import (
	"testing"
)

type testVector struct {
	a, b []uint64
	exp  []uint64
}

var mulFieldTestVectors = []testVector{
	{[]uint64{0x0000000000000000,
		0xFF00000000000000,
		0x0000000000000000,
		0x0001100000000000,
		0x0000000000000000,
		0x0000000023423400,
		0x0000000000000000,
		0x0123240000000000},
		[]uint64{0x0000000000000001,
			0x0000000000000002,
			0x0000000000000003,
			0x0000000000000004,
			0x0000000000000005,
			0x0000000000000006,
			0x0000000000000007,
			0x0000000000000008},
		[]uint64{0x1f8001a71a76ff,
			0x66d800b04b0c01,
			0xae3000d38d390a,
			0xb0c400f6cf6e0b,
			0xf47802c12c1a01,
			0x600400d38d4403,
			0xcb90011a11a50e,
			0xd8280160960e10}},
	{[]uint64{0x20045d1452af094, 0x293bcc10b8e55fe,
		0x2f19b47d0fb5087, 0x24437b2944be0f5,
		0x290bc60577dc5d6, 0x383a28e9270a806,
		0x2ea9ca836c3a1b1, 0x3589540a8cd290e},
		[]uint64{0x20045d1452af094, 0x293bcc10b8e55fe,
			0x2f19b47d0fb5087, 0x24437b2944be0f5,
			0x290bc60577dc5d6, 0x383a28e9270a806,
			0x2ea9ca836c3a1b1, 0x3589540a8cd290e},
		[]uint64{0x95832d37a849e7, 0x41dae0c74dbcbc,
			0x9ccd7c7f37634c, 0x39d374dba29ae1,
			0x9b6621f97180e3, 0x4c1fcd23e399ce,
			0xdad4c9f461298f, 0xf3961d4386b5f2}},
	{[]uint64{0x2f5e2772ab90930, 0x348b8d80b83d373,
		0x36df47e963cdc83, 0x3206afcd7f3d225,
		0x422a89750551b34, 0x468a2fbc5f38d68,
		0x2726803e45153ab, 0x412fcc348519777},
		[]uint64{0x2f5e2772ab90930, 0x348b8d80b83d373,
			0x36df47e963cdc83, 0x3206afcd7f3d225,
			0x422a89750551b34, 0x468a2fbc5f38d68,
			0x2726803e45153ab, 0x412fcc348519777},
		[]uint64{0x4df55532373f66, 0x76395f5ec4c2f7,
			0xea5443fd3997b1, 0x3594eaef164b34,
			0x2f88a5d0f78390, 0xf3530df4fd7198,
			0x7864723f2fd84, 0x918cfbfde10a1c}},
}

func Test_MulField(t *testing.T) {
	out := make([]uint64, NLimbs)
	for ix, v := range mulFieldTestVectors {
		MulField(out, v.a, v.b)
		if out[0] != v.exp[0] ||
			out[1] != v.exp[1] ||
			out[2] != v.exp[2] ||
			out[3] != v.exp[3] ||
			out[4] != v.exp[4] ||
			out[5] != v.exp[5] ||
			out[6] != v.exp[6] ||
			out[7] != v.exp[7] {
			t.Errorf("MulField(#%d) was incorrect, got %#v", ix, out)
		}
	}
}

func Test_MulFieldUnsigned(t *testing.T) {
	out := make([]uint64, NLimbs)
	a := []uint64{0x993a1dc07dc468, 0xb30cb17de4df36,
		0xc383cc018dd112, 0x16a4de8f81ceab,
		0x4dcff3b827e40, 0x7ce47f119bb569,
		0xd7ad6a3cc71dca, 0x4971fec10a566}
	b := uint32(0x13154)
	exp := []uint64{0x84841900445f99, 0xdfa13afd094a78,
		0x2d4cca789fe174, 0xdb4200bdd2744c,
		0xe499a5ff3bcd7a, 0x365dac4a39a540,
		0x6079452bd4053d, 0x8e9eb94a85ccb4}
	MulFieldUnsigned(out, a, b)
	if out[0] != exp[0] ||
		out[1] != exp[1] ||
		out[2] != exp[2] ||
		out[3] != exp[3] ||
		out[4] != exp[4] ||
		out[5] != exp[5] ||
		out[6] != exp[6] ||
		out[7] != exp[7] {
		t.Errorf("MulFieldUnsigned(1) was incorrect, got %#v", out)
	}

	a = []uint64{0xe908e8219da07c, 0x932e20bca1bb24,
		0x5fc4c0d1a8cfc0, 0x822dd3800306c7,
		0x4d1da688efac17, 0x74f5826e83eafa,
		0xede9810ca42ae7, 0x3353ab6d645941}
	b = uint32(0x13154)
	exp = []uint64{0x137c37d3eca1e7, 0x40237a950861c0,
		0xd6490eeae89a8a, 0x5004e99c11c284,
		0xa123ca7aacb706, 0xe0f4776e3d5001,
		0x435e67bcfbd64a, 0x86b1947344d615}
	MulFieldUnsigned(out, a, b)
	if out[0] != exp[0] ||
		out[1] != exp[1] ||
		out[2] != exp[2] ||
		out[3] != exp[3] ||
		out[4] != exp[4] ||
		out[5] != exp[5] ||
		out[6] != exp[6] ||
		out[7] != exp[7] {
		t.Errorf("MulFieldUnsigned(2) was incorrect, got %#v", out)
	}
}

func limbEq(a, b []uint64) bool {
	return a[0] == b[0] &&
		a[1] == b[1] &&
		a[2] == b[2] &&
		a[3] == b[3] &&
		a[4] == b[4] &&
		a[5] == b[5] &&
		a[6] == b[6] &&
		a[7] == b[7]
}

func Test_Square(t *testing.T) {
	out := make([]uint64, NLimbs)
	out2 := make([]uint64, NLimbs)
	out3 := make([]uint64, NLimbs)
	out4 := make([]uint64, NLimbs)
	a := []uint64{0x85222f93790836, 0x3c988366cff42b,
		0xcd54e688f90ea4, 0xcbb5c2f153147b,
		0x4878cc0448beff, 0x598976a599e589,
		0xb68b1de2c8aaa3, 0x2fe2e3722d6d5e}
	exp := []uint64{0x453f90677f6ac2, 0x4435fce2425b2b,
		0x59a92c8ccae3fb, 0xf59bc85fb950ac,
		0x37343e032e48db, 0x337481e66f0a6e,
		0x890dc78d47073f, 0x4795ac2b49dbfa}
	Square(out, a)
	if !limbEq(out, exp) {
		t.Errorf("Square(0) was incorrect, got %#v", out)
	}

	MulField(out2, a, a)
	if !limbEq(out, out2) {
		t.Errorf("Mul(0) and Square(0) should be equal...")
	}

	for i := 0; i < 30; i += 2 {
		Square(out3, out)
		MulField(out4, out2, out2)

		if !limbEq(out3, out4) {
			t.Errorf("Mul(%d) and Square(%d) should be equal...", i+1, i+1)
		}

		Square(out, out3)
		MulField(out2, out4, out4)

		if !limbEq(out, out2) {
			t.Errorf("Mul(%d) and Square(%d) should be equal...", i+2, i+2)
		}
	}
}

func Test_widesub(t *testing.T) {
	one := uint128{0x00, 0x01}
	zero := uint128{0x00, 0x00}
	res := widesub(one, one)
	if res != zero {
		t.Errorf("widesub(0x01, 0x01) was incorrect, got %#v", res)
	}

	res = widesub(one, zero)
	if res != one {
		t.Errorf("widesub(0x01, 0x00) was incorrect, got %#v", res)
	}

	high_one := uint128{0x01, 0x01}
	res = widesub(high_one, one)
	if res != (uint128{0x01, 0x00}) {
		t.Errorf("widesub(0x010000000000000001, 0x01) was incorrect, got %#v", res)
	}

	another := uint128{0x01, 0x00}
	res = widesub(another, one)
	if res != (uint128{0x00, 0xFFFFFFFFFFFFFFFF}) {
		t.Errorf("widesub(0x010000000000000000, 0x01) was incorrect, got %#v", res)
	}

	large := uint128{0xFFFFFFFFFDD123, 0x00FF124323245324}
	other_large := uint128{0xAAA, 0x0000000656456454}
	res = widesub(large, other_large)
	if res != (uint128{0xfffffffffdc679, 0x00ff123cccdeeed0}) {
		t.Errorf("widesub(0xFFFFFFFFFDD12300FF124323245324, 0xAAA0000000656456454) was incorrect, got %#v", res)
	}

	//TODO: wraparound...
}

func Test_wideshiftleft(t *testing.T) {
	one := uint128{0x00, 0x01}
	res := wideshiftleft(one, 0)
	if res != one {
		t.Errorf("wideshiftleft(0x01, 0) was incorrect, got %#v", res)
	}

	res = wideshiftleft(one, 10)
	if res != (uint128{0x00, 0x0400}) {
		t.Errorf("wideshiftleft(0x01, 10) was incorrect, got %#v", res)
	}

	res = wideshiftleft(one, 63)
	if res != (uint128{0x00, 0x8000000000000000}) {
		t.Errorf("wideshiftleft(0x01, 63) was incorrect, got %#v", res)
	}

	res = wideshiftleft(one, 64)
	if res != (uint128{0x01, 0x00}) {
		t.Errorf("wideshiftleft(0x01, 64) was incorrect, got %#v", res)
	}

	res = wideshiftleft(uint128{0x01, 0x01}, 1)
	if res != (uint128{0x02, 0x02}) {
		t.Errorf("wideshiftleft(0x010000000000000001, 1) was incorrect, got %#v", res)
	}
}

func Test_wideshiftright(t *testing.T) {
	one := uint128{0x00, 0x01}
	res := wideshiftright(one, 0)
	if res != one {
		t.Errorf("wideshiftright(0x01, 0) was incorrect, got %#v", res)
	}

	res = wideshiftright(one, 1)
	if res != (uint128{0x00, 0x00}) {
		t.Errorf("wideshiftright(0x00, 1) was incorrect, got %#v", res)
	}

	res = wideshiftright(uint128{0xFF00FF, 0x123}, 64)
	if res != (uint128{0x00, 0xFF00FF}) {
		t.Errorf("wideshiftright(0xFF00FF0000000000000123, 64) was incorrect, got %#v", res)
	}

	res = wideshiftright(uint128{0xFF00FF, 0x123}, 1)
	if res != (uint128{0x7f807f, 0x8000000000000091}) {
		t.Errorf("wideshiftright(0xFF00FF0000000000000123, 1) was incorrect, got %#v", res)
	}
}
