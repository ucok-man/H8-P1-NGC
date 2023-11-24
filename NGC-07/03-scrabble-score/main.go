package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var input string
	fmt.Print("enter word: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	var finalscore int
	for _, letter := range input {
		wg.Add(1)
		go searhscore(&wg, string(letter), &finalscore)
	}

	wg.Wait()
	fmt.Println(input, "|", "scrable score:", finalscore)
}

func searhscore(wg *sync.WaitGroup, input string, finascore *int) {
	defer wg.Done()

	input = strings.ToLower(input)
	for score, list := range scrable {
		for _, letter := range list {
			if letter == input {
				*finascore += score
				return
			}
		}
	}
}

var scrable = map[int][]string{
	1:  {"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"},
	2:  {"d", "g"},
	3:  {"b", "c", "m", "p"},
	4:  {"f", "h", "v", "w", "y"},
	5:  {"k"},
	8:  {"j", "x"},
	10: {"q", "z"},
}
