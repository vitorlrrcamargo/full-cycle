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

	// updating
	var product Product
	db.First(&product, 1)
	product.Name = "New Notebook"
	db.Save(&product)

	var product2 Product
	db.First(&product2, 1)
	fmt.Println("\nproduct2:")
	fmt.Println(product2)

	// deleting
	var product3 Product
	db.First(&product3, 2)
	db.Delete(&product3)
}
