package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	// challenge_1()
	// challenge_2()
	// challenge_3()
	// challenge_4()
	// challenge_5()
	// challenge_6()
	// challenge_7()
	// challenge_8()
	// challenge_9()
}

/*
Buatlah sebuah variable yang merupakan slice yang berisi data map orang-orang dengan data sebagai berikut :
  - name: Hank 		Age: 50 	Job: Polisi
  - name: Heisenberg Age: 52 	Job: Ilmuwan
  - name: Skyler 	Age: 48 	Job: Akuntan

Loop kedalam slice tersebut dan tampilkan tampilkan data-data tiap orang dengan format tulisan :
```Hi Perkenalkan, Nama saya [Nama], umur saya [umur], dan saya bekerja sebagai [pekerjaan]
*/
func challenge_1() {
	slices := []map[string]any{
		{"name": "Hank", "age": 50, "job": "Polisi"},
		{"name": "Heisenberg", "age": 52, "job": "Ilmuwan"},
		{"name": "Skyler", "age": 48, "job": "Akuntan"},
	}

	for _, m := range slices {
		nama := m["nama"]
		umur := m["age"]
		job := m["job"]

		fmt.Printf("Hi Perkenalkan, Nama saya %v, umur saya %v, dan saya bekerja sebagai %v", nama, umur, job)
	}
}

/*
Buatlah sebuah `variable` yang berisikan `slice` dari `float64`,
dan hitunglah `rata-rata`, `jumlah`, dan `median` dari slice tersebut.
Gunakan for dan operasi aritmatika
*/
func challenge_2() {
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	process := func(inputs []float64) (average float64, sum float64, median float64) {

		for _, val := range inputs {
			sum += val
		}
		average = sum / float64(len(inputs))
		median = float64(len(inputs) / 2)

		return average, sum, median
	}

	avg, sum, med := process(slice1)
	fmt.Printf("avg: %.2f \t sum: %.2f \t med: %.2f\n", avg, sum, med)

	avg, sum, med = process(slice2)
	fmt.Printf("avg: %.2f \t sum: %.2f \t med: %.2f\n", avg, sum, med)
}

/*
Diberikan sebuah `variabel` kata bertipe `String`.
Buatlah program yang menampilkan `true` jika kata yang diberikan merupakan `palindrome`
dan false jika bukan.
*/
func challenge_3() {
	process := func(input string) bool {
		length := len(input)

		var reverseInput string
		for i := length - 1; i >= 0; i-- {
			reverseInput += string(input[i])
		}
		return reverseInput == input
	}

	word, err := getinputword("Input word: ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v is palindrome: %v\n", word, process(word))
}

/*
Diberikan sebuah variable kata bertipe String.

Buatlah program dimana program tersebut akan menghitung jumlah karakter `x` dan jumlah karakter `o`.
Setelah perhitungan selesai maka tampilkan `true` jika jumlah karakter `o` dan `x` sama dan `false` jika tidak.
*/
func challenge_4() {
	word, err := getinputword("Input word: ")
	if err != nil {
		log.Fatal(err)
	}

	process := func(input string) bool {
		x_count := 0
		o_count := 0
		for _, val := range word {
			huruf := strings.ToLower(string(val))

			if huruf == "x" {
				x_count++
				continue
			}

			if huruf == "o" {
				o_count++
				continue
			}
		}
		return x_count == o_count
	}

	fmt.Printf("num of 'x' and 'o' in '%s' is same?: %v\n", word, process(word))
}

/*
	Bubble sort
*/

func challenge_5() {
	line, err := getinputline("Input angka: ")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(line, " ")
	sliceInt, err := mapfn(lines, func(param string) (int, error) {
		i, err := strconv.Atoi(param)
		if err != nil {
			return 0, err
		}
		return i, nil
	})
	if err != nil {
		log.Fatalf("error converting str to int: %v", err)
	}
	fmt.Printf("array before: %v\n", sliceInt)

	process := func(arr []int, length int) {
		for i := 0; i < length-1; i++ {
			for j := 0; j < length-i-1; j++ {
				//swap
				if arr[j] > arr[j+1] {
					var temp = arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			}
		}
	}

	process(sliceInt, len(sliceInt))
	fmt.Printf("array after sort: %v\n", sliceInt)
}

/*
Pada tugas ini kamu diminta untuk melakukan looping untuk menampilkan di

	 console barisan asterisks (bintang). Setiap baris memiliki 1 simbol '*'.

	 rows1 // input the number of rows
	 // do loops to display asterisks in the console.
	 ```
	 Output // Jika rows = 5
	 	*
		*
		*
		*
		*

```
*/
func challenge_6() {
	input, err := getinputword("jumlah row: ")
	if err != nil {
		log.Fatal(err)
	}
	rowint, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("error converting to number:", err)
	}
	for i := 0; i < rowint; i++ {
		fmt.Println("*")
	}
}

/*
Pada tugas ini kamu diminta untuk melakukan looping untuk menampilkan di
console barisan asterisks. Setiap baris memiliki jumlah simbol '*' sesuai
dengan jumlah baris. Manfaatkan nested loop untuk menyelesaikan kasus
ini.

	  rows2 // input the number of rows
	  // do loops to display asterisks in the console.
	  ```Output jika
	  	 rows and column = 5

		 *****
		 *****
		 *****
		 *****
		 *****
	  ```
*/
func challenge_7() {
	input, err := getinputword("jumlah row: ")
	if err != nil {
		log.Fatal(err)
	}
	rowint, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("error converting to number:", err)
	}
	for i := 0; i < rowint; i++ {
		for i := 0; i < rowint; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Menyusun tangga bintang.
*/
func challenge_8() {
	input, err := getinputword("jumlah row: ")
	if err != nil {
		log.Fatal(err)
	}
	rowint, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("error converting to number:", err)
	}
	for i := 0; i < rowint; i++ {
		fmt.Println(strings.Repeat("*", i+1))
	}
}

/*
Menyusun tangga bintang terbalik.
*/
func challenge_9() {
	input, err := getinputword("jumlah row: ")
	if err != nil {
		log.Fatal(err)
	}
	rowint, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("error converting to number:", err)
	}

	for i := rowint; i > 0; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}
