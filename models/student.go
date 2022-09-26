package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Code string `json:"code"`
	PrefixId uint `json:"prefix_id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Nickname string `json:"nickname"`
	ParentId uint `json:"parent_id"`
	ClassroomID uint `json:"classroom_id"`
	RoomNumber uint `json:"room_number"`
	Status string `json:"status"`
	YearAt uint `json:"year_at"`
	Term string `json:"term"`
	UserId uint `json:"user_id"`
}