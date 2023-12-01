package main

import (
	"H8-P1-NGC/NGC-15/release-02/db"
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

	// drop table
	err = dtbase.DropAllTable()
	if err != nil {
		log.Fatal(err)
	}

	// setup table
	err = dtbase.CreateTables()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success setup all table!")

	// inserting employee table
	fmt.Println("inserting employe table...")
	employees := []db.Employee{
		{FirstName: "John", LastName: "Doe", Position: "Manager"},
		{FirstName: "Jane", LastName: "Smith", Position: "Waiters"},
		{FirstName: "Robert", LastName: "Brown", Position: "Cook"},
	}
	for _, emp := range employees {
		newemp, err := dtbase.Employee.Insert(&emp)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("\tsuccess:", newemp)
	}

	// inserting menu items
	fmt.Println("inserting menu_items table...")
	menuItems := []db.MenuItem{
		{Name: "Spaghetti Carbonara",
			Description: "Traditional italian dish with eggs, cheese.",
			Price:       12.50,
			Category:    "Main Course",
		},
		{Name: "Caesar Salad",
			Description: "Fresh lectucce with caesar dressing.",
			Price:       6.00,
			Category:    "Starter",
		},
		{Name: "Tiramisu",
			Description: "Classic italian dessert coy!!",
			Price:       12.50,
			Category:    "Dessert",
		},
	}
	for _, item := range menuItems {
		newitem, err := dtbase.MenuItem.Insert(&item)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("\tsuccess:", newitem)
	}
}
