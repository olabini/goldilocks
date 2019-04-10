package goldilocks

const (
	// EdwardsD represents the D constant in the Edwards representation of the Goldilocks curve
	EdwardsD = int32(-39081)

	// TwistedD represents the D constant in the Twisted Edwards representation of the Goldilocks curve
	TwistedD = EdwardsD - 1

	// EffD represents another D constant for the Goldilocks curve...
	EffD = -TwistedD

	// NegativeD is the negation of the D constant
	NegativeD = int32(1)
)
