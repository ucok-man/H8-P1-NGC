package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func promptSplitSize(inputsize int) int {
	fmt.Printf("split size (must not be greater than input size '%v'): ", inputsize)

	var input int
	fmt.Scanln(&input)

	return input
}

func promptInputs() []int {
	fmt.Print("inputs numbers (sparated by space): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	inputsStr := strings.Split(scanner.Text(), " ")

	var inputsNum []int
	for _, inputstr := range inputsStr {
		num, err := strconv.Atoi(inputstr)
		if err != nil {
			log.Fatalf("input must be valid number: %v", inputstr)
		}
		inputsNum = append(inputsNum, num)
	}
	return inputsNum
}
