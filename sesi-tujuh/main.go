package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	ID            int     `json:"id"`
	Customer_name string  `json:"customerName"`
	Ordered_at    string  `json:"orderedAt"`
	Items         []Items `json:"items"`
	Created_at    string  `json:"createdAt"`
	Updated_at    string  `json:"updatedAt"`
}
type Items struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_ID    int    `json:"orderId"`
	Created_at  string `json:"createdAt"`
	Updated_at  string `json:"updatedAt"`
}

const (
	user     = "root"
	password = ""
	dbname   = "order_assignment"
)

var (
	db  *sql.DB
	err error
)

var PORT = ":9090"
var TimeFormat = "2006-01-02 15:04:05"

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}
	// Route

	r := gin.Default()

	r.POST("/orders", createOrder)       //Create Order
	r.GET("/orders", GetOrders)          // Get All Order
	r.GET("/orders/:id", GetOrderByID)   // Get Order By ID
	r.PUT("/orders/:id", UpdateOrder)    // Update Order
	r.DELETE("/orders/:id", DeleteOrder) // Delete Order

	// Set Server
	r.Run(PORT)

}

func createOrder(c *gin.Context) {
	var newOrder Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	dt := time.Now().Format(TimeFormat)
	sqlStmtOrder := `INSERT INTO orders (customer_name,ordered_at,created_at,updated_at) values(?,?,?,?)`

	orderAt, _ := time.Parse(time.RFC3339, newOrder.Ordered_at)
	result, err := db.Exec(sqlStmtOrder, newOrder.Customer_name, orderAt, dt, dt)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	sqlStmtItem := `INSERT INTO items (name,description,quantity,order_id,created_at,updated_at) values(?,?,?,?,?,?)`
	for _, val := range newOrder.Items {
		_, err := db.Exec(sqlStmtItem, val.Name, val.Description, val.Quantity, id, dt, dt)
		if err != nil {
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Berhasil Tambah Data",
	})
}

func GetOrders(c *gin.Context) {

	sqlStatement := `select * from orders INNER JOIN items ON orders.id = items.order_id`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	ordersMap := make(map[int]*Order)
	for rows.Next() {
		var (
			orderID         int
			customerName    string
			orderedAt       string
			orderCreatedAt  string
			orderUpdatedAt  string
			itemID          sql.NullInt64
			itemName        sql.NullString
			itemDescription sql.NullString
			itemQuantity    sql.NullInt64
			itemOrderID     sql.NullInt64
			itemCreatedAt   sql.NullString
			itemUpdatedAt   sql.NullString
		)

		if err := rows.Scan(&orderID, &customerName, &orderedAt, &orderCreatedAt, &orderUpdatedAt,
			&itemID, &itemName, &itemDescription, &itemQuantity, &itemOrderID,
			&itemCreatedAt, &itemUpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		order, exists := ordersMap[orderID]
		if !exists {
			order = &Order{
				ID:            orderID,
				Customer_name: customerName,
				Ordered_at:    orderedAt,
				Created_at:    orderCreatedAt,
				Updated_at:    orderUpdatedAt,
			}
			ordersMap[orderID] = order
		}

		if itemID.Valid {
			item := Items{
				ID:          int(itemID.Int64),
				Name:        itemName.String,
				Description: itemDescription.String,
				Quantity:    int(itemQuantity.Int64),
				Order_ID:    int(itemOrderID.Int64),
				Created_at:  itemCreatedAt.String,
				Updated_at:  itemUpdatedAt.String,
			}
			order.Items = append(order.Items, item)
		}
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var orders []Order
	for _, order := range ordersMap {
		orders = append(orders, *order)
	}

	c.JSON(http.StatusOK, orders)
}
func GetOrderByID(c *gin.Context) {

	orderIDStr := c.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id"})
		return
	}
	orderQuery := `SELECT id, customer_name, ordered_at, created_at, updated_at FROM orders WHERE id = ?`
	var order Order
	err = db.QueryRow(orderQuery, orderID).Scan(
		&order.ID,
		&order.Customer_name,
		&order.Ordered_at,
		&order.Created_at,
		&order.Updated_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Query to get items for the order
	itemsQuery := `SELECT id, name, description, quantity, order_id, created_at, updated_at FROM items WHERE order_id = ?`
	rows, err := db.Query(itemsQuery, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []Items
	for rows.Next() {
		var item Items
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Quantity,
			&item.Order_ID,
			&item.Created_at,
			&item.Updated_at,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}
	order.Items = items
	c.JSON(http.StatusOK, order)
}
func UpdateOrder(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	dt := time.Now().Format(TimeFormat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}
	var input Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	orderAt, _ := time.Parse(time.RFC3339, input.Ordered_at)
	_, err = db.Exec(`UPDATE orders SET customer_name = ?, ordered_at = ?, updated_at = ? WHERE id = ? `, input.Customer_name, orderAt, dt, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec(`DELETE FROM items WHERE order_id = ? `, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sqlStmtItem := `INSERT INTO items (name,description,quantity,order_id,created_at,updated_at) values(?,?,?,?,?,?)`
	for _, val := range input.Items {
		_, err := db.Exec(sqlStmtItem, val.Name, val.Description, val.Quantity, orderID, dt, dt)
		if err != nil {
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order and items updated successfully"})
}
func DeleteOrder(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}
	_, err = db.Exec(`DELETE FROM orders WHERE id = ? `, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec(`DELETE FROM items WHERE order_id = ? `, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order and items deleted successfully"})
}
