package StatTests

import (
	"math"

	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
)

// An enumeration definiting the kinds of options possible to the user
const (
	MEAN = iota
	DIFFERENCE
)

// Run performs the Wold-Wolfowitz test, taking in the random numbers, and an option specifying how the runs should be calculated.
func Run(ds *RandGenerator.Dataset, option int) float64 {

	var (
		runs     []bool
		expected float64 = 0.0
		observed float64 = 0.0
		variance float64 = 0.0
	)

	switch option {
	case MEAN:
		runs = generateMeanRuns(ds)
		expected, observed, variance = calcOption1Stats(runs)
	case DIFFERENCE:
		runs = generateDifferenceRuns(ds)
		expected, observed, variance = calcOption2Stats(runs)
	}

	return (observed - expected) / variance
}

// Calculates the expected and variance by the first method provided.
// The first option maps the difference between the value and the mean to a binary enum
func calcOption1Stats(runs []bool) (expected, observed, variance float64) {
	var (
		numTrue, numFalse int = 0, 0
		lastVal           bool
	)
	expected = 0.0
	variance = 0.0
	observed = 0.0

	// Count the number of trues and falses
	for i, elem := range runs {
		if elem {
			numTrue++
		} else {
			numFalse++
		}
		// Count the number of changes from true to false
		if i != 0 {
			if lastVal != elem {
				observed++
			}
		}
		lastVal = elem
	}

	// From slides
	expected = float64(2*numTrue*numFalse) / float64(numTrue+numFalse)
	expected += 1.0
	variance = (expected - 1.0) * (expected - 2.0) / float64(numTrue+numFalse+1)
	return expected, observed, variance

}

// Calculates the expected and variance by the second method provided.
// The second method takes the difference between increasing indices and maps the sign of the difference to a binary result
func calcOption2Stats(runs []bool) (expected, observed, variance float64) {
	var (
		lastVal bool
	)

	// More magic numbers from the slides
	observed = 0.0
	expected = float64(2*len(runs)) / 3.0
	variance = float64(16*len(runs)-29) / 90.0

	// Count the number of trues and falses
	for i, elem := range runs {
		// Count the number of changes from true to false
		if i != 0 {
			if lastVal != elem {
				observed++
			}
		}
		lastVal = elem
	}

	return expected, observed, variance

}

// The first option maps the difference between the value and the mean to a binary enum
func generateMeanRuns(ds *RandGenerator.Dataset) []bool {
	runs := make([]bool, ds.Len())

	mean := ds.Mean()

	for i := 0; int(i) < int(ds.Len()); i++ {
		point := ds.Get(i)
		runs[i] = math.Signbit(point - mean)
	}
	return runs
}

// The second method takes the difference between increasing indices and maps the sign of the difference to a binary result
func generateDifferenceRuns(ds *RandGenerator.Dataset) []bool {
	runs := make([]bool, ds.Len()-1)

	for i := 0; i < ds.Len()-1; i++ {
		point1 := ds.Get(i)
		point2 := ds.Get(i + 1)

		runs[i] = math.Signbit(point1 - point2)
	}
	return runs
}
