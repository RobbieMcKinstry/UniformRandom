package main

import (
	"fmt"

	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
	"github.com/RobbieMcKinstry/UniformRandom/StatTests"
)

func main() {
	_ = RandGenerator.LipschultzGen()
	gen := RandGenerator.NewRANDUGen()
	gen.SetSeed(122949823)
	ds := RandGenerator.NewDataset(gen, 10000)

	// fmt.Println(StatTests.ChiSquared(ds))
	// mu, variance := StatTests.Autocorrelation(ds, 1, 3)
	// fmt.Printf("Z score: %f", mu/variance)
	fmt.Println(StatTests.Run(ds, StatTests.MEAN))
	fmt.Println(StatTests.Run(ds, StatTests.DIFFERENCE))
	// for i := 0; i < 10000; i++ {
	//	fmt.Println(gen.Float())
	// }
}
