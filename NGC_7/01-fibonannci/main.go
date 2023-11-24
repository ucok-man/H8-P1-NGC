package main

import (
	"fmt"
	"sync"
)

func fibonnaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonnaci(n-1) + fibonnaci(n-2)
}

func syncFibonacci() {
	for i := 1; i <= 10; i++ {
		result := fmt.Sprintf("fibonannci ke %d : %d", i, fibonnaci(i))
		fmt.Println(result)
	}
}

func asyncFibonacci() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			result := fmt.Sprintf("fibonannci ke %d : %d", i, fibonnaci(i))
			fmt.Println(result)
		}(i)
	}

	wg.Wait()
}

func main() {
	fmt.Println("SYNC")
	syncFibonacci()

	fmt.Print("\n\n")
	fmt.Println("ASYNC")
	asyncFibonacci()
}
