package models

import (
	"errors"
	"log"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Adress    string `json:"adress"`
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

	result := config.GormDb().First(&owner, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &owner
}

func AllOwners() *[]Owner {
	var owners []Owner

	config.GormDb().Find(&owners)
	return &owners
}

func UpdateOwner(owner *Owner) {
	config.GormDb().Model(&owner).Updates(owner)
}

func DeleteOwnerById(id int) *Owner {
	var owner Owner
	result := config.GormDb().Delete(&owner, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &owner
}
