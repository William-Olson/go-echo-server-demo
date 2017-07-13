package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for gorm
)

// User : model for app users
type User struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Username string    `gorm:"text;not null;unique" json:"username"`
	First    string    `gorm:"text" json:"first"`
	Last     string    `gorm:"text" json:"last"`
	Password string    `gorm:"text" json:"-"`
	Sessions []Session `gorm:"ForeignKey:UserID" json:"sessions"`
}

// UsersAPI : used for interacting with user models
type UsersAPI struct {
	db *gorm.DB
}

// Create : create a user in the db
func (u UsersAPI) Create(un string, f string, l string, p string) {

	c := newCrypt()
	pw, _ := c.hash(p)
	u.db.Create(&User{
		Username: un,
		First:    f,
		Last:     l,
		Password: pw,
	})

}

// ByID : get a user by its id
func (u UsersAPI) ByID(id uint) User {

	var user User
	u.db.Find(&user, id)
	return user

}

// GetAll : get all current users in db
func (u UsersAPI) GetAll() []User {

	var users []User
	u.db.Find(&users)
	return users

}

// Login : attempt to login a user or error if login failed
func (u UsersAPI) Login(un string, pw string) (User, error) {

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
