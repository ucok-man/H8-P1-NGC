package main

import (
	"fmt"
)

func main() {
	// input := rand.New(rand.NewSource(time.Now().UnixMicro())).Intn(10) + 1
	input := 12
	inputCh := make(chan int)

	// execute function
	resultch := factorial(inputCh)

	// send input
	fmt.Println("INPUT :", input)
	inputCh <- input
	close(inputCh)

	// print
	fmt.Println("RESULT FACTORIAL:", <-resultch)
}

func factorial(inputch <-chan int) <-chan int {
	resultch := make(chan int)

	go func() {
		defer close(resultch)

		input := <-inputch

		var sum int = 1
		for i := 1; i <= input; i++ {
			sum *= i
		}

		resultch <- sum
	}()

	return resultch
}
