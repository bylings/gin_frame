package models

import (
	"gorm.io/gorm"
)

type Regions struct {
	gorm.Model
	//Id   int
	Name string
	Code int
}
