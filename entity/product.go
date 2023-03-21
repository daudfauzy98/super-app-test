package entity

type Product struct {
	ID       string  `gorm:"size:36;not null;primaryKey" json:"id"`
	Name     string  `gorm:"size:100;not null" json:"name" binding:"required"`
	Category string  `gorm:"size:50;not null" json:"category" binding:"required"`
	Brand    string  `gorm:"size:50; not null" json:"brand" binding:"required"`
	Stocks   []Stock `json:"stocks" binding:"required"`
}
