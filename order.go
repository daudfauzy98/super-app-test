package main

import (
	"log"
	"stock-management/db"
	"stock-management/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// API: /stock/:id (product_id)
func stockOrder(r *gin.Engine) {
	r.PUT("/order/:id", func(ctx *gin.Context) {
		productID := ctx.Param("id")
		numOrder, err := strconv.Atoi(ctx.Query("order"))
		if err != nil {
			log.Println(err)
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			return
		}

		err = db.MainDB.Transaction(func(tx *gorm.DB) error {
			// Search available stock at the warehouses randomly
			var availableStock *entity.Stock
			err = db.MainDB.Where("product_id = ? AND count >= ?", productID, numOrder).
				First(&availableStock).Error
			if err != nil {
				return err
			}

			condition := "product_id = ? AND warehouse_id = ?"
			err = tx.Model(&entity.Stock{}).Where(condition, productID, availableStock.WarehouseID).
				Update("count", availableStock.Count-numOrder).Error
			if err != nil {
				return err
			}
			return nil
		})
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
	})
}
