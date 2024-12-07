package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create categories
	db.Create([]Category{
		{Name: "Electronics"},
		{Name: "Kitchen"},
	})

	// create products
	db.Create([]Product{
		{Name: "Monitor", Price: 1399.99, CategoryID: 1},
		{Name: "Mouse", Price: 299.99, CategoryID: 1},
		{Name: "Keyboard", Price: 599.99, CategoryID: 1},
		{Name: "Refrigerator", Price: 1599.99, CategoryID: 2},
		{Name: "Cooktop", Price: 899.99, CategoryID: 2},
	})

	// create serial numbers
	db.Create([]SerialNumber{
		{Number: "A1B2C3D4E5", ProductID: 1},
		{Number: "B2C3D4E5F6", ProductID: 2},
		{Number: "C3D4E5F6G7", ProductID: 3},
		{Number: "D4E5F6G7H8", ProductID: 4},
		{Number: "E5F6G7H8I9", ProductID: 5},
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println("- Name:", product.Name, "| Serial Number:", product.SerialNumber.Number)
		}
	}
}
