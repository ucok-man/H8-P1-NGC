package main

import (
	"fmt"
)

type Result struct {
	fnName string
	value  int
}

func main() {
	var (
		input  = 2
		result = make(chan Result)
	)

	var taskfns = []func(int, chan Result){SumSquare, SquareSum}

	for _, taskfn := range taskfns {
		go taskfn(input, result)
	}

	for i := 0; i < 2; i++ {
		res := <-result
		fmt.Println(res.fnName, ":", res.value)
	}

}

func SumSquare(input int, result chan Result) {
	var sum int
	for i := 1; i <= input; i++ {
		sum += i
	}
	result <- Result{
		fnName: "SumSquare",
		value:  (sum * sum),
	}
}

func SquareSum(input int, result chan Result) {
	var sum int
	for i := 1; i <= input; i++ {
		sum += (i * i)
	}
	result <- Result{
		fnName: "SquareSum",
		value:  sum,
	}
}
