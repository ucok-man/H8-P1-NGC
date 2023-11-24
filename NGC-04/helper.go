package main

import (
	"bufio"
	"fmt"
	"os"
)

func getinputword(prompt string) (string, error) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error while scanning user input: %v", err)
	}

	return scanner.Text(), nil
}

func getinputline(prompt string) (string, error) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error while scanning user input: %v", err)
	}

	return scanner.Text(), nil
}

func mapfn[T any, Y any](input []T, callback func(param T) (Y, error)) ([]Y, error) {
	var result []Y
	for _, val := range input {
		r, err := callback(val)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}
