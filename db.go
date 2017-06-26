package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/matryer/try.v1"
	"math"
	"time"
)

const (
	dbDriver           = "postgres"
	dbOpts             = "host=db user=postgres dbname=postgres sslmode=disable password=postgres"
	DB_CONNECT_RETRIES = 10
	RETRY_FACTOR       = 1.7
)

type DB struct {
	client *gorm.DB
	users  *Users
}

func (dB DB) sync() {

	// sync schema
	dB.client.AutoMigrate(&User{})

}

func (dB DB) addTestData() {

	// test data creation
	dB.users.create("testFirst", "testLast")

}

func (dB *DB) connect() {

	var db *gorm.DB

	// retry connecting to database until threshold reached or successful
	// connection
	err := try.Do(func(attempt int) (bool, error) {

		var err error
		shouldRetry := attempt <= DB_CONNECT_RETRIES
		timeout := time.Second * time.Duration(math.Pow(RETRY_FACTOR, float64(attempt)))

		// connect
		fmt.Printf("db connection attempt: %v\n", attempt)
		db, err = gorm.Open(dbDriver, dbOpts)

		// connect err
		if err != nil {
			time.Sleep(timeout)
			return shouldRetry, err
		}

		// ping err
		err = db.DB().Ping()
		if err != nil {
			time.Sleep(timeout)
		}

		return shouldRetry, err

	})

	// fail on exhausted retries
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("Could not connect to db")
	}

	// attach db and models to struct
	dB.client = db
	dB.users = &Users{db}

}
