package main

import (
	"fmt"
	"strings"
)

func checkAnagram(source, target string) string {
	if len(source) != len(target) {
		return fmt.Sprintf("%s & %s bukan merupakan anagram", source, target)
	}

	sourceMap := make(map[rune]int)
	targetMap := make(map[rune]int)

	for _, letter := range source {
		sourceMap[letter]++
	}

	for _, letter := range target {
		targetMap[letter]++
	}

	for letter, sourceCount := range sourceMap {
		if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
			return fmt.Sprintf("%s & %s bukan merupakan anagram", source, target)
		}
	}

	return fmt.Sprintf("%s & %s merupakan anagram", source, target)
}

func validate(inputs ...string) error {
	if len(inputs) < 1 {
		panic("validate fn: function have zero parameter")
	}

	for _, input := range inputs {
		if len(input) < 10 {
			return errValidation(fmt.Sprintf("(%s) must be at least 10 character long", input))
		}

		if strings.ContainsAny(input, "!@#$%^&*()-+[]{};:'\"<>?/.,_\t\r 123456789") {
			return errValidation(fmt.Sprintf("(%s) must be not contain any withespace, number or symbol", input))
		}
	}

	return nil
}
