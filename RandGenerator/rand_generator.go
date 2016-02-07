package RandGenerator

const (
	_        = iota
	HalfWord = 2 << (32 * iota)
	Word     = 2 << (32 * iota)

	GCCMultiplier = 1103515245
	GCCIncrement  = 12345

	NumericalRecipiesMultipler = 1664525
	NumericalRecipiesIncrement = 1013904223

	VisualMultiplier = 214013
	VisualIncrement  = 2531011

	MMIXMultiplier = 6364136223846793005
	MMIXIncrement  = 1442695040888963407
)

type Generator interface {
	Float() float64
	SetSeed() int64
}

func New(multiplierVal, incrementVal, modulus int64) Generator {
	return &generator{
		x:          1,
		increment:  incrementVal,
		mod:        modulus,
		multiplier: multiplierVal,
	}
}

func NewWithSeed(multiplierVal, incrementVal, modulus, seed int64) Generator {
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

func NewMMIXGen() Generator {
	return &generator{
		x:          1,
		increment:  MMIXIncrement,
		mod:        Word,
		multiplier: MMIXMultiplier,
	}
}

//a = 101427, c = 321, m = 216
func LipschultzGen() Generator {
	return &generator{
		x:          1,
		increment:  321,
		mod:        216,
		multiplier: 101427,
	}
}

func NewRANDUGen() Generator {
	return &generator{
		// a = 65539, c = 0, m = 231
		x:          1,
		increment:  0,
		mod:        231,
		multiplier: 65539,
	}
}

type generator struct {
	x          int64
	increment  int64
	mod        int64
	multiplier int64
}

func (g *generator) SetSeed(seed int64) {
	g.x = seed
}

func (g *generator) Float() float64 {
	// Xi+1 = (aXi + c) mod m. divide xi+1 by m to get the value.
	x1 := (g.multipler * g.x) % g.mod
	g.x = x1

	return (x1 * 1.0) / g.mod
}
