package models

import (
	"errors"
	"log"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Adress        string `json:"adress"`
	PropertyRefer uint   `json:"property"`
}

func NewTenant(propertyId int, tenant *Tenant) *Tenant {
	if tenant == nil {
		log.Fatal(tenant)
	}
	tenant.PropertyRefer = uint(propertyId)
	config.GormDb().Create(&tenant)

	return tenant
}

func FindTenantById(propertyId int, id int) *Tenant {
	var tenant Tenant

	result := config.GormDb().First(&tenant, "property_refer = ? AND id = ?", propertyId, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &tenant
}

func AllTenants(propertyId int) *[]Tenant {
	var tenants []Tenant

	config.GormDb().Where("property_refer = ?", propertyId).Find(&tenants)
	return &tenants
}

func UpdateTenant(tenant *Tenant) {
	config.GormDb().Session(&gorm.Session{FullSaveAssociations: true}).Save(&tenant)
}

func DeleteTenantById(tenant *Tenant) *Tenant {
	result := config.GormDb().Delete(&tenant)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return tenant
}
