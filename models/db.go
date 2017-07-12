package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/matryer/try.v1"
	"math"
	"time"
)

const (
	dbDriver    = "postgres"
	dbOpts      = "host=db user=postgres dbname=postgres sslmode=disable password=postgres"
	maxRetries  = 10
	retryFactor = 1.7
)

type DB struct {
	Client   *gorm.DB
	Users    *UsersApi
	Sessions *SessionsApi
}

func (dB *DB) Start() {

	dB.connect()
	dB.sync()
	dB.addTestData()

}

func (dB *DB) sync() {

	// wire up migration helper
	migrations := migrationsApi{db: dB.Client}
	migrations.loadQueue()

	// attach models
	dB.Users = &UsersApi{dB.Client}
	dB.Sessions = &SessionsApi{dB.Client}

	// sync schema
	dB.Client.AutoMigrate(&Migration{})
	dB.Client.AutoMigrate(&User{})
	dB.Client.AutoMigrate(&Session{})

	// run migrations
	err := migrations.run()

	if err != nil {
		panic(err)
	}

}

func (dB DB) addTestData() {

	// test data creation
	dB.Users.Create("admin", "testFirst", "testLast", "testPassword")

}

func (dB *DB) connect() {

	var db *gorm.DB

	// retry connecting to database until threshold reached or successful
	// connection
	err := try.Do(func(attempt int) (bool, error) {

		var err error
		shouldRetry := attempt <= maxRetries
		timeout := time.Second * time.Duration(math.Pow(retryFactor, float64(attempt)))

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

	// attach db to struct
	dB.Client = db

}
