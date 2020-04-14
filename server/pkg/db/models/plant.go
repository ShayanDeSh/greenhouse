package models

import (
	"github.com/jinzhu/gorm"
)

type Plant struct {
	gorm.Model
	Ip         string `gorm:"size:15;unique;not null" json:"Ip"`
	Humidity   int    `json:"Humidity"`
	NeedWater  bool   `json:"NeedWater"`
	Threshhold int    `json:"Threshhold"`
}

type Plants struct {
	db *gorm.DB
}

func NewPlants(db *gorm.DB) *Plants {
	return &Plants{db}
}

func (p *Plants) Add(plant *Plant) (err error) {
	err = p.db.Create(plant).Error
	return err
}

func (p *Plants) Delete(ID int) (err error) {
	err = p.db.Unscoped().Where("id = ?", ID).Delete(&Plant{}).Error
	return err
}

func (p *Plants) Update(plant *Plant) error {
	err := p.db.Model(&plant).Updates(&plant).Error
	return err
}
