package models

import "gorm.io/gorm"

type Parent struct {
	gorm.Model
	PrefixId uint `json:"prefix_id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}