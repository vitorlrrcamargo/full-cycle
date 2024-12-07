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

	// select with where
	var products []Product
	db.Limit(2).Offset(2).Find(&products)
	fmt.Println("\nproducts:")
	for _, product := range products {
		fmt.Println(product)
	}

	var products2 []Product
	db.Where("price > ?", 1000).Find(&products2)
	fmt.Println("\nproducts2:")
	for _, product := range products2 {
		fmt.Println(product)
	}

	var products3 []Product
	db.Where("name LIKE ?", "Mo%").Find(&products3)
	fmt.Println("\nproducts3:")
	for _, product := range products3 {
		fmt.Println(product)
	}
}
