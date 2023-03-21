package controllers

import (
	"log"
	"stock-management/db"
	"stock-management/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API: /stock/:id
func GetStock(ctx *gin.Context) {
	var stocks []entity.Stock
	err := db.MainDB.Find(&stocks, "product_id = ?", ctx.Param("id")).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "success",
		"data":    stocks,
	})
}

// API: /stock/:id (stock ID)
// Req: query param "count"
func AddStock(ctx *gin.Context) {
	stockID := ctx.Param("id")
	var stock *entity.Stock
	err := db.MainDB.Find(&stock, "id = ?", stockID).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	updatedStock, err := strconv.Atoi(ctx.Query("count"))
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	err = db.MainDB.Model(&entity.Stock{}).Where("id = ?", stockID).
		Update("count", stock.Count+updatedStock).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "success",
	})
}
