package entity

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name  string
	Price float64 `Validate : "Price must be between 1 and 1000(1|1000)"`
	SKU   string
}
