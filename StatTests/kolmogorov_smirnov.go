package StatTests

import (
	"math"

	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
)

func KolmogorovSmirnov(ds *RandGenerator.Dataset) float64 {
	ds = ds.Subset(0, 100).Sort() // only select the first 100 numbers, and sort them
	max := 0.0

	// Find the maximum difference from the expected value to the observed value
	for i := 0; i < ds.Len(); i++ {
		observed := ds.Get(i)
		expected := float64(i) / float64(ds.Len())

		distence := math.Abs(observed - expected)
		if distence > max {
			max = distence
		}
	}

	return max
}
