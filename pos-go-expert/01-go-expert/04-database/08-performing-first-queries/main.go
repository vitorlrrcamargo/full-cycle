package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// select one
	var product Product
	db.First(&product, 2)
	fmt.Println("product:")
	fmt.Println(product)

	var product2 Product
	db.First(&product2, "name = ?", "Mouse")
	fmt.Println("\nproduct2:")
	fmt.Println(product2)

	// select all
	var products []Product
	db.Find(&products)
	fmt.Println("\nproducts:")
	for _, product := range products {
		fmt.Println(product)
	}
}
