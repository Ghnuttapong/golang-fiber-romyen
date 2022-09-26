package models

import "gorm.io/gorm"

type Finance struct {
	gorm.Model
	Name string `json:"name"`
	Expenses string `json:"expenses"`
	Note string `json:"note "`
	CategoryId uint `json:"category_id"`
}