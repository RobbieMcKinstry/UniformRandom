package StatTests

import (
	"math"
)

const (
	MEAN       = iota
	DIFFERENCE = iota
)

func Run(ds *RandGen.Dataset, option int) int {

	var runs []int64

	switch option {
	case MEAN:
		runs = generateMeanRuns(ds)
	case DIFFERENCE:
		runs = generateDifferenceRuns(ds)
	}

	// TODO do something with *runs* value
}

func generateMeanRuns(ds *Dataset) []int64 {
	runs := make([]int, ds.len())

	mean := ds.Mean()

	for i := 0; i < ds.len(); i++ {
		point := ds.Get(i)
		runs[i] = math.Signbit(point - mean)
	}
	return runs
}

func generateDifferenceRuns(ds *Dataset) []int64 {
	runs := make([]int, ds.len()-1)

	for i := 0; i < ds.len()-1; i++ {
		point1 := ds.Get(i)
		point2 := ds.Get(i + 1)

		runs[i] = math.Signbit(point1 - point2)
	}
	return runs
}
