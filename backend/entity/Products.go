package entity

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name string `valid:"required~Name is required"`
	Price float64 `valid:"required~Price is required, range(1|1000)"`
	SKU string `valid:"required~SKU is required"`
}