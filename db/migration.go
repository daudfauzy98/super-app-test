package db

import (
	"log"
	"stock-management/entity"
)

func Migration() {
	log.Println("[MESSAGE] start migrating DB tables..")

	entities := []interface{}{
		entity.Product{},
		entity.Stock{},
	}

	err := MainDB.Debug().AutoMigrate(entities...)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("[MESSAGE] migrating DB done!")
}

func DropTables() {
	log.Println("[MESSAGE] start droping all tables..")
	entities := []interface{}{
		entity.Product{},
		entity.Stock{},
	}
	err := MainDB.Debug().Migrator().DropTable(entities...)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("[MESSAGE] droping all tables done!")
}
