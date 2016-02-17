package main

import (
	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
	"github.com/RobbieMcKinstry/UniformRandom/StatTests"
)

func main() {
	_ = RandGenerator.LipschultzGen()
	gen := RandGenerator.NewRANDUGen()
	gen.SetSeed(131)
	ds := RandGenerator.NewDataset(gen, 10000)

	//fmt.Println(StatTests.ChiSquared(ds))
	StatTests.Run(ds, StatTests.MEAN)
	StatTests.Run(ds, StatTests.DIFFERENCE)
	/*
		for i := 0; i < 10000; i++ {
			fmt.Println(gen.Float())

		}
	*/
}
