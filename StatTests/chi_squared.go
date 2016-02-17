package StatTests

import (
	"log"

	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
)

func ChiSquared(ds *RandGenerator.Dataset) float64 {

	var buckets []int = make([]int, 10)
	for i := 0; i < 10; i++ {
		buckets[i] = 0
	}

	for i := 0; i < ds.Len(); i++ {
		num := ds.Get(i)
		switch {
		case num < 0.1:
			buckets[0] += 1
		case 0.1 <= num && num < 0.2:
			buckets[1] += 1
		case 0.2 <= num && num < 0.3:
			buckets[2] += 1
		case 0.3 <= num && num < 0.4:
			buckets[3] += 1
		case 0.4 <= num && num < 0.5:
			buckets[4] += 1
		case 0.5 <= num && num < 0.6:
			buckets[5] += 1
		case 0.6 <= num && num < 0.7:
			buckets[6] += 1
		case 0.7 <= num && num < 0.8:
			buckets[7] += 1
		case 0.8 <= num && num < 0.9:
			buckets[8] += 1
		case 0.9 <= num && num < 1.0:
			buckets[9] += 1

		default:
			log.Fatal("Generate produced a value greater than 1.0")
		}
	}

	testStatistic := 0.0
	for _, bucket := range buckets {
		var expected = int(ds.Len() / 10.0)
		testStatistic += float64(((bucket - expected) * (bucket - expected))) / float64((expected))
	}
	return testStatistic
}
