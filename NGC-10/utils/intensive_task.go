package utils

import "math"

func IntensiveTask() {
	for i := 0; i < 1000000000; i++ {
		_ = math.Sqrt(float64(i)) - math.Sqrt(float64(i))
	}
}
