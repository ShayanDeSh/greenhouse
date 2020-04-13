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

type Plants struct {
	db *gorm.DB
}

func NewPlants(db *gorm.DB) *Plants {
	return &Plants{db}
}

func (p *Plants) AddPlant(threshhold int, needwater bool, humidity int, ip string) {
	var plant = Plant{
		Ip:         ip,
		Humidity:   humidity,
		NeedWater:  needwater,
		Threshhold: threshhold,
	}
	p.db.Create(&plant)
}

func (p *Plants) DeletePlant(ID int) (err error) {
	err = p.db.Where("id = ?", ID).Delete(&Plant{}).Error
	return err
}

func (p *Plants) UpdatePlant(plant Plant) error {
	err := p.db.Save(&plant).Error
	return err
}
