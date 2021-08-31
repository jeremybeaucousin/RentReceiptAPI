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
	Properties []Property `json:"properties" gorm:"foreignKey:OwnerRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	config.GormDb().Model(&owner).Preload("Properties").Updates(owner)
}

func DeleteOwnerById(id int) *Owner {
	var owner Owner
	result := config.GormDb().Preload("Properties").Delete(&owner, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &owner
}
