// +build go1.12

package field

import "github.com/olabini/goldilocks/internal/field/field64"

const NLimbs = field64.NLimbs
const LimbBits = field64.LimbBits
const Bits = field64.Bits

var One = CreateElementFrom([]uint64{1, 0, 0, 0, 0, 0, 0, 0})
