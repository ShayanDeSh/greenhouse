package models

import (
	"github.com/jinzhu/gorm"
)

type Plant struct {
	gorm.Model
	Ip        string `gorm:"size:15;unique;not null"`
	Humidity  int
	NeedWater bool
}
