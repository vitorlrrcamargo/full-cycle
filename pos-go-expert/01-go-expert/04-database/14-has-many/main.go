package main

import (
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
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

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

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println("- Name:", product.Name)
		}
	}
}
