package models

import (
	"fmt"
	"time"
)

type Migration struct {
	Name      string `gorm:"text;primary_key;" json:"name"`
	Done      bool   `sql:"DEFAULT:false"`
	CreatedAt time.Time
}

type migrater interface {
	up(*DB) error
	getName() string
}

type migrationsApi struct {
	count   uint
	pending uint
	db      *DB
	migs    map[string]migrater
}

/*

	Init the migration queue to be ran

*/
func (migApi *migrationsApi) loadQueue() {

	// init migs map
	migApi.migs = make(map[string]migrater)

	// queue up migrations
	migApi.register(sessionForeignKey{})
	migApi.register(addTestData{})

}

/*

	Queue up a migration to be checked and executed

*/
func (migApi *migrationsApi) register(mig migrater) {

	migApi.migs[mig.getName()] = mig

}

/*

	Execute Migrations

*/
func (migApi *migrationsApi) run() error {

	toRun := map[string]migrater{}

	for name, mig := range migApi.migs {

		migApi.count = migApi.count + 1

		// attempt to find already run migration
		migModel := Migration{Name: name}
		migApi.db.Client.Find(&migModel)

		// mark as pending if not found
		if migModel.Done == false {
			migApi.pending = migApi.pending + 1
			toRun[name] = mig
		}

	}

	fmt.Printf("Total Migrations Found: %d ", migApi.count)
	fmt.Printf("(%d of which are Pending Migrations)\n", migApi.pending)

	// run pending migrations
	for name, mig := range toRun {

		// log which migration is executing
		fmt.Printf("Running Migration: %s\n", name)

		// run it
		err := mig.up(migApi.db)

		// early out if errors are encountered
		if err != nil {
			return err
		}

		// save migration status if successful
		migApi.db.Client.Create(&Migration{
			Name:      name,
			Done:      true,
			CreatedAt: time.Now(),
		})
	}

	return nil

}

/*

	Use a special name for migrations table

*/
func (mig Migration) TableName() string {

	return "$migrations"

}
