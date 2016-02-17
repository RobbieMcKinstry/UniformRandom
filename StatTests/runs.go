package StatTests

import (
	"fmt"
	"math"

	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
)

const (
	MEAN       = iota
	DIFFERENCE = iota
)

func Run(ds *RandGenerator.Dataset, option int) int {

	var runs []bool

	switch option {
	case MEAN:
		runs = generateMeanRuns(ds)
	case DIFFERENCE:
		runs = generateDifferenceRuns(ds)
	}

	for _, run := range runs {
		if run {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}

	// TODO do something with *runs* value
	_ = runs
	return 0
}

func generateMeanRuns(ds *RandGenerator.Dataset) []bool {
	runs := make([]bool, ds.Len())

	mean := ds.Mean()

	for i := 0; int(i) < int(ds.Len()); i++ {
		point := ds.Get(i)
		runs[i] = math.Signbit(point - mean)
	}
	return runs
}

func generateDifferenceRuns(ds *RandGenerator.Dataset) []bool {
	runs := make([]bool, ds.Len()-1)

	for i := 0; i < ds.Len()-1; i++ {
		point1 := ds.Get(i)
		point2 := ds.Get(i + 1)

		runs[i] = math.Signbit(point1 - point2)
	}
	return runs
}
