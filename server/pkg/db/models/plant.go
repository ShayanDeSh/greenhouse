package models

import (
	"github.com/jinzhu/gorm"
)

type Plant struct {
	gorm.Model
	Ip         string `gorm:"size:15;unique;not null"`
	Humidity   int
	NeedWater  bool
	Threshhold int
}

func AddPlant(threshhold int, needwater bool, humidity int, ip string, db *gorm.DB) {
	var plant = Plant{
		Ip:         ip,
		Humidity:   humidity,
		NeedWater:  needwater,
		Threshhold: threshhold,
	}
	db.Create(&plant)
}

func DeletePlant(ID int, db *gorm.DB) (err error) {
	err = db.Where("id = ?", ID).Delete(&Plant{}).Error
	return err
}

func UpdatePlant(plant Plant, db *gorm.DB) error {
	err := db.Save(&plant).Error
	return err
}
