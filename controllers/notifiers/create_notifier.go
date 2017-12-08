package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Person struct represent person model.
type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// RegisterNotifier will push a message to single client
func main() {
	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()

	// migrate scheme to Person struct
	db.AutoMigrate(&Person{})

	p1 := Person{FirstName: "Iqbal", LastName: "Dwi"}
	p2 := Person{FirstName: "Mochamad", LastName: "Cahyo"}

	// insert value from p1 & p2 person
	db.Create(&p1)
	db.Create(&p2)

	var p3 Person

	// find the first row on the table
	db.First(&p3)

	fmt.Print("p1: ")
	fmt.Println(p1)
	fmt.Print("p2: ")
	fmt.Println(p2)
	fmt.Print("p3: ")
	fmt.Println(p3)
}
