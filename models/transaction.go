package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Pay int `json:"pay"`
	Balance int `json:"balance"`
	StudentId uint `json:"student_id"`
	FinanceId uint `json:"finance_id"`
	Quantity uint `json:"quantity"`
}