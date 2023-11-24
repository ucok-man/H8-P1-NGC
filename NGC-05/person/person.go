package person

import (
	"fmt"
	"reflect"
)

/*
	Buatlah sebuah struct yang `Person` yang mempunyai properti `Name string`,
	`Age int`, dan `Job int` lalu Buatlah sebuah method `GetInfo` yang akan
	menampilkan/print informasi dengan format sebagai berikut :

	```bash
	Name: Bambang
	Age: 20
	Job: Gambler
	```

	Buatlah lagi sebuah method `AddYear` yang akan menambahkan `Age` sebesar
	`1`, dan jika setelah penjumlahan `Age` bernilai >= 50 maka secara
	otomatis `Job` akan berubah menjadi `Retired`

	Buatlah studi kasus pada `Person` dengan `Age` awal bernilai 45 dan dengan menjalankan `AddYear` sebanyak 5x maka seharusnya value `Age` akan menjadi 50 dan `Job` akan menjadi `Retired`

	* Pastikan ketika `AddYear` mengubah struct aslinya, bukan membuat copy baru. Silahkan explore dari penggunaan Pointer pada golang
*/

type Person struct {
	Name string
	Age  int
	Job  string
}

func (p Person) GetInfo() string {
	var format string
	rv := reflect.ValueOf(&p)
	rt := rv.Elem().Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := reflect.Indirect(rv).FieldByName(field.Name)
		format += fmt.Sprintf("%v\t\t: %v\n", field.Name, value)
	}
	return format
}

func (p *Person) AddYear(sum ...int) {
	var addAge int
	if len(sum) > 0 {
		for _, val := range sum {
			addAge += val
		}
	} else {
		addAge = 1
	}

	p.Age += addAge
	if p.Age >= 50 {
		p.Job = "Retired"
	}
}

func Challenge_1() {
	bambang := Person{
		Name: "Bambang",
		Age:  20,
		Job:  "Gambler",
	}
	fmt.Println(bambang.GetInfo())

	bengbeng := Person{
		Name: "Beng Beng",
		Age:  45,
		Job:  "Gambler",
	}
	fmt.Println(bengbeng.GetInfo())

	fmt.Println("beng beng add 5 year...")
	bengbeng.AddYear(5)
	fmt.Println(bengbeng.GetInfo())

}

/*
	buatlah sebuah slice of Person, dan lakukan pengulangan terhadap slice tersebut lalu disetiap perulangan tampilkan informasi dengan menggunakan method `GetInfo`
*/

func Challenge_2() {
	persons := []Person{
		{
			Name: "Bambang",
			Age:  20,
			Job:  "Gambler",
		},
		{
			Name: "Beng Beng",
			Age:  45,
			Job:  "Gambler",
		},
		{
			Name: "Bung Bung",
			Age:  10,
			Job:  "Gambler",
		},
	}

	for _, person := range persons {
		fmt.Println(person.GetInfo())
	}
}
