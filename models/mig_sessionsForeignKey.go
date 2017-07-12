package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type sessionForeignKey struct{}

func (s sessionForeignKey) up(db *gorm.DB) error {

	// add the session's user_id forein key
	db.Model(&Session{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	return nil

}

func (s sessionForeignKey) getName() string {

	return "$sessionForeignKey"

}
