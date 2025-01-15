package entity

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name string `valid:"required~Name is required"`
	Price float64 `valid:"required~Price is required"`
	SKU string `valid:"required~SKU is required"`
}