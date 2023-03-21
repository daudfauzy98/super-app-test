package main

import (
	"stock-management/db"
	"stock-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.InitDB()
	db.DropTables()
	db.Migration()
	db.Seeder()

	go stockOrder(r)
	go stockAlert(30)

	routes.InitializeProductsRoutes(r)
	routes.InitializeStockRoutes(r)

	r.Run("localhost:8000")
}
