package models

import (
	"errors"
	"log"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Name       string  `json:"name"`
	Adress     string  `json:"adress"`
	Rent       float32 `json:"rent"`
	Charges    float32 `json:"charges"`
	Tenant     Tenant  `json:"tenant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PropertyRefer;"`
	OwnerRefer uint    `json:"owner"`
}

func NewProperty(ownerId int, property *Property) *Property {
	if property == nil {
		log.Fatal(property)
	}
	property.OwnerRefer = uint(ownerId)

	config.GormDb().Create(&property)

	return property
}

func FindPropertyById(ownerId int, id int) *Property {
	var property Property

	result := config.GormDb().Preload("Tenant").First(&property, "owner_refer = ? AND id = ?", ownerId, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &property
}

func AllProperties(ownerId int) *[]Property {
	var propertys []Property

	config.GormDb().Preload("Tenant").Where("owner_refer = ?", ownerId).Find(&propertys)
	return &propertys
}

func UpdateProperty(property *Property) {
	config.GormDb().Session(&gorm.Session{FullSaveAssociations: true}).Save(&property)
}

func DeletePropertyById(property *Property) *Property {
	result := config.GormDb().Select("Tenant").Delete(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return property
}
