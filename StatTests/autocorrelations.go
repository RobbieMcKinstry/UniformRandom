package StatTests

import (
	"math"

	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
)

// Autocorrelation tests whether or not a generator produces independent results.
// The parameters are the dataset ds,
// start, the first index looked at by the test
// and jump, the ith index looked starting at (start)
func Autocorrelation(ds *RandGenerator.Dataset, start, jump int) (float64, float64) {

	// Index1 and Index2 keep tracks of when to break from the loop, as the array will be out of bounds when index2 is exceeds M
	var (
		index1 = start
		index2 = start + jump
		total  = 0.0

		k int
	)

	// Summation of each of the potentially correlated indices
	for k = 0; index2 < ds.Len(); k++ {
		total += ds.Get(index1) * ds.Get(index2)
		index1 = start + jump*k
		index2 = start + jump*(k+1)
	}

	// Now, take the total and divide it by the number of iterations + 1
	mu := total / float64((k + 1))

	mu -= 0.25 // Magic number provided by Professor Lipschultz in the slides

	// calculate the variance
	squaredVal := float64(13*k + 7)
	demoninator := float64(12 * (k + 1))
	variance := math.Sqrt(squaredVal) / demoninator
	return mu, variance
}
