package routes

import (
	"stock-management/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeProductsRoutes(r *gin.Engine) {
	r.GET("/product/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "success",
			"message": "This is product API",
		})
	})
	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:id", controllers.GetProduct)
	r.POST("/product", controllers.AddProduct)
}
