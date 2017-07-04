package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	First    string `gorm:"text" json:"first"`
	Last     string `gorm:"text" json:"last"`
	Password string `gorm:"text" json:"-"`
}

type UsersApi struct {
	db *gorm.DB
}

// create a user in the db
func (u UsersApi) Create(f string, l string, p string) {

	u.db.Create(&User{First: f, Last: l, Password: p})

}

func (u UsersApi) ById(id uint) User {

	var user User
	u.db.Find(&user, id)
	return user

}

func (u UsersApi) GetAll() []User {

	var users []User
	u.db.Find(&users)
	return users

}
