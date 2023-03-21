package controllers

import (
	"encoding/json"
	"log"
	"stock-management/db"
	"stock-management/entity"

	"github.com/gin-gonic/gin"
)

// API: /products
func GetProducts(ctx *gin.Context) {
	var products []entity.Product
	err := db.MainDB.Preload("Stocks").Order("name asc").Find(&products).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var response []map[string]interface{}
	b, err := json.Marshal(&products)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	for i, val := range products {
		count := 0
		for _, val := range val.Stocks {
			count += val.Count
		}
		if val.ID == response[i]["id"] {
			response[i]["stocks"] = count
		}
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "success",
		"data":    response,
	})
}

// API: /product/:id
func GetProduct(ctx *gin.Context) {
	var product entity.Product
	err := db.MainDB.Preload("Stocks").Find(&product, "id = ?", ctx.Param("id")).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var response map[string]interface{}
	b, err := json.Marshal(&product)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	count := 0
	for _, val := range product.Stocks {
		count += val.Count
	}
	response["stocks"] = count

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "success",
		"data":    response,
	})
}
