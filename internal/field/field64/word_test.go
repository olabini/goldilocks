package field64

import "testing"

func Test_LeftShiftExtend(t *testing.T) {
	total := LeftShiftExtend(0x1)
	if total != Word(0xFFFFFFFFFFFFFFFF) {
		t.Errorf("1 mask was incorrect, got %x", total)
	}

	total = LeftShiftExtend(0x0)
	if total != Word(0x00) {
		t.Errorf("0 mask was incorrect, got %x", total)
	}
}

func Test_WordIsZero(t *testing.T) {
	res := WordIsZero(0x00)
	if res != Word(0x01) {
		t.Errorf("WordIsZero(0x00) was incorrect, got 0x%x", res)
	}

	res = WordIsZero(0xFFFFFFFFFFFFFFFF)
	if res != Word(0x00) {
		t.Errorf("WordIsZero(0xFFFFFFFFFFFFFFFF) was incorrect, got 0x%x", res)
	}

	res = WordIsZero(0x01)
	if res != Word(0x00) {
		t.Errorf("WordIsZero(0x01) was incorrect, got 0x%x", res)
	}
}

func Test_WideMul(t *testing.T) {
	hi, low := WideMul(0x00, 0x00)
	if hi != Word(0x00) || low != Word(0x00) {
		t.Errorf("WideMul(0x00, 0x00) was incorrect, got 0x%x, 0x%x", hi, low)
	}

	hi, low = WideMul(0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF)
	if hi != Word(0xfffffffffffffffe) || low != Word(0x0000000000000001) {
		t.Errorf("WideMul(0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF) was incorrect, got 0x%x, 0x%x", hi, low)
	}

	hi, low = WideMul(0x42, 0xFFFFFFFF1FFFFFFF)
	if hi != Word(0x41) || low != Word(0xffffffc63fffffbe) {
		t.Errorf("WideMul(0x42, 0xFFFFFFFF1FFFFFFF) was incorrect, got 0x%x, 0x%x", hi, low)
	}
}
