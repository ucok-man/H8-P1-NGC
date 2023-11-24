package main

import (
	"fmt"
	"math"
	"sync"
)

type Circle struct {
	Diameter float64
	Luas     float64
	Keliling float64
}

func process(inputs []float64) <-chan Circle {
	var wg sync.WaitGroup
	var resultch = make(chan Circle)

	for _, input := range inputs {
		wg.Add(1)
		go func(diameter float64) {
			defer wg.Done()

			luas := luasLingkaran(diameter)
			keliling := kelilingLingkaran(diameter)

			resultch <- Circle{
				Diameter: diameter,
				Luas:     luas,
				Keliling: keliling,
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(resultch)
	}()

	return resultch
}

func luasLingkaran(diameter float64) float64 {
	return math.Phi * ((diameter * diameter) / 4)
}

func kelilingLingkaran(diameter float64) float64 {
	return math.Phi * diameter
}

func main() {
	// input diameter lingkaran
	inputs := []float64{100, 200, 300}

	// process
	resultch := process(inputs)

	// print
	for circle := range resultch {
		fmt.Printf("Diameter %.2f\n\tKeliling: %.2f\n\tLuas: %.2f\n\n", circle.Diameter, circle.Keliling, circle.Luas)
	}
}
