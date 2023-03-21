package db

import (
	"log"
	"stock-management/entity"

	"github.com/google/uuid"
)

func Seeder() {
	log.Println("[MESSAGE] start seeding Product table..")
	prodID := uuid.NewString()
	products := []entity.Product{
		{
			ID:       prodID,
			Name:     "Head&Shoulder Cooling Menthol",
			Category: "Shampoo",
			Brand:    "Head&Shoulder",
		},
	}
	for _, val := range products {
		err := MainDB.Debug().Create(val).Error
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("[MESSAGE] finish seeding Product table..")

	log.Println("[MESSAGE] start seeding Stock table..")
	stocks := []entity.Stock{
		{
			ID:           uuid.NewString(),
			ProductID:    prodID,
			WarehouseID:  uuid.NewString(),
			WarehouseLoc: "Ngawi",
			Count:        10,
		},
		{
			ID:           uuid.NewString(),
			ProductID:    prodID,
			WarehouseID:  uuid.NewString(),
			WarehouseLoc: "Malang",
			Count:        15,
		},
	}
	for _, val := range stocks {
		err := MainDB.Debug().Create(val).Error
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("[MESSAGE] finish Product table..")
}
