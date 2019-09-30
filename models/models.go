package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null" json:"-"`
}

// Note model
type Note struct {
	gorm.Model
	UserID uint
	User   User   `gorm:"foreignkey:UserID;association_foreignkey:ID;save_associations:false"`
	Title  string `gorm:"index"`
	Body   string `gorm:"index"`
}
