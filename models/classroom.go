package models

type ClassRoom struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Class  string `json:"class"`
}