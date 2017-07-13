package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // required for gorm
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

// DB : The db dependency
type DB struct {
	Client   *gorm.DB
	Users    *UsersAPI
	Sessions *SessionsAPI
}

// Start : connect the db client, sync models, and run migrations
func (dB *DB) Start() {

	dB.connect()
	dB.sync()

}

/*

	Add missing fields to tables & migrate the db

*/
func (dB *DB) sync() {

	// wire up migration helper
	migrations := migrationsAPI{db: dB}
	migrations.loadQueue()

	// attach models
	dB.Users = &UsersAPI{dB.Client}
	dB.Sessions = &SessionsAPI{dB.Client}

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

/*

	Establish a db client connection

*/
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
