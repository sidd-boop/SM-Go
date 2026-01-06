package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"UniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func Migrate(db *gorm.DB) error{
	return db.AutoMigrate(&User{})
}
