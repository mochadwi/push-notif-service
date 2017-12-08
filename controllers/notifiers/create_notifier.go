package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Person struct represent person model.
type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var db *gorm.DB
var err error

// RegisterNotifier will push a message to single client
func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// migrate scheme to Person struct
	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/list-notifiers", GetProjects)

	r.Run(":8080")

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

func GetProjects(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}
