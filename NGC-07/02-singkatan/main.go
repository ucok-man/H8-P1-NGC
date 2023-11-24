package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var inputs []string
	for {
		input, err := prompt("enter word ('q' for break) : ")
		if err != nil {
			log.Fatal(err)
		}

		input = strings.ToLower(input)
		if input == "q" {
			break
		}
		inputs = append(inputs, input)
	}

	var results []string
	for _, input := range inputs {
		wg.Add(1)
		go persingkatFN(&wg, input, &results)
	}

	wg.Wait()

	fmt.Println()
	for _, val := range results {
		fmt.Println(val)
	}
}

func persingkatFN(wg *sync.WaitGroup, input string, results *[]string) {
	defer wg.Done()

	input = strings.ReplaceAll(input, "-", " ")
	inputs := strings.Split(input, " ")

	var result string
	for _, val := range inputs {
		result += strings.ToUpper(string(val[0]))
	}

	*results = append(*results, fmt.Sprintf("%s | %s", input, result))
}

func prompt(question string) (string, error) {
	fmt.Print(question)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
