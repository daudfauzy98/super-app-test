package entity

type Stock struct {
	ID           string `gorm:"size:36;not null;uniqueIndex;primaryKey" json:"id"`
	ProductID    string `gorm:"size:36;not null" json:"product_id"`
	WarehouseID  string `gorm:"size:36;not null" json:"warehouse_id"`
	WarehouseLoc string `gorm:"size:100" json:"warehouse_loc"`
	Count        int    `gorm:"size:10" json:"count"`
}
