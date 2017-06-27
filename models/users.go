package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	First string `gorm:"text" json:"first"`
	Last  string `gorm:"text" json:"last"`
}

type UsersApi struct {
	db *gorm.DB
}

// create a user in the db
func (u UsersApi) Create(f string, l string) {

	u.db.Create(&User{First: f, Last: l})

}
