package routes

import (
	"stock-management/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeStockRoutes(r *gin.Engine) {
	r.GET("/stock", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "success",
			"message": "This is stock API",
		})
	})
	r.GET("/stock/:id", controllers.GetStock)
	r.PUT("/stock/:id", controllers.AddStock)
}
