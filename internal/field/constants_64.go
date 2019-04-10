// +build go1.12

package field

import "github.com/olabini/goldilocks/internal/field/field64"

// NLimbs contains the number of limbs for this representation of the Galois Field
const NLimbs = field64.NLimbs

// LimbBits contains the number of bits each limb uses
const LimbBits = field64.LimbBits

// Bits contains the total number of bits an element uses
const Bits = field64.Bits

// One contains the element representing the number one. It should NOT be modified
var One = CreateElementFrom([]uint64{1, 0, 0, 0, 0, 0, 0, 0})
