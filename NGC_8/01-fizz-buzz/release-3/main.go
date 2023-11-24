package main

import (
	"bytes"
	"fmt"
	"sync"
)

type Result struct {
	mu        *sync.Mutex
	fizzbuzz  bytes.Buffer
	sum       int
	totalodd  int
	totaleven int
}

func main() {
	var (
		wg       sync.WaitGroup
		result   = Result{mu: &sync.Mutex{}}
		inputCh  = make(chan int)
		resultCh = make(chan *Result)
	)

	for i := 1; i <= 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			process(inputCh, resultCh)
		}()
		inputCh <- i
		resultCh <- &result
	}

	wg.Wait()

	fmt.Println(result.fizzbuzz.String())
	fmt.Println("\nSUM:", result.sum)
	fmt.Println("TOTAL EVEN:", result.totaleven)
	fmt.Println("TOTAL ODD:", result.totalodd)
}

func process(inputCh chan int, resultCh chan *Result) {
	input := <-inputCh
	result := <-resultCh

	result.mu.Lock()
	defer result.mu.Unlock()

	result.fizzbuzz.WriteString(fizzbuzz(input))
	result.sum += input

	if isEven(input) {
		result.totaleven += 1
	} else {
		result.totalodd += 1
	}
}

func fizzbuzz(input int) string {
	switch {
	case input%15 == 0:
		return fmt.Sprintf("%d%s ", input, "fizzbuzz")
	case input%5 == 0:
		return fmt.Sprintf("%d%s ", input, "buzz")
	case input%3 == 0:
		return fmt.Sprintf("%d%s ", input, "fizz")
	default:
		return fmt.Sprintf("%d ", input)
	}
}

func isEven(input int) bool {
	return input%2 == 0
}
