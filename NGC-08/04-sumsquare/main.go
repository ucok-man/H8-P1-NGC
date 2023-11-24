package main

import (
	"fmt"
	"log"
	"sync"
)

func sumOfSquare(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += (val * val)
	}
	return sum
}

func split(splitsize int, inputs []int) (map[int][]int, error) {
	if splitsize > len(inputs) {
		return nil, fmt.Errorf("num of splitsize must be less than inputs length")
	}

	var container = make(map[int][]int)

	// var startidx = 0
	for i := 0; i < len(inputs); i += splitsize {
		var endidx = i + splitsize
		if endidx > len(inputs) {
			endidx = len(inputs)
		}
		container[i] = inputs[i:endidx]
	}

	return container, nil
}

func spawn(inputs map[int][]int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Printf("processing. Spawning %d goroutine...\n", len(inputs))

	var allsum int
	for _, input := range inputs {
		wg.Add(1)
		go func(numbers []int) {
			defer wg.Done()

			sum := sumOfSquare(numbers)

			mu.Lock()
			allsum += sum
			mu.Unlock()

		}(input)
	}

	wg.Wait()
	return allsum
}

func main() {
	inputs := promptInputs()
	splitsize := promptSplitSize(len(inputs))

	splitInput, err := split(splitsize, inputs)
	if err != nil {
		log.Fatal(err)
	}

	result := spawn(splitInput)
	fmt.Println("Sum of square:", result)
}
