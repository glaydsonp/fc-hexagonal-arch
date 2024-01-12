package main

import (
	"database/sql"

	db2 "github.com/glaydsonp/go-hexagonal/adapters/db"
	"github.com/glaydsonp/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Test", 10.0)
	productService.Enable(product)
}
