package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null; unique"`
	Password string `gorm:"not null"`
	Items    []Item `gorm:"constraint:OnDelete:Cascade;"`
}
