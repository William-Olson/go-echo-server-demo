package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	First string `gorm:"text"`
	Last  string `gorm:"text"`
}

type Users struct {
	db *gorm.DB
}

func (u Users) create(f string, l string) {
	u.db.Create(&User{First: f, Last: l})
}
