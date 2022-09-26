package models

import _"gorm.io/gorm"

type Prefix struct {
	ID   uint   `gorm:"primaryKey"; json:"id"`
	Name string `json:"name"`
}