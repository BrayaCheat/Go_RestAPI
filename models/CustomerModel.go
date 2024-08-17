package models

import "gorm.io/gorm"

type Customer struct{
	ID uint `gorm:"primaryKey;autoIncrement;type:bigint"`
	Name string `gorm:"size:255"`
	Email string `gorm:"unique;size:255"`
	Phone string `gorm:"size:255"`
}

// TableName overrides the default table name.
func (Customer) TableName() string {
	return "tb_customer"
}

//migrate schema
func Migrate(db *gorm.DB){
	db.AutoMigrate(&Customer{})
}