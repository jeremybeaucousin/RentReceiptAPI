package models

import (
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

	config.GormDb().First(&owner, "id = ?", id)
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

func DeleteOwnerById(id int) {
	var owner Owner
	config.GormDb().Delete(&owner, id)
}
