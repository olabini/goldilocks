package field64

import . "gopkg.in/check.v1"

type wordSuite struct{}

var _ = Suite(&wordSuite{})

func (s *wordSuite) Test_LeftShiftExtend(c *C) {
	total := LeftShiftExtend(0x1)
	c.Assert(total, Equals, Word(0xFFFFFFFFFFFFFFFF))

	total = LeftShiftExtend(0x0)
	c.Assert(total, Equals, Word(0x00))
}

func (s *wordSuite) Test_WordIsZero(c *C) {
	res := WordIsZero(0x00)
	c.Assert(res, Equals, Word(0x01))

	res = WordIsZero(0xFFFFFFFFFFFFFFFF)
	c.Assert(res, Equals, Word(0x00))

	res = WordIsZero(0x01)
	c.Assert(res, Equals, Word(0x00))
}

// func Test_WideMul(t *testing.T) {
// 	hi, low := WideMul(0x00, 0x00)
// 	if hi != Word(0x00) || low != Word(0x00) {
// 		t.Errorf("WideMul(0x00, 0x00) was incorrect, got 0x%x, 0x%x", hi, low)
// 	}

// 	hi, low = WideMul(0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF)
// 	if hi != Word(0xfffffffffffffffe) || low != Word(0x0000000000000001) {
// 		t.Errorf("WideMul(0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF) was incorrect, got 0x%x, 0x%x", hi, low)
// 	}

// 	hi, low = WideMul(0x42, 0xFFFFFFFF1FFFFFFF)
// 	if hi != Word(0x41) || low != Word(0xffffffc63fffffbe) {
// 		t.Errorf("WideMul(0x42, 0xFFFFFFFF1FFFFFFF) was incorrect, got 0x%x, 0x%x", hi, low)
// 	}
// }
