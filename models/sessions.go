package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Session struct {
	Token        string    `gorm:"primary_key;text" json:"token"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	ClientIP     string    `gorm:"text" json:"client_ip"`
	LastActivity time.Time `json:"last_activity"`
}

type SessionsApi struct {
	db *gorm.DB
}

// create a user in the db
func (sesh SessionsApi) Create(user User, token, client string) {

	sesh.db.Create(&Session{
		UserID:       user.ID,
		Token:        token,
		ClientIP:     client,
		LastActivity: time.Now(),
	})

}

func (sesh SessionsApi) ByToken(token uint) Session {

	var session Session
	sesh.db.Find(&session, token)
	return session

}

func (sesh SessionsApi) Activity(s *Session) {

	sesh.db.Update(&s, "last_activity", time.Now())

}
