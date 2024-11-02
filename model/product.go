package model

type Product struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	SKU          string  `json:"sku" gorm:"unique;not null"`
	Name         string  `json:"name" gorm:"not null"`
	Price        float64 `json:"price" gorm:"not null"`
	InventoryQty int     `json:"inventoryQty" gorm:"not null"`
}
