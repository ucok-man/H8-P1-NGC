package main

import (
	"fmt"
	"math"
)

func sumOfSquare(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += (val * val)
	}
	return sum
}

func spawn(numOfRoutine int, inputs []int) error {
	if numOfRoutine > len(inputs) {
		return fmt.Errorf("num of routine must be less than inputs length")
	}

	fmt.Println(math.Ceil(float64(len(inputs)) / float64(numOfRoutine)))

	return nil
}

func main() {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(spawn(3, inputs))
}
