package field

import "github.com/awnumar/memguard"

// Element contains the limbs for fields
// It is the equivalent of gf_448_s, gf_448_p and gf
// in the libgoldilocks source
type Element struct {
	Limb *memguard.LockedBuffer
}

func CreateFieldElement() *Element {
	var e Element
	e.Limb, _ = memguard.NewImmutable(NLimbs * WordSize)
	return &e
}

func (e *Element) Destroy() {
	e.Limb.Destroy()
}

// SquareN will square x, n times
// Function: gf_sqrn
func SquareN(y, x *Element, n int) {
	tmp := CreateFieldElement()
	defer tmp.Destroy()

	if n&1 != 0 {
		Square(y, x)
		n--
	} else {
		Square(tmp, x)
		Square(y, tmp)
		n -= 2
	}
	for ; n > 0; n -= 2 {
		Square(tmp, y)
		Square(y, tmp)
	}
}

// #define gf_add_nr gf_add_RAW

// /** Subtract mod p.  Bias by 2 and don't reduce  */
// static inline void gf_sub_nr ( gf c, const gf a, const gf b ) {
//     gf_sub_RAW(c,a,b);
//     gf_bias(c, 2);
//     if (GF_HEADROOM < 3) gf_weak_reduce(c);
// }

// /** Subtract mod p. Bias by amt but don't reduce.  */
// static inline void gf_subx_nr ( gf c, const gf a, const gf b, int amt ) {
//     gf_sub_RAW(c,a,b);
//     gf_bias(c, amt);
//     if (GF_HEADROOM < amt+1) gf_weak_reduce(c);
// }

// /** Mul by signed int.  Not constant-time WRT the sign of that int. */
// static inline void gf_mulw(gf c, const gf a, int32_t w) {
//     if (w>0) {
//         gf_mulw_unsigned(c, a, w);
//     } else {
//         gf_mulw_unsigned(c, a, -w);
//         gf_sub(c,ZERO,c);
//     }
// }

// /** Constant time, x = is_z ? z : y */
// static inline void gf_cond_sel(gf x, const gf y, const gf z, mask_t is_z) {
//     constant_time_select(x,y,z,sizeof(gf),is_z,0);
// }

// /** Constant time, if (neg) x=-x; */
// static inline void gf_cond_neg(gf x, mask_t neg) {
//     gf y;
//     gf_sub(y,ZERO,x);
//     gf_cond_sel(x,x,y,neg);
// }

// /** Constant time, if (swap) (x,y) = (y,x); */
// static inline void
// gf_cond_swap(gf x, gf_s *__restrict__ y, mask_t swap) {
//     constant_time_cond_swap(x,y,sizeof(gf_s),swap);
// }

// static GOLDILOCKS_INLINE void gf_mul_qnr(gf_s *__restrict__ out, const gf x) {
// #if P_MOD_8 == 7
//     gf_sub(out,ZERO,x);
// #else
//     #error "Only supporting p=7 mod 8"
// #endif
// }

// static GOLDILOCKS_INLINE void gf_div_qnr(gf_s *__restrict__ out, const gf x) {
// #if P_MOD_8 == 7
//     gf_sub(out,ZERO,x);
// #else
//     #error "Only supporting p=7 mod 8"
// #endif
// }
