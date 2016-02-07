package StatTests

import (
	"log"

	"github.com/RobbieMcKinstry/UniformRandom/RandGen"
)

func ChiSquared(gen RandGen.Generator, sampleSize int) int {

	var buckets []int = make([]int, 10)
	for i := 0; i < 10; i++ {
		buckets[i] = 0
	}

	for i := 0; i < sampleSize; i++ {
		num := gen.Float()
		switch {
		case num < 0.1:
			buckets[0] = num
		case 0.1 <= num && num < 0.2:
			buckets[1] = num
		case 0.2 <= num && num < 0.3:
			buckets[2] = num
		case 0.3 <= num && num < 0.4:
			buckets[3] = num
		case 0.4 <= num && num < 0.5:
			buckets[4] = num
		case 0.5 <= num && num < 0.6:
			buckets[5] = num
		case 0.6 <= num && num < 0.7:
			buckets[6] = num
		case 0.7 <= num && num < 0.8:
			buckets[7] = num
		case 0.8 <= num && num < 0.9:
			buckets[8] = num
		case 0.9 <= num && num < 1.0:
			buckets[9] = num

		default:
			log.Fatal("Generate produced a value greater than 1.0")
		}
	}

	testStatistic := 0
	for i, bucket := range buckets {
		var expected = sampleSize / 10.0
		testStatistic += (bucket - expected) * (bucket - expected) / (expected)
	}
	return testStatistic
}
