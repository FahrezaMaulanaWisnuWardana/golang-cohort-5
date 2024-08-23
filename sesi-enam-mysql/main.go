package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "root"
	password = ""
	dbname   = "product_go_sql_cohort-5"
)

var (
	db  *sql.DB
	err error
)

type Variants struct {
	ID           int
	Variant_name string
	Quantity     int
	Product_id   int
	Created_at   string
	Updated_at   string
	Product      Products
}
type Products struct {
	ID         int
	Name       string
	Created_at string
	Updated_at string
}

func main() {

	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to database")
	// CreateProduct()
	// UpdateProduct()
	// GetProductById(1)
	// CreateVariant()
	// UpdateVariantById()
	// DeleteVariantById()
	GetProductWithVariant(5)
}
func CreateProduct() {
	var product = Products{}

	dt := time.Now().Format("2006-01-02 15:04:05")
	sqlStatement := `INSERT INTO products (name,created_at,updated_at) values(?,?,?)`
	result, err := db.Exec(sqlStatement, "Product 1", dt, dt)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`
	err = db.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &product.Created_at, &product.Updated_at)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
func UpdateProduct() {
	dt := time.Now().Format("2006-01-02 15:04:05")
	sqlStatement := `UPDATE products SET name=?,updated_at=? WHERE id = ?`
	result, err := db.Exec(sqlStatement, "Product satu update", dt, 1)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
func GetProductById(id int) {
	var product = Products{}
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &product.Created_at, &product.Updated_at)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
func CreateVariant() {
	var variant = Variants{}
	dt := time.Now().Format("2006-01-02 15:04:05")

	sqlStatement := `INSERT INTO variants (variant_name,quantity,product_id,created_at,updated_at) values(?,?,?,?,?)`
	result, err := db.Exec(sqlStatement, "Variant satu", 10, 1, dt, dt)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	sqlRetrieve := `SELECT * FROM variants LEFT JOIN products ON variants.product_id = products.id WHERE variants.id = ?`
	err = db.QueryRow(sqlRetrieve, id).Scan(&variant.ID, &variant.Variant_name, &variant.Quantity, &variant.Product_id, &variant.Created_at, &variant.Updated_at, &variant.Product.ID, &variant.Product.Name, &variant.Product.Created_at, &variant.Product.Updated_at)
	if err != nil {
		panic(err)
	}
	fmt.Println(variant)
}
func UpdateVariantById() {
	dt := time.Now().Format("2006-01-02 15:04:05")
	sqlStatement := `UPDATE variants SET variant_name=?,quantity=?,product_id=?,updated_at=? WHERE id = ?`
	result, err := db.Exec(sqlStatement, "Variant satu update", 10, 1, dt, 1)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
func DeleteVariantById() {
	sqlStatement := `DELETE from variants WHERE id = ?`
	result, err := db.Exec(sqlStatement, 2)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
func GetProductWithVariant(id int) {
	var variant = Variants{}
	sqlRetrieve := `SELECT * FROM variants LEFT JOIN products ON variants.product_id = products.id WHERE variants.id = ?`
	err = db.QueryRow(sqlRetrieve, id).Scan(&variant.ID, &variant.Variant_name, &variant.Quantity, &variant.Product_id, &variant.Created_at, &variant.Updated_at, &variant.Product.ID, &variant.Product.Name, &variant.Product.Created_at, &variant.Product.Updated_at)
	if err != nil {
		panic(err)
	}
	fmt.Println(variant)
}
