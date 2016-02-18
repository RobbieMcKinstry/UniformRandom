package RandGenerator

const (
	_        = iota
	HalfWord = 1 << (32 * iota)

	GCCMultiplier = 1103515245
	GCCIncrement  = 12345

	NumericalRecipiesMultiplier = 1664525
	NumericalRecipiesIncrement  = 1013904223

	VisualMultiplier = 214013
	VisualIncrement  = 2531011

	MMIXMultiplier = 6364136223846793005
	MMIXIncrement  = 1442695040888963407
)

type Generator interface {
	Float64() float64
	Seed(int64)
}

func New(multiplierVal, incrementVal, modulus uint64) Generator {
	return &generator{
		x:          1,
		increment:  incrementVal,
		mod:        modulus,
		multiplier: multiplierVal,
	}
}

func NewWithSeed(multiplierVal, incrementVal, modulus, seed uint64) Generator {
	return &generator{
		x:          seed,
		increment:  incrementVal,
		mod:        modulus,
		multiplier: multiplierVal,
	}

}

func NewGCCGen() Generator {
	return &generator{
		x:          1,
		increment:  GCCIncrement,
		mod:        HalfWord,
		multiplier: GCCMultiplier,
	}
}

func NewNumericalRecipiesGen() Generator {
	return &generator{
		x:          1,
		increment:  NumericalRecipiesIncrement,
		mod:        HalfWord,
		multiplier: NumericalRecipiesMultiplier,
	}
}

func NewVisualGen() Generator {
	return &generator{
		x:          1,
		increment:  VisualIncrement,
		mod:        HalfWord,
		multiplier: VisualMultiplier,
	}
}

func LipschultzGen() Generator {
	return &generator{
		x:          1,
		increment:  321,
		mod:        1 << 16,
		multiplier: 101427,
	}
}

// a = 65539, c = 0, m = 2^31
func NewRANDUGen() Generator {
	return &generator{
		x:          1,
		increment:  0,
		mod:        1 << 31,
		multiplier: 65539,
	}
}

type generator struct {
	x          uint64
	increment  uint64
	mod        uint64
	multiplier uint64
}

func (g *generator) Seed(seed int64) {
	g.x = uint64(seed)
}

func (g *generator) Float64() float64 {
	x1 := (g.multiplier*g.x + g.increment) % g.mod
	g.x = x1

	return float64(x1) / float64(g.mod)
}
