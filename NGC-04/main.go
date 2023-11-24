package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	challenge_1()
	// challenge_2()
}

/*
	Buatlah sebuah variadic function yang dapat menerima argumen string yang
	tak terbatas jumlahnya.Function tersebut akan me-return sebuah kalimat
	string, yang merupakan gabungan dari tiap kata yang ada di argumen.

	Contoh:
	```
	AlayGen("hello", "world", "zzz") // akan menghasilkan "hello world zzz"
	```

	selain mengabungkan katanya, kita perlu membuat tiap kata menjadi alay,
	dengan mengubah ```a -> 4e -> 3i -> !l -> 1n -> Ns -> 5x -> *```

	untuk mempermudah pengerjaan gunakan konsep modular function
*/

func challenge_1() {
	inputs := getinputs()
	joinedstr := joinStr(inputs...)
	var result string
	for _, val := range joinedstr {
		result += toAlay(string(val))
	}

	fmt.Println("result:", result)
}

func joinStr(params ...string) string {
	var result string
	for _, val := range params {
		result += " " + val
	}
	return result
}

func getinputs() []string {
	iterationStr, err := getinputword("how many word? ")
	if err != nil {
		log.Fatal(err)
	}
	iteration, err := strconv.Atoi(iterationStr)
	if err != nil {
		log.Fatal("error input must be number:", err)
	}

	var inputs []string
	for i := 0; i < iteration; i++ {
		word, err := getinputword(fmt.Sprintf("word %v: ", i))
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, word)
	}
	return inputs
}

func toAlay(input string) string {
	container := map[string]string{
		"a": "4",
		"e": "3",
		"i": "!",
		"l": "1",
		"n": "N",
		"s": "5",
		"x": "*",
	}
	for key, val := range container {
		if key == input {
			return val
		}
	}
	return input
}

/* ---------------------------------------------------------------- */
/*                         fibonacci squence                        */
/* ---------------------------------------------------------------- */

func challenge_2() {
	input, err := getinputword("index fibonacci squence: ")
	if err != nil {
		log.Fatal(err)
	}
	nth, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("error input must be number:", err)
	}
	fmt.Println(fibonacci(nth))
}

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
