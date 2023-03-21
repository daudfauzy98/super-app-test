package main

import (
	"log"
	"stock-management/db"
	"stock-management/entity"
	"time"
)

func stockAlert(threshold int) {
	for {
		var stocks []*entity.Stock
		err := db.MainDB.Find(&stocks).Error
		if err != nil {
			log.Println(err)
		}
		for _, val := range stocks {
			if val.Count < threshold {
				msg := "[WARNING] stock of product %s at the ware house %s is below the threshold!\n"
				log.Printf(msg, val.ProductID, val.WarehouseLoc)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
