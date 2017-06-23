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

type DB struct {
	db    *gorm.DB
	users *Users
}

const (
	dbDriver = "postgres"
	dbOpts   = "host=db user=postgres dbname=postgres sslmode=disable password=postgres"
)

var users = &Users{}
var db *gorm.DB

func (dB DB) connect() *gorm.DB {

	// TODO: add retry logic for db connection

	// attempt connection
	db, err := gorm.Open(dbDriver, dbOpts)

	// shut down on errors
	if err != nil {
		panic("failed to connect database")
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	// attach db and models to struct
	dB.db = db
	users.db = db
	dB.users = users

	return db
}
