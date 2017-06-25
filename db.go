package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	client *gorm.DB
	users  *Users
}

const (
	dbDriver = "postgres"
	dbOpts   = "host=db user=postgres dbname=postgres sslmode=disable password=postgres"
)

func (dB DB) sync() {

	// sync schema
	dB.client.AutoMigrate(&User{})
}

func (dB DB) addTestData() {

	// test data creation
	dB.users.create("testFirst", "testLast")

}

func (dB *DB) connect() {

	db := new(gorm.DB)
	users := new(Users)

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
	dB.client = db
	users.db = db
	dB.users = users
}
