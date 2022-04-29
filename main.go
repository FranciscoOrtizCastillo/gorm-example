// Package main provide CRUD Product with GORM.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FranciscoOrtizCastillo/gorm-example/connection"
	"github.com/FranciscoOrtizCastillo/gorm-example/models"
	"gorm.io/gorm"
)

func errorFatal(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func CreateProduct(connn *gorm.DB) {
	/*
		product := models.Product{
			Id: "002",
			Name: "Product 2",
			Price: 12.34,
		}

		rs := connn.Create(&product)
	*/

	p2 := models.Product{
		Id:    "002",
		Name:  "Product 2",
		Price: 12.34,
	}

	p3 := models.Product{
		Id:    "003",
		Name:  "Product 3",
		Price: 20.99,
	}

	p4 := models.Product{
		Id:    "004",
		Name:  "Product 4",
		Price: 9.99,
	}

	rs := connn.Create([]models.Product{p2, p3, p4})

	if rs.Error != nil {
		fmt.Println(rs.Error)
	} else {
		fmt.Println("Product created")
	}
}

func ReadProducts(connn *gorm.DB) {
	var productos []models.Product
	if err := connn.Find(&productos).Error; err != nil {
		log.Println(err.Error())
	}

	for _, producto := range productos {
		fmt.Println(producto)
	}
}

func ReadProductById(connn *gorm.DB, id string) {
	var producto models.Product
	//if err := connn.Where("id = ?", id).Limit(1).Find(&producto).Error; err != nil {
	rs := connn.Where("id = ?", id).First(&producto)

	if rs.Error != nil {
		log.Println(rs.Error.Error())
		return
	}

	if rs.RowsAffected == 0 {
		fmt.Println("Product not found")
		return
	}

	fmt.Println(producto)
}

func UpdateProduct(connn *gorm.DB, id string) {
	var producto = models.Product{Price: 99.99}
	rs := connn.Where("id = ?", id).Updates(&producto)

	if rs.Error != nil {
		log.Println(rs.Error)
		return
	}

	fmt.Println("Product updated")
}

func DeleteProduct(connn *gorm.DB, id string) {
	rs := connn.Where("id = ?", id).Delete(&models.Product{})

	if rs.Error != nil {
		log.Println(rs.Error)
		return
	}

	fmt.Println("Product deleted")
}

func main() {
	//fmt.Println("Initializing DB connection")

	db_host := os.Getenv("DATABASE_HOST")
	db_port := os.Getenv("DATABASE_PORT")
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")

	conn, err := connection.GetConnection(db_host, db_user, db_pass, db_name, db_port)

	errorFatal(err)

	_ = conn
	fmt.Println("Connection DB successful")

	CreateProduct(conn)

	ReadProducts(conn)

	UpdateProduct(conn, "002")

	ReadProductById(conn, "002")

	DeleteProduct(conn, "002")
}
