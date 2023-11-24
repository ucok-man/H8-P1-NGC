package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg       sync.WaitGroup
		rangenum = 100
		inputCh  = make(chan int)
	)

	for i := 1; i <= rangenum; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			fizzbuzz(inputCh)
		}()
		inputCh <- i
	}

	wg.Wait()
}

func fizzbuzz(inputCh chan int) {
	input := <-inputCh
	printfizzbuzz(input)
}

func printfizzbuzz(input int) {
	switch {
	case input%15 == 0:
		fmt.Printf("%d%s\n", input, "fizzbuzz")
	case input%5 == 0:
		fmt.Printf("%d%s", input, "buzz")
	case input%3 == 0:
		fmt.Printf("%d%s", input, "fizz")
	default:
		fmt.Printf("%d ", input)
	}
}
