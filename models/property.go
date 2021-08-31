package models

import (
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Name       string `json:"name"`
	Adress     string `json:"adress"`
	OwnerRefer uint   `json:"ownerRefer"`
}
