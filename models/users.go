package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Username string    `gorm:"text;not null;unique" json:"username"`
	First    string    `gorm:"text" json:"first"`
	Last     string    `gorm:"text" json:"last"`
	Password string    `gorm:"text" json:"-"`
	Sessions []Session `gorm:"ForeignKey:UserID" json:"sessions"`
}

type UsersApi struct {
	db *gorm.DB
}

// create a user in the db
func (u UsersApi) Create(un string, f string, l string, p string) {

	c := newCrypt()
	pw, _ := c.hash(p)
	u.db.Create(&User{
		Username: un,
		First:    f,
		Last:     l,
		Password: pw,
	})

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

func (u UsersApi) Login(un string, pw string) (User, error) {

	c := newCrypt()
	var user User

	// look up user
	u.db.Where("username = ?", un).Find(&user)

	if user.Username != un {
		return user, fmt.Errorf("No user found by that username")
	}

	if c.check(pw, user.Password) {
		return user, nil
	}

	return user, fmt.Errorf("Invalid login")

}
