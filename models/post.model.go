package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID     string `gorm:"primaryKey;type:string;default:uuid_generate_v4()"`
	UserID string
	Text   string `gorm:"size:255;"`
	Url    string
	Likes  float32
}
