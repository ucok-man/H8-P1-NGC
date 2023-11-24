package main

import (
	"fmt"
	"strings"
)

func main() {
	// challenge_1()
	// challenge_2()
	// challenge_3()
	// challenge_4()
	// challenge_5()
}

func sparator(name string) {
	n := 80
	if name == "" {
		println(strings.Repeat("=", n), "\n")
		return
	}

	n = n - len(name)
	fmt.Print(strings.Repeat("=", (n/2)-1), " ")
	fmt.Print(name)

	var m int
	if n%2 != 0 {
		m = (n / 2) + 1
	} else {
		m = n / 2
	}
	fmt.Print(" ", strings.Repeat("=", m-1), "\n")
}

/*
A.Buatlah sebuah variabel dengan nama `myNum` dengan tipe data int32 dan assign nilai 50 ke dalam variabel tersebut Lalu, print variable tersebut ke dalam terminal

# B.Buatlah sebuah variabel dengan nama `myNum2` dengan tipe data float32 dan assign nilai 51 ke dalam variabel tersebut Lalu, print variable tersebut ke dalam terminal

C.Buatlah sebuah variabel dengan nama `myNumStr` dengan tipe data string dan assign karakter "50" ke dalam variabel tersebut Lalu, print variable tersebut ke dalam terminal
*/
func challenge_1() {
	sparator("Challenge 1")

	var mynum int32 = 50
	var myfloat float64 = 51
	var mystr string = "50"

	fmt.Printf("mynum: %v \t myfloat: %.2f \t mystr: %v\n", mynum, myfloat, mystr)

	sparator("")
}

/*
Buatlah variabel dengan dengan ketentuan berikut:
- x dan assign nilai int32 5
- y dan assign nilai int32 10
Buatlah variabel baru dengan nama z, yang merupakan hasil penjumlahan dari x dan y. Lalu print ke dalam terminal
*/
func challenge_2() {
	sparator("Challenge 2")

	var x int32 = 5
	var y int32 = 10
	var z = x + y
	fmt.Println("x + y : ", z)

	sparator("")
}

/*
Membuat program CLI yang menerima input nama (string) dan akan menampilkan `hello [nama]`
*/
func challenge_3() {
	sparator("Challenge 3")

	var nama string
	fmt.Print("masukan nama : ")
	fmt.Scan(&nama)
	fmt.Println("Hello", nama)

	sparator("")
}

/*
Buatlah slice dengan nama `people` yang berisikan nama-nama orang `Walt, Jesse, Skyler, Saul`
- hitung berapa panjang slice tersebut dan tampilkan ke dalam terminal
- tambahkan `Hank` dan `Marie` ke dalam slice `people`
- hitung berapa panjang slice setelah ditambahkan `Hank` dan `Marie` dan tampilkan ke dalam terminal
- tampilkan `people` ke dalam terminal
*/
func challenge_4() {
	sparator("Challenge 4")

	people := []string{"Walt", "Jesse", "Skyler", "Saul"}
	fmt.Println("slice:", people)
	fmt.Println("Initial length:", len(people))

	people = append(people, "Hank", "Marie")
	fmt.Println("After 'Marie' and 'Hank' length is:", len(people))
	fmt.Println("slice:", people)

	sparator("")
}

/*
Buatlah sebuah array dengan kapasitas 3 elemen yang berisikan `map` yang memiliki keys: `name` dan `gender`.
lalu isi array tersebut dengan data berikut, menggunakan perintah append

  - tampilkan array tersebut ke dalam terminal.

    Setelah itu modifikasi array tadi, untuk setiap data dengan `gender` `F` tambahkan `Mrs` di depan namanya, dan `M` tambahkan `Mr`.

  - tampilkan kembali ke dalam terminal
*/
func challenge_5() {
	sparator("Challenge 5")

	arr := [3]map[string]string{}

	slices := arr[:]
	names := []string{"Hank", "Heisenberg", "Skyler"}
	genders := []string{"M", "M", "F"}

	for idx := range slices {
		// delete and append
		slices = append(slices[1:], map[string]string{
			"name":   names[idx],
			"gender": genders[idx],
		})
	}
	fmt.Println("slices :", slices)

	for _, mapvalue := range slices {
		name := mapvalue["name"]

		if mapvalue["gender"] == "M" {
			mapvalue["name"] = fmt.Sprintf("Mr %s", name)
			continue
		}

		if mapvalue["gender"] == "F" {
			mapvalue["name"] = fmt.Sprintf("Mrs %s", name)
			continue
		}

	}
	fmt.Println("slices :", slices)

	sparator("")
}
