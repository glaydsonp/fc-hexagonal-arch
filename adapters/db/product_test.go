package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/glaydsonp/go-hexagonal/adapters/db"
	"github.com/glaydsonp/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	if Db == nil {
		log.Fatal("Failed to create db connection")
	}
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	createTableSql := `create table products (
		"id" string,
		"name" string,
		"status" string,
		"price" float
	);`

	stmt, err := db.Prepare(createTableSql)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insertProductSql := `insert into products values("abc", "Product Test", "disabled", 0.0);`

	stmt, err := db.Prepare(insertProductSql)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
	require.Equal(t, 0.0, product.GetPrice())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 10.0

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, productResult.GetName(), product.GetName())
	require.Equal(t, productResult.GetStatus(), product.GetStatus())
	require.Equal(t, productResult.GetPrice(), product.GetPrice())
	require.Equal(t, productResult.GetID(), product.GetID())

	product.Name = "Product Test Updated"
	product.Price = 20.0

	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, productResult.GetName(), product.GetName())
	require.Equal(t, productResult.GetStatus(), product.GetStatus())
	require.Equal(t, productResult.GetPrice(), product.GetPrice())
	require.Equal(t, productResult.GetID(), product.GetID())
}
