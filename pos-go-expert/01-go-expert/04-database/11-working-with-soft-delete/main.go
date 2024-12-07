package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1799.99,
	})

	// create products
	products := []Product{
		{Name: "Monitor", Price: 1399.99},
		{Name: "Mouse", Price: 299.99},
		{Name: "Keyboard", Price: 599.99},
	}
	db.Create(&products)

	// updating
	var product Product
	db.First(&product, 1)
	product.Name = "New Notebook"
	db.Save(&product)

	// deleting
	var product2 Product
	db.First(&product2, 2)
	db.Delete(&product2)
}
