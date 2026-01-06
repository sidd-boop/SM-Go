package thought

import "gorm.io/gorm"

type Thought struct{
	gorm.Model
	Content string `gorm :"type:text,not null"`
	Tag string `gorm:"index"`
	UserID uint `gorm:"not null"`
}

func Migrate(db *gorm.DB) error{
	return db.AutoMigrate(&Thought{})
}