package models

import (
	"errors"
	"log"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	FirstName  string     `json:"firstname"`
	LastName   string     `json:"lastname"`
	Adress     string     `json:"adress"`
	Properties []Property `json:"properties" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OwnerRefer;"`
}

func NewOwner(owner *Owner) *Owner {
	if owner == nil {
		log.Fatal(owner)
	}

	config.GormDb().Create(&owner)

	return owner
}

func FindOwnerById(id int) *Owner {
	var owner Owner

	result := config.GormDb().Preload("Properties").First(&owner, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &owner
}

func AllOwners() *[]Owner {
	var owners []Owner

	config.GormDb().Preload("Properties").Find(&owners)
	return &owners
}

func UpdateOwner(owner *Owner) {
	config.GormDb().Session(&gorm.Session{FullSaveAssociations: true}).Save(&owner)
}

func DeleteOwnerById(owner *Owner) *Owner {
	result := config.GormDb().Select("Properties").Delete(&owner)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return owner
}
