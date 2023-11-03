package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:string;default:uuid_generate_v4()"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;"`
	Name     string `gorm:"size:255;"`
	Author   []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Define the User foreign key relationship for book rating
type UserLikes struct {
	UserID string `gorm:"primaryKey;autoIncrement:false"`
	PostID string `gorm:"primaryKey;autoIncrement:false"`
}
