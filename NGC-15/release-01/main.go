package main

import (
	"H8-P1-NGC/NGC-15/release-01/db"
	"fmt"
	"log"
)

func main() {
	// create connection
	dtbase, err := db.New("root:@tcp(localhost:3306)/tasty_bytes_go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success establish connection to mysql!")

	// setup table
	err = dtbase.CreateTables()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success setup all table!")
}
