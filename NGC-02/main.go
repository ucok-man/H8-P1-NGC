package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// challenge_1()
	// challenge_2()
}

/*
Buatlah sebuah program CLI sederhana yang menerima input`Nama`Program tersebut akan me-random angka `1-100` dan akan menampilkan :
  - `Selamat [Name], anda sangat beruntung` jika mendapat angka > 80
  - `Selamat [Name], anda beruntung` jika mendapat angka <= 80 dan > 60
  - `Mohon maaf [Name], anda kurang beruntung` jika mendapat angka <= 60 dan > 40
  - `Mohon maaf [Name], anda sial` jika mendapat angka <= 40
*/
func challenge_1() {
	fmt.Print("Input name: ")

	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		log.Fatalf("error getting username: %v", err)
	}

	random := rand.New(rand.NewSource(time.Now().UnixMilli()))
	randnum := random.Intn(101)

	println("number: ", randnum)
	switch {
	case randnum > 80:
		fmt.Printf("Selamat %v anda sangat beruntung\n", name)
	case randnum <= 80 && randnum > 60:
		fmt.Printf("Selamat %v anda beruntung\n", name)
	case randnum <= 60 && randnum > 40:
		fmt.Printf("Mohon maaf %s, anda kurang beruntung\n", name)
	default:
		fmt.Printf("Mohon maaf %s, anda sial\n", name)
	}
}

/*
Buatlah sebuah aplikasi CLI yang akan menerima input berupa `Nama` dan `Umur`.
Jika`Umur`lebih besar dari 18 maka aplikasi akan menampilkan`Silahkan masuk`dan jika`Umur` lebih kecil sama dengan 18 maka akan menampilkan `Dilarang masuk, maksimal umur 19`.

Dan perlu perlu juga menampilkan pesan error jika:
- `Umur < 0`, `Umur > 100`,
- Umur bukan merupakan angka,
- PESAN ERROR dibebaskan selama masuk akal.
*/
func challenge_2() {
	fmt.Print("Input name: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		log.Fatalf("error: getting user name: %v", err)
	}

	fmt.Print("Input Age: ")
	var age int
	_, err = fmt.Scan(&age)
	if err != nil {
		log.Fatalf("error: getting user age: %v", err)
	}

	if age < 0 && age > 100 {
		fmt.Fprintln(os.Stderr, "error: age must be `0 < age < 100`")
		os.Exit(1)
	}

	switch {
	case age > 18:
		fmt.Println("Silahkan masuk")
	default:
		fmt.Println("Dilarang masuk, maksimal umur 19")
	}

}
