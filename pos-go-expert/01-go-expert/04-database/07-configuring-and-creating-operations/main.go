package main

import (
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
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1799.99,
	})

	// create batch
	products := []Product{
		{Name: "Monitor", Price: 1399.99},
		{Name: "Mouse", Price: 299.99},
		{Name: "Keyboard", Price: 599.99},
	}
	db.Create(&products)
}
