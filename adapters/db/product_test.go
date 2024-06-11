package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Gabriel-Macedogmc/hexagonal-architecture/adapters/db"
	"github.com/Gabriel-Macedogmc/hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name" string, "price" float, "status" string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products ("id", "name", "price", "status") VALUES ("abc", "Product A", 0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product A", product.GetName())
	require.Equal(t, float64(0.0), product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	product := application.NewProduct()
	product.Name = "Product A"
	product.Price = float64(10)

	productResult, err := productDB.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = application.ENABLED

	productResult, err = productDB.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Status, productResult.GetStatus())
}
