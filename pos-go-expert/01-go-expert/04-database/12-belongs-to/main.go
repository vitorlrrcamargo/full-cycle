package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	// create products
	db.Create([]Product{
		{Name: "Monitor", Price: 1399.99, CategoryID: category.ID},
		{Name: "Mouse", Price: 299.99, CategoryID: category.ID},
	})

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println("Name:", product.Name, "| Category:", product.Category.Name)
	}
}
