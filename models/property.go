package models

import (
	"errors"
	"log"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Name       string `json:"name"`
	Adress     string `json:"adress"`
	Tenant     Tenant `json:"tenant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PropertyRefer;"`
	OwnerRefer uint   `json:"owner"`
}

func NewProperty(property *Property) *Property {
	if property == nil {
		log.Fatal(property)
	}

	config.GormDb().Create(&property)

	return property
}

func FindPropertyById(id int) *Property {
	var property Property

	result := config.GormDb().First(&property, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &property
}

func AllPropertys() *[]Property {
	var propertys []Property

	config.GormDb().Find(&propertys)
	return &propertys
}

func UpdateProperty(property *Property) {
	config.GormDb().Session(&gorm.Session{FullSaveAssociations: true}).Save(&property)
}

func DeletePropertyById(property *Property) *Property {
	result := config.GormDb().Delete(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return property
}
