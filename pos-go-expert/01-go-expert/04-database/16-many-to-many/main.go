package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create categories
	category := Category{Name: "Kitchen"}
	db.Create(&category)

	category2 := Category{Name: "Eletronics"}
	db.Create(&category2)

	// create products
	db.Create([]Product{
		{Name: "Refrigerator", Price: 1599.99, Categories: []Category{category, category2}},
		{Name: "Induction Cooktop", Price: 899.99, Categories: []Category{category, category2}},
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
