package entity

type Product struct {
	ID       string  `gorm:"size:36;not null;primaryKey" json:"id"`
	Name     string  `gorm:"size:100" json:"name"`
	Category string  `gorm:"size:50" json:"category"`
	Brand    string  `gorm:"size:50" json:"brand"`
	Stocks   []Stock `json:"stocks"`
}
