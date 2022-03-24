package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Id       uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Link     string
	DeepLink string
	Type     int
}
