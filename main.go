package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Product struct {
	Id    int    `json:"id,omitempty" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Price int    `json:"price" gorm:"column:price;"`
}

func (Product) TableName() string {
	return "products"
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBConnectionStr")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	newProduct := Product{Name: "Iphone 13 pro max", Price: 99999}
	db.Create(&newProduct)

	fmt.Println(newProduct)

	var products []Product

	db.Find(&products)

	fmt.Println(products)

	var product Product

	db.Where("id = 1").First(&product)

	fmt.Println(product)
}
