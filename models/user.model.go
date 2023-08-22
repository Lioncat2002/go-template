package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email  string `gorm:"size:255;not null;unique"`
	Name   string `gorm:"size:255;"`
	Author []Item //things that the user has posted to sell on the platform
	Owns   []Item `gorm:"foreignKey:CurrentOwnerID"`
	Coins  float32
}
