package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string
	CategoryId   int
}