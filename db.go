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

func (dB DB) init() {

	driver := "postgres"
	opts := "host=db user=postgres dbname=postgres sslmode=disable password=postgres"

	// TODO: add retry logic for db connection

	// attempt connection
	db, err := gorm.Open(driver, opts)

	// shut down on errors
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// attach db and models to struct
	dB.db = db
	dB.users = &Users{db: db}

	// sync schema
	db.AutoMigrate(&User{})

	// test data creation
	dB.users.create("admin", "admin")
}
