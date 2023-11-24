package main

import (
	"fmt"
	"math"
	"sync"
)

const (
	RECTANGLE = "RECTANGLE"
	TRIANGLE  = "TRIANGLE"
	CIRCLE    = "CIRCLE"
)

type Shape struct {
	ShapeType string
	Length    int
	Area      float32
}

func calculateArea(input Shape) Shape {
	switch input.ShapeType {
	case RECTANGLE:
		input.Area = float32(input.Length) * float32(input.Length)
	case TRIANGLE:
		input.Area = 0.5 * float32(input.Length) * float32(input.Length)
	case CIRCLE:
		input.Area = math.Phi * (0.25 * float32(input.Length))
	default:
		panic(fmt.Sprintf("unhandled posible case %v\n", input.ShapeType))
	}

	return input
}

func main() {
	var (
		wg       sync.WaitGroup
		inputch  = make(chan Shape)
		resultch = make(chan Shape)
	)

	inputs := []Shape{
		{ShapeType: RECTANGLE, Length: 5},
		{ShapeType: CIRCLE, Length: 3},
		{ShapeType: TRIANGLE, Length: 5},
		{ShapeType: RECTANGLE, Length: 15},
		{ShapeType: CIRCLE, Length: 5},
	}

	for _, shape := range inputs {
		wg.Add(1)
		go func(inputch <-chan Shape) {
			defer wg.Done()

			input := <-inputch

			result := calculateArea(input)
			resultch <- result
		}(inputch)
		inputch <- shape
	}

	go func() {
		for res := range resultch {
			fmt.Printf("%s: LENGTH %d | AREA %.2f\n", res.ShapeType, res.Length, res.Area)
		}
	}()

	wg.Wait()
	close(inputch)
	close(resultch)
}
